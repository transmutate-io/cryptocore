package cryptocore

import (
	"github.com/transmutate-io/cryptocore/block"
	"github.com/transmutate-io/cryptocore/tx"
	"github.com/transmutate-io/cryptocore/types"
)

var (
	_ Client                 = (*ltcClient)(nil)
	_ TargetedBlockGenerator = (*ltcClient)(nil)
	_ AddressGenerator       = (*ltcClient)(nil)
	_ Sender                 = (*ltcClient)(nil)
	_ RawTransactionSender   = (*ltcClient)(nil)
	_ Balancer               = (*ltcClient)(nil)
	_ AddressLister          = (*ltcClient)(nil)
)

type ltcClient struct{ baseBTCClient }

func NewClientLTC(addr, user, pass string, tlsConf *TLSConfig) (Client, error) {
	c, err := newJsonRpcClient(addr, user, pass, tlsConf)
	if err != nil {
		return nil, err
	}
	return &ltcClient{baseBTCClient{*c}}, nil
}

func (c *ltcClient) Block(hash types.Bytes) (block.Block, error) {
	r := &block.BlockLTC{}
	if err := c.block(r, args(hash, 1)); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *ltcClient) RawBlock(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getblock", args(hash, 0))
}

func (c *ltcClient) Transaction(hash types.Bytes) (tx.Tx, error) {
	r := &tx.TxLTC{}
	if err := c.transaction(r, args(hash, true)); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *ltcClient) GenerateToAddress(nBlocks int, addr string) ([]types.Bytes, error) {
	return c.doSliceBytes("generatetoaddress", args(nBlocks, addr))
}
