package block

import (
	"encoding/json"

	"github.com/transmutate-io/cryptocore/types"
)

type BlockDCR struct{ commonBlock }

func (blk *BlockDCR) Hash() types.Bytes { return blk.commonBlock.Hash }

func (blk *BlockDCR) Confirmations() int { return blk.commonBlock.Confirmations }

func (blk *BlockDCR) Height() int { return blk.commonBlock.Height }

func (blk *BlockDCR) Transactions() []types.Bytes { return blk.commonBlock.Transactions }

func (blk *BlockDCR) Time() types.UnixTime { return blk.commonBlock.Time }

func (blk *BlockDCR) PreviousBlockHash() types.Bytes { return blk.commonBlock.PreviousBlockHash }

func (blk *BlockDCR) NextBlockHash() types.Bytes { return blk.commonBlock.NextBlockHash }

func (blk *BlockDCR) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &blk.commonBlock) }
