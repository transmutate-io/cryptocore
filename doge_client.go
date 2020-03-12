package btccore

import "transmutate.io/pkg/btccore/types"

type dogeClient struct{ *baseClient }

func NewClientDOGE(addr, user, pass string, useHTTPS bool) Client {
	return &dogeClient{baseClient: newBaseClient(addr, user, pass, useHTTPS)}
}

func (c *dogeClient) RawBlock(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getblock", args(hash, false))
}

func (c *dogeClient) Block(hash types.Bytes) (*types.Block, error) {
	return c.block(hash, true)
}
