package cryptocore

import (
	"github.com/transmutate-io/cryptocore/block"
	"github.com/transmutate-io/cryptocore/tx"
	"github.com/transmutate-io/cryptocore/types"
)

type dcrClient struct{ *baseClient }

func NewClientDCR(addr, user, pass string, tlsConf *TLSConfig) (Client, error) {
	b, err := newBaseClient(addr, user, pass, tlsConf)
	if err != nil {
		return nil, err
	}
	return &dcrClient{baseClient: b}, nil
}

func (c *dcrClient) RawBlock(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getblock", args(hash, false))
}

func (c *dcrClient) Block(hash types.Bytes) (block.Block, error) {
	r := &block.BlockDCR{}
	if err := c.block(r, args(hash, true)); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *dcrClient) Transaction(hash types.Bytes) (tx.Tx, error) {
	r := &tx.TxDCR{}
	if err := c.transaction(r, args(hash, 1)); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *dcrClient) CanGenerateBlocksToAddress() bool { return false }

func (c *dcrClient) GenerateBlocksToAddress(nBlocks int, addr string) ([]types.Bytes, error) {
	panic("can't call \"generatetoaddress\" method")
}

func (c *dcrClient) Balance(minConf int64) (types.Amount, error) {
	var (
		a = mkArgs(2, "*")
		r = &struct {
			Balances []struct {
				Total types.Amount `json:"total"`
			} `json:"balances"`
		}{}
	)
	if minConf >= 0 {
		a = append(a, minConf)
	}
	if err := c.do("getbalance", a, &r); err != nil {
		return "", err
	}
	total := uint64(0)
	for _, i := range r.Balances {
		total += i.Total.UInt64((8))
	}
	return types.NewAmount(total, 8), nil
}
