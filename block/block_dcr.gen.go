package block

import (
	"encoding/json"

	"github.com/transmutate-io/cryptocore/types"
)

var (
	_ Block                  = (*BlockDCR)(nil)
	_ TransactionsLister     = (*BlockDCR)(nil)
	_ ConfirmationCounter    = (*BlockDCR)(nil)
	_ ForwardBlockNavigator  = (*BlockDCR)(nil)
	_ BackwardBlockNavigator = (*BlockDCR)(nil)
)

type BlockDCR struct{ baseBTCBlock }

func (blk *BlockDCR) Hash() types.Bytes { return blk.baseBTCBlock.Hash }

func (blk *BlockDCR) Confirmations() int { return blk.baseBTCBlock.Confirmations }

func (blk *BlockDCR) Height() int { return blk.baseBTCBlock.Height }

func (blk *BlockDCR) TransactionsHashes() []types.Bytes { return blk.baseBTCBlock.Transactions }

func (blk *BlockDCR) Time() types.UnixTime { return blk.baseBTCBlock.Time }

func (blk *BlockDCR) PreviousBlockHash() types.Bytes { return blk.baseBTCBlock.PreviousBlockHash }

func (blk *BlockDCR) NextBlockHash() types.Bytes { return blk.baseBTCBlock.NextBlockHash }

func (blk *BlockDCR) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &blk.baseBTCBlock) }
