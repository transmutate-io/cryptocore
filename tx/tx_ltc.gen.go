package tx

import (
	"encoding/json"

	"github.com/transmutate-io/cryptocore/types"
)

type TxLTC struct{ txBTC }

func (tx *TxLTC) ID() types.Bytes { return tx.txBTC.ID }

func (tx *TxLTC) Hash() types.Bytes { return tx.txBTC.Hash }

func (tx *TxLTC) BlockHash() types.Bytes { return tx.txBTC.BlockHash }

func (tx *TxLTC) Confirmations() int { return tx.txBTC.Confirmations }

func (tx *TxLTC) BlockTime() types.UnixTime { return tx.txBTC.BlockTime }

func (tx *TxLTC) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &tx.txBTC) }

func (tx *TxLTC) UTXO() (TxUTXO, bool) { return tx, true }

func (tx *TxLTC) LockTime() types.UnixTime { return tx.txBTC.LockTime }

func (tx *TxLTC) Inputs() []Input {
	r := make([]Input, 0, len(tx.txBTC.Inputs))
	for _, i := range tx.txBTC.Inputs {
		r = append(r, &wrapInput{i})
	}
	return r
}

func (tx *TxLTC) Outputs() []Output {
	r := make([]Output, 0, len(tx.txBTC.Outputs))
	for _, i := range tx.txBTC.Outputs {
		r = append(r, &wrapOutput{i})
	}
	return r
}
