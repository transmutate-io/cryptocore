package cryptocore

import (
	"strings"

	"github.com/transmutate-io/cryptocore/block"
	"github.com/transmutate-io/cryptocore/tx"
	"github.com/transmutate-io/cryptocore/types"
)

type bchClient struct{ *baseClient }

func NewClientBCH(addr, user, pass string, tlsConf *TLSConfig) (Client, error) {
	b, err := newBaseClient(addr, user, pass, tlsConf)
	if err != nil {
		return nil, err
	}
	return &bchClient{baseClient: b}, nil
}

func (c *bchClient) RawBlock(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getblock", args(hash, 0))
}

func (c *bchClient) Block(hash types.Bytes) (block.Block, error) {
	r := &block.BlockBCH{}
	if err := c.block(r, args(hash, 1)); err != nil {
		return nil, err
	}
	return r, nil
}

func trimAddress(addr string) string {
	t := strings.Split(addr, ":")
	if len(t) < 2 {
		return t[0]
	} else {
		return t[1]
	}
}

func (c *bchClient) Transaction(hash types.Bytes) (tx.Tx, error) {
	r := &tx.TxBCH{}
	if err := c.transaction(r, args(hash, true)); err != nil {
		return nil, err
	}
	for _, i := range r.Outputs() {
		for j, k := range i.LockScript().Addresses() {
			i.LockScript().Addresses()[j] = trimAddress(k)
		}
	}
	return r, nil
}

func (c *bchClient) ReceivedByAddress(minConf, includeEmpty, includeWatchOnly interface{}) ([]*types.AddressFunds, error) {
	r, err := c.baseClient.ReceivedByAddress(minConf, includeEmpty, includeWatchOnly)
	if err != nil {
		return nil, err
	}
	for _, i := range r {
		i.Address = trimAddress((i.Address))
	}
	return r, nil
}
