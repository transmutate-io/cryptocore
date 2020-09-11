package block

import (
	"encoding/json"

	"github.com/transmutate-io/cryptocore/types"
)

var (
	_ Block                  = (*BlockBCH)(nil)
	_ TransactionsLister     = (*BlockBCH)(nil)
	_ ConfirmationCounter    = (*BlockBCH)(nil)
	_ ForwardBlockNavigator  = (*BlockBCH)(nil)
	_ BackwardBlockNavigator = (*BlockBCH)(nil)
)

type BlockBCH struct{ baseBTCBlock }

func (blk *BlockBCH) Hash() types.Bytes { return blk.baseBTCBlock.Hash }

func (blk *BlockBCH) Confirmations() int { return blk.baseBTCBlock.Confirmations }

func (blk *BlockBCH) Height() int { return blk.baseBTCBlock.Height }

func (blk *BlockBCH) TransactionsHashes() []types.Bytes { return blk.baseBTCBlock.Transactions }

func (blk *BlockBCH) Time() types.UnixTime { return blk.baseBTCBlock.Time }

func (blk *BlockBCH) PreviousBlockHash() types.Bytes { return blk.baseBTCBlock.PreviousBlockHash }

func (blk *BlockBCH) NextBlockHash() types.Bytes { return blk.baseBTCBlock.NextBlockHash }

func (blk *BlockBCH) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &blk.baseBTCBlock) }
