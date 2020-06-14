package tx

import "transmutate.io/pkg/cryptocore/types"

type (
	txBTC struct {
		ID            types.Bytes    `json:"txid"`
		Hash          types.Bytes    `json:"hash"`
		BlockHash     types.Bytes    `json:"blockhash"`
		Confirmations int            `json:"confirmations"`
		BlockTime     types.UnixTime `json:"blocktime"`
		LockTime      types.UnixTime `json:"locktime"`
		Inputs        []*inputBTC    `json:"vin"`
		Outputs       []*outputBTC   `json:"vout"`
	}

	inputBTC struct {
		TransactionID types.Bytes   `json:"txid"`
		N             int           `json:"vout"`
		UnlockScript  *scriptSigBTC `json:"scriptSig"`
		Sequence      int           `json:"sequence"`
		Coinbase      types.Bytes   `json:"coinbase"`
	}

	outputBTC struct {
		Value      types.Amount     `json:"value"`
		N          int              `json:"n"`
		LockScript *scriptPubKeyBTC `json:"scriptPubKey"`
	}

	scriptSigBTC struct {
		Hex types.Bytes `json:"hex"`
		Asm string      `json:"asm"`
	}

	scriptPubKeyBTC struct {
		Hex                types.Bytes `json:"hex"`
		Asm                string      `json:"asm"`
		RequiredSignatures int         `json:"reqSigs"`
		Type               string      `json:"type"`
		Addresses          []string    `json:"addresses"`
	}
)

type wrapScriptSig struct{ *scriptSigBTC }

func (w *wrapScriptSig) Bytes() types.Bytes { return w.scriptSigBTC.Hex }
func (w *wrapScriptSig) Asm() string        { return w.scriptSigBTC.Asm }

type wrapInput struct{ *inputBTC }

func (w *wrapInput) TransactionID() types.Bytes { return w.inputBTC.TransactionID }
func (w *wrapInput) N() int                     { return w.inputBTC.N }
func (w *wrapInput) UnlockScript() ScriptSig    { return &wrapScriptSig{w.inputBTC.UnlockScript} }
func (w *wrapInput) Sequence() int              { return w.inputBTC.Sequence }
func (w *wrapInput) Coinbase() types.Bytes      { return w.inputBTC.Coinbase }

type wrapScriptPubKey struct{ *scriptPubKeyBTC }

func (w *wrapScriptPubKey) Bytes() types.Bytes      { return w.scriptPubKeyBTC.Hex }
func (w *wrapScriptPubKey) Asm() string             { return w.scriptPubKeyBTC.Asm }
func (w *wrapScriptPubKey) RequiredSignatures() int { return w.scriptPubKeyBTC.RequiredSignatures }
func (w *wrapScriptPubKey) Type() string            { return w.scriptPubKeyBTC.Type }
func (w *wrapScriptPubKey) Addresses() []string     { return w.scriptPubKeyBTC.Addresses }

type wrapOutput struct{ *outputBTC }

func (w *wrapOutput) Value() types.Amount      { return w.outputBTC.Value }
func (w *wrapOutput) N() int                   { return w.outputBTC.N }
func (w *wrapOutput) LockScript() ScriptPubKey { return &wrapScriptPubKey{w.outputBTC.LockScript} }
