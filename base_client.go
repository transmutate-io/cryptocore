package cryptocore

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"transmutate.io/pkg/cryptocore/types"
)

type baseClient struct {
	Address   string
	Username  string
	Password  string
	TLSConfig *tls.Config
}

func newBaseClient(addr, user, pass string, tlsConf *tls.Config) *baseClient {
	return &baseClient{
		Address:   addr,
		Username:  user,
		Password:  pass,
		TLSConfig: tlsConf,
	}
}

func (c *baseClient) do(method string, params interface{}, r interface{}) error {
	b := bytes.NewBuffer(make([]byte, 0, 1024))
	err := json.NewEncoder(b).Encode(&rpcRequest{
		JsonRPC: "1.0",
		ID:      "go-cryptocore",
		Method:  method,
		Params:  params,
	})
	if err != nil {
		return err
	}
	var useHTTPS string
	if c.TLSConfig != nil {
		useHTTPS = "s"
	}
	resp, err := http.Post(fmt.Sprintf("http%s://%s:%s@%s/", useHTTPS, c.Username, c.Password, c.Address), "application/json", b)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	rr := &rpcResponse{Result: r}
	if err = json.NewDecoder(resp.Body).Decode(rr); err != nil {
		return err
	}
	if rr.Error != nil {
		return rr.Error
	}
	return nil
}

func (c *baseClient) doString(method string, args interface{}) (string, error) {
	var r string
	if err := c.do(method, args, &r); err != nil {
		return "", err
	}
	return r, nil
}

func (c *baseClient) doUint64(method string, args interface{}) (uint64, error) {
	var r uint64
	if err := c.do(method, args, &r); err != nil {
		return 0, err
	}
	return r, nil
}

func (c *baseClient) doBytes(method string, args interface{}) (types.Bytes, error) {
	var r types.Bytes
	if err := c.do(method, args, &r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *baseClient) doSliceBytes(method string, args interface{}) ([]types.Bytes, error) {
	var r []types.Bytes
	if err := c.do(method, args, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type rpcRequest struct {
	JsonRPC string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type rpcResponse struct {
	ID     string       `json:"id"`
	Error  *ClientError `json:"error"`
	Result interface{}  `json:"result"`
}

type ClientError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (we *ClientError) Error() string { return we.Message }

func (c *baseClient) NewAddress() (string, error) { return c.doString("getnewaddress", nil) }

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

func (c *baseClient) DumpPrivateKey(addr string) (string, error) {
	return c.doString("dumpprivkey", args(addr))
}

func (c *baseClient) GenerateToAddress(nBlocks int, addr string) ([]types.Bytes, error) {
	return c.doSliceBytes("generatetoaddress", args(nBlocks, addr))
}

func (c *baseClient) SendToAddress(addr string, value types.Amount) (types.Bytes, error) {
	return c.doBytes("sendtoaddress", args(addr, value))
}

func (c *baseClient) BlockCount() (uint64, error) { return c.doUint64("getblockcount", nil) }

func (c *baseClient) BlockHash(height uint64) (types.Bytes, error) {
	return c.doBytes("getblockhash", args(height))
}

func (c *baseClient) SendRawTransaction(tx types.Bytes) (types.Bytes, error) {
	return c.doBytes("sendrawtransaction", args(tx.Hex()))
}

func (c *baseClient) RawTransaction(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getrawtransaction", args(hash.Hex(), false))
}

func (c *baseClient) transaction(hash types.Bytes) (*types.Transaction, error) {
	r := &types.Transaction{}
	if err := c.do("getrawtransaction", args(hash, true), r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *baseClient) Balance(minConf int64) (types.Amount, error) {
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

func (c *baseClient) block(args ...interface{}) (*types.Block, error) {
	r := &types.Block{}
	if err := c.do("getblock", args, r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *baseClient) ReceivedByAddress(minConf, includeEmpty, includeWatchOnly interface{}) ([]*types.AddressFunds, error) {
	args := mkArgs(3)
	if minConf != nil {
		args = append(args, minConf)
	}
	if includeEmpty != nil {
		args = append(args, includeEmpty)
	}
	if includeWatchOnly != nil {
		args = append(args, includeWatchOnly)
	}
	r := []*types.AddressFunds{}
	if err := c.do("listreceivedbyaddress", args, &r); err != nil {
		return nil, err
	}
	return r, nil
}

var blockIteratorTimeout = time.Second

func NewBlockIterator(cl Client, blockHeight uint64) (BlockFunc, CloseFunc) {
	cc := make(chan struct{}, 1)
	bc := make(chan *types.Block)
	errc := make(chan error)
	go func() {
		defer close(bc)
		defer close(errc)
		var (
			bh  types.Bytes
			err error
		)
		for {
			if bh == nil {
				if bh, err = cl.BlockHash(blockHeight); err != nil {
					e, ok := err.(*ClientError)
					if !ok || e.Code != -8 {
						errc <- err
						return
					}
					time.Sleep(blockIteratorTimeout)
					continue
				}
			}
			blk, err := cl.Block(bh)
			if err != nil {
				errc <- err
				return
			}
			bh = blk.NextBlockHash
			blockHeight++
			select {
			case bc <- blk:
			case <-cc:
				return
			}
		}
	}()
	return func() (*types.Block, error) {
		select {
		case blk := <-bc:
			return blk, nil
		case err := <-errc:
			return nil, err
		}
	}, func() { close(cc) }
}

func NewTransactionIterator(cl Client, blockHeight uint64) (TransactionFunc, CloseFunc) {
	cc := make(chan struct{}, 1)
	tc := make(chan *types.Transaction)
	errc := make(chan error)
	go func() {
		defer close(tc)
		defer close(errc)
		nextBlk, closeIter := NewBlockIterator(cl, blockHeight)
		defer closeIter()
		for {
			blk, err := nextBlk()
			if err != nil {
				errc <- err
				return
			}
			for _, i := range blk.Transactions {
				tx, err := cl.Transaction(i)
				if err != nil {
					errc <- err
					return
				}
				tc <- tx
			}
		}
	}()
	return func() (*types.Transaction, error) {
		select {
		case tx := <-tc:
			return tx, nil
		case err := <-errc:
			return nil, err
		}
	}, func() { close(cc) }
}

// func (c *baseClient) NewTransactionIterator(firstBlockHeight int) (chan *types.Transaction, CloseFunc) {
// 	cc := make(chan struct{})
// 	// r := make(chan *types.Transaction)
// 	return nil, func() { close(cc) }
// }
