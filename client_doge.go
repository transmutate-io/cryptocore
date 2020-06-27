package cryptocore

import (
	"transmutate.io/pkg/cryptocore/block"
	"transmutate.io/pkg/cryptocore/tx"
	"transmutate.io/pkg/cryptocore/types"
)

type dogeClient struct{ *baseClient }

func NewClientDOGE(addr, user, pass string, tlsConf *TLSConfig) Client {
	return &dogeClient{baseClient: newBaseClient(addr, user, pass, tlsConf)}
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
