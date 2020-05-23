package cryptocore

import (
	"crypto/tls"

	"transmutate.io/pkg/cryptocore/types"
)

type ltcClient struct{ *baseClient }

func NewClientLTC(addr, user, pass string, tlsConf *tls.Config) Client {
	return &ltcClient{baseClient: newBaseClient(addr, user, pass, tlsConf)}
}

func (c *ltcClient) Block(hash types.Bytes) (*types.Block, error) {
	return c.block(hash, 1)
}

func (c *ltcClient) RawBlock(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getblock", args(hash, 0))
}

func (c *ltcClient) Transaction(hash types.Bytes) (*types.Transaction, error) {
	return c.transaction(hash)
}
