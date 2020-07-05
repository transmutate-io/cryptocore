package tx

import (
	"encoding/json"

	"github.com/transmutate-io/cryptocore/types"
)

type TxDCR struct{ txBTC }

func (tx *TxDCR) ID() types.Bytes { return tx.txBTC.ID }

func (tx *TxDCR) Hash() types.Bytes { return tx.txBTC.Hash }

func (tx *TxDCR) BlockHash() types.Bytes { return tx.txBTC.BlockHash }

func (tx *TxDCR) Confirmations() int { return tx.txBTC.Confirmations }

func (tx *TxDCR) BlockTime() types.UnixTime { return tx.txBTC.BlockTime }

func (tx *TxDCR) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &tx.txBTC) }

func (tx *TxDCR) UTXO() (TxUTXO, bool) { return tx, true }

func (tx *TxDCR) LockTime() types.UnixTime { return tx.txBTC.LockTime }

func (tx *TxDCR) Inputs() []Input {
	r := make([]Input, 0, len(tx.txBTC.Inputs))
	for _, i := range tx.txBTC.Inputs {
		r = append(r, &wrapInput{i})
	}
	return r
}

func (tx *TxDCR) Outputs() []Output {
	r := make([]Output, 0, len(tx.txBTC.Outputs))
	for _, i := range tx.txBTC.Outputs {
		r = append(r, &wrapOutput{i})
	}
	return r
}
