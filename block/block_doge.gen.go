package block

import (
	"encoding/json"

	"transmutate.io/pkg/cryptocore/types"
)

type BlockDOGE struct{ commonBlock }

func (blk *BlockDOGE) Hash() types.Bytes { return blk.commonBlock.Hash }

func (blk *BlockDOGE) Confirmations() int { return blk.commonBlock.Confirmations }

func (blk *BlockDOGE) Height() int { return blk.commonBlock.Height }

func (blk *BlockDOGE) Transactions() []types.Bytes { return blk.commonBlock.Transactions }

func (blk *BlockDOGE) Time() types.UnixTime { return blk.commonBlock.Time }

func (blk *BlockDOGE) PreviousBlockHash() types.Bytes { return blk.commonBlock.PreviousBlockHash }

func (blk *BlockDOGE) NextBlockHash() types.Bytes { return blk.commonBlock.NextBlockHash }

func (blk *BlockDOGE) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &blk.commonBlock) }
