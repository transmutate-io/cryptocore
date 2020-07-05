package cryptocore

import (
	"github.com/transmutate-io/cryptocore/block"
	"github.com/transmutate-io/cryptocore/tx"
	"github.com/transmutate-io/cryptocore/types"
)

type btcClient struct{ *baseClient }

func NewClientBTC(addr, user, pass string, tlsConf *TLSConfig) Client {
	return &btcClient{baseClient: newBaseClient(addr, user, pass, tlsConf)}
}

func (c *btcClient) RawBlock(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getblock", args(hash, 0))
}

func (c *btcClient) Block(hash types.Bytes) (block.Block, error) {
	r := &block.BlockBTC{}
	if err := c.block(r, args(hash, 1)); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *btcClient) Transaction(hash types.Bytes) (tx.Tx, error) {
	r := &tx.TxBTC{}
	if err := c.transaction(r, args(hash, true)); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *btcClient) CanGenerateBlocks() bool { return false }

func (c *btcClient) GenerateBlocks(nBlocks int) ([]types.Bytes, error) {
	panic("can't call \"generate\" method")
}
