package cryptocore

import (
	"strings"

	"transmutate.io/pkg/cryptocore/types"
)

type bchClient struct{ *baseClient }

func NewClientBCH(addr, user, pass string, tlsConf *TLSConfig) Client {
	return &bchClient{baseClient: newBaseClient(addr, user, pass, tlsConf)}
}

func (c *bchClient) RawBlock(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getblock", args(hash, 0))
}

func (c *bchClient) Block(hash types.Bytes) (*types.Block, error) {
	r, err := c.block(hash, 1)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *bchClient) Transaction(hash types.Bytes) (*types.Transaction, error) {
	r, err := c.transaction(hash)
	if err != nil {
		return nil, err
	}
	for _, i := range r.Outputs {
		for j, k := range i.UnlockScript.Addresses {
			t := strings.Split(k, ":")
			if len(t) < 2 {
				i.UnlockScript.Addresses[j] = t[0]
			} else {
				i.UnlockScript.Addresses[j] = t[1]
			}
		}
	}
	return r, nil
}
