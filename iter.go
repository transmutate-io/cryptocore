package cryptocore

import (
	"errors"
	"time"

	"transmutate.io/pkg/cryptocore/tx"

	"transmutate.io/pkg/cryptocore/block"
	"transmutate.io/pkg/cryptocore/types"
)

type blockIterator struct {
	cl         Client
	nextHeight uint64
	nextHash   types.Bytes
}

func NewBlockIterator(cl Client, blockHeight uint64) *blockIterator {
	var bh uint64
	if blockHeight < 2 {
		bh = 1
	} else {
		bh = blockHeight
	}
	return &blockIterator{cl: cl, nextHeight: bh}
}

var ErrNoBlock = errors.New("no block")

func (bi *blockIterator) block(hash types.Bytes) (block.Block, error) {
	blk, err := bi.cl.Block(hash)
	if err != nil {
		return nil, err
	}
	bi.nextHeight = uint64(blk.Height()) + 1
	bi.nextHash = blk.NextBlockHash()
	return blk, nil
}

func (bi *blockIterator) blockAtHeight(height uint64) (block.Block, error) {
	bh, err := bi.cl.BlockHash(uint64(height) + 1)
	if err != nil {
		if e, ok := err.(*ClientError); ok && e.Code == -8 {
			return nil, ErrNoBlock
		}
		return nil, err
	}
	return bi.block(bh)
}

func (bi *blockIterator) Next() (block.Block, error) {
	if len(bi.nextHash) != 0 {
		return bi.block(bi.nextHash)
	}
	return bi.blockAtHeight(bi.nextHeight)
}

var iteratorTimeout = time.Second

func (bi *blockIterator) WaitNext() (block.Block, error) {
	for {
		blk, err := bi.Next()
		if err != nil {
			if err != ErrNoBlock {
				return nil, err
			}
			time.Sleep(iteratorTimeout)
			continue
		}
		return blk, nil
	}
}

type txIterator struct {
	blkIter *blockIterator
	blk     block.Block
	txIdx   int
}

func NewTransactionIterator(cl Client, blockHeight uint64) *txIterator {
	return &txIterator{
		blkIter: NewBlockIterator(cl, blockHeight),
		txIdx:   0,
	}
}

func (ti *txIterator) Next() (tx.Tx, error) {
	if ti.blk == nil {
		var err error
		if ti.blk, err = ti.blkIter.Next(); err != nil {
			return nil, err
		}
		ti.txIdx = 0
	}
	txs := ti.blk.Transactions()
	if ti.txIdx > len(txs)-1 {
		ti.blk = nil
		return ti.Next()
	}
	idx := ti.txIdx
	ti.txIdx++
	return ti.blkIter.cl.Transaction(txs[idx])
}

func (ti *txIterator) WaitNext() (tx.Tx, error) {
	for {
		tx, err := ti.Next()
		if err != nil {
			if err != ErrNoBlock {
				return nil, err
			}
			time.Sleep(iteratorTimeout)
			continue
		}
		return tx, nil
	}
}
