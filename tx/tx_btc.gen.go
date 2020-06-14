package tx

import (
	"encoding/json"

	"transmutate.io/pkg/cryptocore/types"
)

type TxBTC struct{ txBTC }

func (tx *TxBTC) ID() types.Bytes { return tx.txBTC.ID }

func (tx *TxBTC) Hash() types.Bytes { return tx.txBTC.Hash }

func (tx *TxBTC) BlockHash() types.Bytes { return tx.txBTC.BlockHash }

func (tx *TxBTC) Confirmations() int { return tx.txBTC.Confirmations }

func (tx *TxBTC) BlockTime() types.UnixTime { return tx.txBTC.BlockTime }

func (tx *TxBTC) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &tx.txBTC) }

func (tx *TxBTC) UTXO() (TxUTXO, bool) { return tx, true }

func (tx *TxBTC) LockTime() types.UnixTime { return tx.txBTC.LockTime }

func (tx *TxBTC) Inputs() []Input {
	r := make([]Input, 0, len(tx.txBTC.Inputs))
	for _, i := range tx.txBTC.Inputs {
		r = append(r, &wrapInput{i})
	}
	return r
}

func (tx *TxBTC) Outputs() []Output {
	r := make([]Output, 0, len(tx.txBTC.Outputs))
	for _, i := range tx.txBTC.Outputs {
		r = append(r, &wrapOutput{i})
	}
	return r
}
