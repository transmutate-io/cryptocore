package cryptocore

import (
	"github.com/transmutate-io/cryptocore/block"
	"github.com/transmutate-io/cryptocore/tx"
	"github.com/transmutate-io/cryptocore/types"
)

var (
	_ Client                 = (*dogeClient)(nil)
	_ TargetedBlockGenerator = (*dogeClient)(nil)
	_ AddressGenerator       = (*dogeClient)(nil)
	_ Sender                 = (*dogeClient)(nil)
	_ RawTransactionSender   = (*dogeClient)(nil)
	_ Balancer               = (*dogeClient)(nil)
	_ AddressLister          = (*dogeClient)(nil)
)

type dogeClient struct{ baseBTCClient }

func NewClientDOGE(addr, user, pass string, tlsConf *TLSConfig) (Client, error) {
	c, err := newJsonRpcClient(addr, user, pass, tlsConf)
	if err != nil {
		return nil, err
	}
	return &dogeClient{baseBTCClient{*c}}, nil
}

func (c *dogeClient) RawBlock(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getblock", args(hash, false))
}

func (c *dogeClient) Block(hash types.Bytes) (block.Block, error) {
	r := &block.BlockDOGE{}
	if err := c.block(r, args(hash, true)); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *dogeClient) Transaction(hash types.Bytes) (tx.Tx, error) {
	r := &tx.TxDOGE{}
	if err := c.transaction(r, args(hash, true)); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *dogeClient) GenerateToAddress(nBlocks int, addr string) ([]types.Bytes, error) {
	return c.doSliceBytes("generatetoaddress", args(nBlocks, addr))
}
