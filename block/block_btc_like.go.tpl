package {{ .Values.package }}

import (
	"encoding/json"

	"transmutate.io/pkg/cryptocore/types"
)

type Block{{ .Values.short }} struct{ commonBlock }

func (blk *Block{{ .Values.short }}) Hash() types.Bytes { return blk.commonBlock.Hash }

func (blk *Block{{ .Values.short }}) Confirmations() int { return blk.commonBlock.Confirmations }

func (blk *Block{{ .Values.short }}) Height() int { return blk.commonBlock.Height }

func (blk *Block{{ .Values.short }}) Transactions() []types.Bytes { return blk.commonBlock.Transactions }

func (blk *Block{{ .Values.short }}) Time() types.UnixTime { return blk.commonBlock.Time }

func (blk *Block{{ .Values.short }}) PreviousBlockHash() types.Bytes { return blk.commonBlock.PreviousBlockHash }

func (blk *Block{{ .Values.short }}) NextBlockHash() types.Bytes { return blk.commonBlock.NextBlockHash }

func (blk *Block{{ .Values.short }}) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &blk.commonBlock) }
