package cryptocore

import (
	"github.com/transmutate-io/cryptocore/block"
	"github.com/transmutate-io/cryptocore/tx"
	"github.com/transmutate-io/cryptocore/types"
)

var (
	_ Client                 = (*btcClient)(nil)
	_ TargetedBlockGenerator = (*btcClient)(nil)
	_ AddressGenerator       = (*btcClient)(nil)
	_ Sender                 = (*btcClient)(nil)
	_ RawTransactionSender   = (*btcClient)(nil)
	_ Balancer               = (*btcClient)(nil)
	_ AddressLister          = (*btcClient)(nil)
)

type btcClient struct{ baseBTCClient }

func NewClientBTC(addr, user, pass string, tlsConf *TLSConfig) (Client, error) {
	c, err := newJsonRpcClient(addr, user, pass, tlsConf)
	if err != nil {
		return nil, err
	}
	return &btcClient{baseBTCClient{*c}}, nil
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

func (c *btcClient) GenerateToAddress(nBlocks int, addr string) ([]types.Bytes, error) {
	return c.doSliceBytes("generatetoaddress", args(nBlocks, addr))
}
