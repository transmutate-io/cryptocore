package cryptocore

import (
	"github.com/transmutate-io/cryptocore/block"
	"github.com/transmutate-io/cryptocore/tx"
	"github.com/transmutate-io/cryptocore/types"
)

type dogeClient struct{ *baseClient }

func NewClientDOGE(addr, user, pass string, tlsConf *TLSConfig) (Client, error) {
	b, err := newBaseClient(addr, user, pass, tlsConf)
	if err != nil {
		return nil, err
	}
	return &dogeClient{baseClient: b}, nil
}

func (c *dogeClient) RawBlock(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getblock", args(hash, false))
}

func (c *dogeClient) Block(hash types.Bytes) (block.Block, error) {
	r := &block.BlockDOGE{}
	if err := c.block(r, args(hash, true)); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *dogeClient) Transaction(hash types.Bytes) (tx.Tx, error) {
	r := &tx.TxDOGE{}
	if err := c.transaction(r, args(hash, true)); err != nil {
		return nil, err
	}
	return r, nil
}
