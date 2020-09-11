package block

import (
	"encoding/json"

	"github.com/transmutate-io/cryptocore/types"
)

var (
	_ Block                  = (*BlockLTC)(nil)
	_ TransactionsLister     = (*BlockLTC)(nil)
	_ ConfirmationCounter    = (*BlockLTC)(nil)
	_ ForwardBlockNavigator  = (*BlockLTC)(nil)
	_ BackwardBlockNavigator = (*BlockLTC)(nil)
)

type BlockLTC struct{ baseBTCBlock }

func (blk *BlockLTC) Hash() types.Bytes { return blk.baseBTCBlock.Hash }

func (blk *BlockLTC) Confirmations() int { return blk.baseBTCBlock.Confirmations }

func (blk *BlockLTC) Height() int { return blk.baseBTCBlock.Height }

func (blk *BlockLTC) TransactionsHashes() []types.Bytes { return blk.baseBTCBlock.Transactions }

func (blk *BlockLTC) Time() types.UnixTime { return blk.baseBTCBlock.Time }

func (blk *BlockLTC) PreviousBlockHash() types.Bytes { return blk.baseBTCBlock.PreviousBlockHash }

func (blk *BlockLTC) NextBlockHash() types.Bytes { return blk.baseBTCBlock.NextBlockHash }

func (blk *BlockLTC) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &blk.baseBTCBlock) }
