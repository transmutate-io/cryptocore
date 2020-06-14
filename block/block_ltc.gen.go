package block

import (
	"encoding/json"

	"transmutate.io/pkg/cryptocore/types"
)

type BlockLTC struct{ commonBlock }

func (blk *BlockLTC) Hash() types.Bytes { return blk.commonBlock.Hash }

func (blk *BlockLTC) Confirmations() int { return blk.commonBlock.Confirmations }

func (blk *BlockLTC) Height() int { return blk.commonBlock.Height }

func (blk *BlockLTC) Transactions() []types.Bytes { return blk.commonBlock.Transactions }

func (blk *BlockLTC) Time() types.UnixTime { return blk.commonBlock.Time }

func (blk *BlockLTC) PreviousBlockHash() types.Bytes { return blk.commonBlock.PreviousBlockHash }

func (blk *BlockLTC) NextBlockHash() types.Bytes { return blk.commonBlock.NextBlockHash }

func (blk *BlockLTC) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &blk.commonBlock) }
