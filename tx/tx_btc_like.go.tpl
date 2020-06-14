package {{ .Values.package }}

import (
	"encoding/json"

	"transmutate.io/pkg/cryptocore/types"
)

type Tx{{ .Values.short }} struct{ txBTC }

func (tx *Tx{{ .Values.short }}) ID() types.Bytes { return tx.txBTC.ID }

func (tx *Tx{{ .Values.short }}) Hash() types.Bytes { return tx.txBTC.Hash }

func (tx *Tx{{ .Values.short }}) BlockHash() types.Bytes { return tx.txBTC.BlockHash }

func (tx *Tx{{ .Values.short }}) Confirmations() int { return tx.txBTC.Confirmations }

func (tx *Tx{{ .Values.short }}) BlockTime() types.UnixTime { return tx.txBTC.BlockTime }

func (tx *Tx{{ .Values.short }}) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &tx.txBTC) }

func (tx *Tx{{ .Values.short }}) UTXO() (TxUTXO, bool) { return tx, true }

func (tx *Tx{{ .Values.short }}) LockTime() types.UnixTime { return tx.txBTC.LockTime }

func (tx *Tx{{ .Values.short }}) Inputs() []Input {
	r := make([]Input, 0, len(tx.txBTC.Inputs))
	for _, i := range tx.txBTC.Inputs {
		r = append(r, &wrapInput{i})
	}
	return r
}

func (tx *Tx{{ .Values.short }}) Outputs() []Output {
	r := make([]Output, 0, len(tx.txBTC.Outputs))
	for _, i := range tx.txBTC.Outputs {
		r = append(r, &wrapOutput{i})
	}
	return r
}
