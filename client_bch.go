package cryptocore

import (
	"strings"

	"github.com/transmutate-io/cryptocore/block"
	"github.com/transmutate-io/cryptocore/tx"
	"github.com/transmutate-io/cryptocore/types"
)

var (
	_ Client                 = (*bchClient)(nil)
	_ TargetedBlockGenerator = (*bchClient)(nil)
	_ AddressGenerator       = (*bchClient)(nil)
	_ Sender                 = (*bchClient)(nil)
	_ RawTransactionSender   = (*bchClient)(nil)
	_ Balancer               = (*bchClient)(nil)
	_ AddressLister          = (*bchClient)(nil)
)

type bchClient struct{ baseBTCClient }

func NewClientBCH(addr, user, pass string, tlsConf *TLSConfig) (Client, error) {
	c, err := newJsonRpcClient(addr, user, pass, tlsConf)
	if err != nil {
		return nil, err
	}
	return &bchClient{baseBTCClient{*c}}, nil
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

func (c *bchClient) GenerateToAddress(nBlocks int, addr string) ([]types.Bytes, error) {
	return c.doSliceBytes("generatetoaddress", args(nBlocks, addr))
}

func (c *bchClient) AvailableAddresses() ([]string, error) {
	r, err := c.baseBTCClient.AvailableAddresses()
	if err != nil {
		return nil, err
	}
	for i, v := range r {
		r[i] = trimAddress(v)
	}
	return r, nil
}
