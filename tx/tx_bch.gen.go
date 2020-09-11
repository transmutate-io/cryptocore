package tx

import (
	"encoding/json"

	"github.com/transmutate-io/cryptocore/types"
)

var (
	_ Tx     = (*TxBCH)(nil)
	_ TxUTXO = (*TxBCH)(nil)
)

type TxBCH struct{ txBTC }

func (tx *TxBCH) ID() types.Bytes { return tx.txBTC.ID }

func (tx *TxBCH) Hash() types.Bytes { return tx.txBTC.Hash }

func (tx *TxBCH) BlockHash() types.Bytes { return tx.txBTC.BlockHash }

func (tx *TxBCH) Confirmations() int { return tx.txBTC.Confirmations }

func (tx *TxBCH) BlockTime() types.UnixTime { return tx.txBTC.BlockTime }

func (tx *TxBCH) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &tx.txBTC) }

func (tx *TxBCH) UTXO() (TxUTXO, bool) { return tx, true }

func (tx *TxBCH) LockTime() types.UnixTime { return tx.txBTC.LockTime }

func (tx *TxBCH) Inputs() []Input {
	r := make([]Input, 0, len(tx.txBTC.Inputs))
	for _, i := range tx.txBTC.Inputs {
		r = append(r, &wrapInput{i})
	}
	return r
}

func (tx *TxBCH) Outputs() []Output {
	r := make([]Output, 0, len(tx.txBTC.Outputs))
	for _, i := range tx.txBTC.Outputs {
		r = append(r, &wrapOutput{i})
	}
	return r
}
