package cryptocore

import (
	"errors"

	"github.com/transmutate-io/cryptocore/types"
)

type baseBTCClient struct{ jsonRPCClient }

func (c *baseBTCClient) BlockCount() (uint64, error) { return c.doUint64("getblockcount", nil) }

var ErrNoBlock = errors.New("no block")

func (c *baseBTCClient) BlockHash(height uint64) (types.Bytes, error) {
	r, err := c.doBytes("getblockhash", args(height))
	if err != nil {
		if e, ok := err.(*ClientError); ok && e.Code == -8 {
			return nil, ErrNoBlock
		}
		return nil, err
	}
	return r, nil
}

func (c *baseBTCClient) NewAddress() (string, error) { return c.doString("getnewaddress", nil) }

func mkArgs(n int, args ...interface{}) []interface{} {
	var sz int
	if argsSz := len(args); n > argsSz {
		sz = n
	} else {
		sz = argsSz
	}
	return append(make([]interface{}, 0, sz), args...)
}

func args(a ...interface{}) []interface{} { return a }

func (c *baseBTCClient) SendToAddress(addr string, value types.Amount) (types.Bytes, error) {
	return c.doBytes("sendtoaddress", args(addr, value))
}

var ErrNonFinal = errors.New("non final")

func (c *baseBTCClient) SendRawTransaction(tx types.Bytes) (types.Bytes, error) {
	r, err := c.doBytes("sendrawtransaction", args(tx.Hex()))
	if err != nil {
		e, ok := err.(*ClientError)
		if ok && e.Code == -26 {
			return nil, ErrNonFinal
		}
		return nil, err
	}
	return r, nil
}

func (c *baseBTCClient) RawTransaction(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getrawtransaction", args(hash.Hex(), false))
}

func (c *baseBTCClient) transaction(t interface{}, args interface{}) error {
	return c.do("getrawtransaction", args, t)
}

func (c *baseBTCClient) Balance(minConf int64) (types.Amount, error) {
	var (
		a = mkArgs(2, "*")
		r types.Amount
	)
	if minConf >= 0 {
		a = append(a, minConf)
	}
	if err := c.do("getbalance", a, &r); err != nil {
		return "", err
	}
	return r, nil
}

func (c *baseBTCClient) block(b interface{}, args interface{}) error {
	return c.do("getblock", args, b)
}

type addressFunds struct {
	WatchOnly     bool          `json:"involvesWatchonly"`
	Address       string        `json:"address"`
	Amount        types.Amount  `json:"amount"`
	Confirmations int           `json:"confirmations"`
	Label         string        `json:"label"`
	TxIDs         []types.Bytes `json:"txids"`
}

func (c *baseBTCClient) AvailableAddresses() ([]string, error) {
	addrFunds := []addressFunds{}
	if err := c.do("listreceivedbyaddress", args(0, true, true), &addrFunds); err != nil {
		return nil, err
	}
	r := make([]string, 0, len(addrFunds))
	for _, i := range addrFunds {
		r = append(r, i.Address)
	}
	return r, nil
}
