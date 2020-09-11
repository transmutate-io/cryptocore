package cryptocore

import (
	"bytes"
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/transmutate-io/cryptocore/block"
	"github.com/transmutate-io/cryptocore/tx"
	"github.com/transmutate-io/cryptocore/types"
)

var (
	_ Client           = (*ethClient)(nil)
	_ TargetedBalancer = (*ethClient)(nil)
)

type ethClient struct{ cl *ethclient.Client }

func NewClientETH(addr, user, pass string, tlsConf *TLSConfig) (Client, error) {
	c, err := ethclient.Dial(clientURL(addr, user, pass, tlsConf))
	if err != nil {
		return nil, err
	}
	return &ethClient{c}, nil
}

func (c *ethClient) BalanceOf(addr string, minConf int64) (types.Amount, error) {
	b, err := c.cl.BalanceAt(context.Background(), common.HexToAddress(addr), nil)
	if err != nil {
		return "", err
	}
	return types.NewAmountBig(b, 18), nil
}

func (c *ethClient) getBlockAtHeight(h *big.Int) (*ethtypes.Block, error) {
	return c.cl.BlockByNumber(context.Background(), h)
}

func (c *ethClient) BlockCount() (uint64, error) {
	blk, err := c.getBlockAtHeight(nil)
	if err != nil {
		return 0, err
	}
	return blk.NumberU64(), nil
}

func (c *ethClient) BlockHash(height uint64) (types.Bytes, error) {
	blk, err := c.getBlockAtHeight(nil)
	if err != nil {
		return nil, err
	}
	return types.Bytes(blk.Hash().Bytes()), nil
}

func (c *ethClient) RawBlock(hash types.Bytes) (types.Bytes, error) {
	blk, err := c.getBlockAtHeight(nil)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(make([]byte, 0, 1024))
	if err = blk.EncodeRLP(b); err != nil {
		return nil, err
	}
	return types.Bytes(b.Bytes()), nil
}

func (c *ethClient) Block(hash types.Bytes) (block.Block, error) {
	blk, err := c.getBlockAtHeight(nil)
	if err != nil {
		return nil, err
	}
	return &block.BlockETH{Block: blk}, nil
}

var ErrInvalidTransaction = errors.New("invalid transaction")

func (c *ethClient) SendTransaction(t tx.Tx) (types.Bytes, error) {
	tt, ok := t.(*tx.TxETH)
	if !ok {
		return nil, ErrInvalidTransaction
	}
	return nil, c.cl.SendTransaction(context.Background(), tt.Tx)
}

func (c *ethClient) getTransaction(hash types.Bytes) (*ethtypes.Transaction, bool, error) {
	return c.cl.TransactionByHash(context.Background(), common.BytesToHash(hash))
}

func (c *ethClient) RawTransaction(hash types.Bytes) (types.Bytes, error) {
	tx, pending, err := c.getTransaction(hash)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(make([]byte, 0, 1024))
	if err = tx.EncodeRLP(b); err != nil {
		return nil, err
	}
	if pending {
		err = ErrTransactionPending
	}
	return types.Bytes(b.Bytes()), err
}

var ErrTransactionPending = errors.New("transactin pending")

func (c *ethClient) Transaction(hash types.Bytes) (tx.Tx, error) {
	t, pending, err := c.getTransaction(hash)
	if err != nil {
		return nil, err
	}
	if pending {
		err = ErrTransactionPending
	}
	return &tx.TxETH{Tx: t}, err
}
