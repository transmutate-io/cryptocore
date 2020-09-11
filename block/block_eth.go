package block

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/transmutate-io/cryptocore/tx"
	"github.com/transmutate-io/cryptocore/types"
)

var (
	_ Block                  = (*BlockETH)(nil)
	_ BackwardBlockNavigator = (*BlockETH)(nil)
	_ TransactionsFetcher    = (*BlockETH)(nil)
)

type BlockETH struct{ Block *ethtypes.Block }

func (blk *BlockETH) Hash() types.Bytes { return types.Bytes(blk.Block.Hash().Bytes()) }

func (blk *BlockETH) Height() int { return int(blk.Block.NumberU64()) }

func (blk *BlockETH) Time() types.UnixTime { return types.UnixTime(blk.Block.Time()) }

func (blk *BlockETH) PreviousBlockHash() types.Bytes {
	return types.Bytes(blk.Block.ParentHash().Bytes())
}

func (blk *BlockETH) Transactions() []tx.Tx {
	txs := blk.Block.Transactions()
	r := make([]tx.Tx, 0, len(txs))
	for _, i := range txs {
		r = append(r, &tx.TxETH{Tx: i})
	}
	return r
}
