package btccore

import "transmutate.io/pkg/btccore/types"

type ltcClient struct{ *baseClient }

func NewClientLTC(addr, user, pass string, useHTTPS bool) Client {
	return &ltcClient{baseClient: newBaseClient(addr, user, pass, useHTTPS)}
}

func (c *ltcClient) Block(hash types.Bytes) (*types.Block, error) {
	return c.block(hash, 1)
}

func (c *ltcClient) RawBlock(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getblock", args(hash, 0))
}
