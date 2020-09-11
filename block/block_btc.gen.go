package block

import (
	"encoding/json"

	"github.com/transmutate-io/cryptocore/types"
)

var (
	_ Block                  = (*BlockBTC)(nil)
	_ TransactionsLister     = (*BlockBTC)(nil)
	_ ConfirmationCounter    = (*BlockBTC)(nil)
	_ ForwardBlockNavigator  = (*BlockBTC)(nil)
	_ BackwardBlockNavigator = (*BlockBTC)(nil)
)

type BlockBTC struct{ baseBTCBlock }

func (blk *BlockBTC) Hash() types.Bytes { return blk.baseBTCBlock.Hash }

func (blk *BlockBTC) Confirmations() int { return blk.baseBTCBlock.Confirmations }

func (blk *BlockBTC) Height() int { return blk.baseBTCBlock.Height }

func (blk *BlockBTC) TransactionsHashes() []types.Bytes { return blk.baseBTCBlock.Transactions }

func (blk *BlockBTC) Time() types.UnixTime { return blk.baseBTCBlock.Time }

func (blk *BlockBTC) PreviousBlockHash() types.Bytes { return blk.baseBTCBlock.PreviousBlockHash }

func (blk *BlockBTC) NextBlockHash() types.Bytes { return blk.baseBTCBlock.NextBlockHash }

func (blk *BlockBTC) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &blk.baseBTCBlock) }
