package cryptocore

import (
	"github.com/transmutate-io/cryptocore/block"
	"github.com/transmutate-io/cryptocore/tx"
	"github.com/transmutate-io/cryptocore/types"
)

type ltcClient struct{ *baseClient }

func NewClientLTC(addr, user, pass string, tlsConf *TLSConfig) Client {
	return &ltcClient{baseClient: newBaseClient(addr, user, pass, tlsConf)}
}

func (c *ltcClient) Block(hash types.Bytes) (block.Block, error) {
	r := &block.BlockLTC{}
	if err := c.block(r, args(hash, 1)); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *ltcClient) RawBlock(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getblock", args(hash, 0))
}

func (c *ltcClient) Transaction(hash types.Bytes) (tx.Tx, error) {
	r := &tx.TxLTC{}
	if err := c.transaction(r, args(hash, true)); err != nil {
		return nil, err
	}
	return r, nil
}
