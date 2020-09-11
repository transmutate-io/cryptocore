package tx

import (
	"encoding/json"

	"github.com/transmutate-io/cryptocore/types"
)

var (
	_ Tx     = (*TxDOGE)(nil)
	_ TxUTXO = (*TxDOGE)(nil)
)

type TxDOGE struct{ txBTC }

func (tx *TxDOGE) ID() types.Bytes { return tx.txBTC.ID }

func (tx *TxDOGE) Hash() types.Bytes { return tx.txBTC.Hash }

func (tx *TxDOGE) BlockHash() types.Bytes { return tx.txBTC.BlockHash }

func (tx *TxDOGE) Confirmations() int { return tx.txBTC.Confirmations }

func (tx *TxDOGE) BlockTime() types.UnixTime { return tx.txBTC.BlockTime }

func (tx *TxDOGE) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &tx.txBTC) }

func (tx *TxDOGE) UTXO() (TxUTXO, bool) { return tx, true }

func (tx *TxDOGE) LockTime() types.UnixTime { return tx.txBTC.LockTime }

func (tx *TxDOGE) Inputs() []Input {
	r := make([]Input, 0, len(tx.txBTC.Inputs))
	for _, i := range tx.txBTC.Inputs {
		r = append(r, &wrapInput{i})
	}
	return r
}

func (tx *TxDOGE) Outputs() []Output {
	r := make([]Output, 0, len(tx.txBTC.Outputs))
	for _, i := range tx.txBTC.Outputs {
		r = append(r, &wrapOutput{i})
	}
	return r
}
