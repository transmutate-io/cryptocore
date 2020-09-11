package block

import (
	"encoding/json"

	"github.com/transmutate-io/cryptocore/types"
)

var (
	_ Block                  = (*BlockDOGE)(nil)
	_ TransactionsLister     = (*BlockDOGE)(nil)
	_ ConfirmationCounter    = (*BlockDOGE)(nil)
	_ ForwardBlockNavigator  = (*BlockDOGE)(nil)
	_ BackwardBlockNavigator = (*BlockDOGE)(nil)
)

type BlockDOGE struct{ baseBTCBlock }

func (blk *BlockDOGE) Hash() types.Bytes { return blk.baseBTCBlock.Hash }

func (blk *BlockDOGE) Confirmations() int { return blk.baseBTCBlock.Confirmations }

func (blk *BlockDOGE) Height() int { return blk.baseBTCBlock.Height }

func (blk *BlockDOGE) TransactionsHashes() []types.Bytes { return blk.baseBTCBlock.Transactions }

func (blk *BlockDOGE) Time() types.UnixTime { return blk.baseBTCBlock.Time }

func (blk *BlockDOGE) PreviousBlockHash() types.Bytes { return blk.baseBTCBlock.PreviousBlockHash }

func (blk *BlockDOGE) NextBlockHash() types.Bytes { return blk.baseBTCBlock.NextBlockHash }

func (blk *BlockDOGE) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &blk.baseBTCBlock) }
