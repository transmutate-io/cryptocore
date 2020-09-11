package {{ .Values.package }}

import (
	"encoding/json"

	"github.com/transmutate-io/cryptocore/types"
)

var (
	_ Block                  = (*Block{{ .Values.short }})(nil)
	_ TransactionsLister     = (*Block{{ .Values.short }})(nil)
	_ ConfirmationCounter    = (*Block{{ .Values.short }})(nil)
	_ ForwardBlockNavigator  = (*Block{{ .Values.short }})(nil)
	_ BackwardBlockNavigator = (*Block{{ .Values.short }})(nil)
)

type Block{{ .Values.short }} struct{ baseBTCBlock }

func (blk *Block{{ .Values.short }}) Hash() types.Bytes { return blk.baseBTCBlock.Hash }

func (blk *Block{{ .Values.short }}) Confirmations() int { return blk.baseBTCBlock.Confirmations }

func (blk *Block{{ .Values.short }}) Height() int { return blk.baseBTCBlock.Height }

func (blk *Block{{ .Values.short }}) TransactionsHashes() []types.Bytes { return blk.baseBTCBlock.Transactions }

func (blk *Block{{ .Values.short }}) Time() types.UnixTime { return blk.baseBTCBlock.Time }

func (blk *Block{{ .Values.short }}) PreviousBlockHash() types.Bytes { return blk.baseBTCBlock.PreviousBlockHash }

func (blk *Block{{ .Values.short }}) NextBlockHash() types.Bytes { return blk.baseBTCBlock.NextBlockHash }

func (blk *Block{{ .Values.short }}) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &blk.baseBTCBlock) }
