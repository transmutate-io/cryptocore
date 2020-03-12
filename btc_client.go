package btccore

import (
	"transmutate.io/pkg/btccore/types"
)

type btcClient struct{ *baseClient }

func NewClientBTC(addr, user, pass string, useHTTPS bool) Client {
	return &btcClient{baseClient: newBaseClient(addr, user, pass, useHTTPS)}
}

func (c *btcClient) RawBlock(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getblock", args(hash, 0))
}

func (c *btcClient) Block(hash types.Bytes) (*types.Block, error) {
	return c.block(hash, 1)
}