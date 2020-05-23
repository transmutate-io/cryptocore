package cryptocore

import (
	"crypto/tls"

	"transmutate.io/pkg/cryptocore/types"
)

type dogeClient struct{ *baseClient }

func NewClientDOGE(addr, user, pass string, tlsConf *tls.Config) Client {
	return &dogeClient{baseClient: newBaseClient(addr, user, pass, tlsConf)}
}

func (c *dogeClient) RawBlock(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getblock", args(hash, false))
}

func (c *dogeClient) Block(hash types.Bytes) (*types.Block, error) {
	return c.block(hash, true)
}

func (c *dogeClient) Transaction(hash types.Bytes) (*types.Transaction, error) {
	return c.transaction(hash)
}
