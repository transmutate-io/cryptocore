package block

import (
	"encoding/json"

	"github.com/transmutate-io/cryptocore/types"
)

type BlockBCH struct{ commonBlock }

func (blk *BlockBCH) Hash() types.Bytes { return blk.commonBlock.Hash }

func (blk *BlockBCH) Confirmations() int { return blk.commonBlock.Confirmations }

func (blk *BlockBCH) Height() int { return blk.commonBlock.Height }

func (blk *BlockBCH) Transactions() []types.Bytes { return blk.commonBlock.Transactions }

func (blk *BlockBCH) Time() types.UnixTime { return blk.commonBlock.Time }

func (blk *BlockBCH) PreviousBlockHash() types.Bytes { return blk.commonBlock.PreviousBlockHash }

func (blk *BlockBCH) NextBlockHash() types.Bytes { return blk.commonBlock.NextBlockHash }

func (blk *BlockBCH) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &blk.commonBlock) }
