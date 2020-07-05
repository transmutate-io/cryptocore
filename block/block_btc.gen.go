package block

import (
	"encoding/json"

	"github.com/transmutate-io/cryptocore/types"
)

type BlockBTC struct{ commonBlock }

func (blk *BlockBTC) Hash() types.Bytes { return blk.commonBlock.Hash }

func (blk *BlockBTC) Confirmations() int { return blk.commonBlock.Confirmations }

func (blk *BlockBTC) Height() int { return blk.commonBlock.Height }

func (blk *BlockBTC) Transactions() []types.Bytes { return blk.commonBlock.Transactions }

func (blk *BlockBTC) Time() types.UnixTime { return blk.commonBlock.Time }

func (blk *BlockBTC) PreviousBlockHash() types.Bytes { return blk.commonBlock.PreviousBlockHash }

func (blk *BlockBTC) NextBlockHash() types.Bytes { return blk.commonBlock.NextBlockHash }

func (blk *BlockBTC) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &blk.commonBlock) }
