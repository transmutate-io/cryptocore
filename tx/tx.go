package tx

import "github.com/transmutate-io/cryptocore/types"

type (
	Tx interface {
		ID() types.Bytes
		Hash() types.Bytes
	}

	TxUTXO interface {
		BlockHash() types.Bytes
		Confirmations() int
		BlockTime() types.UnixTime
		LockTime() types.UnixTime
		Inputs() []Input
		Outputs() []Output
	}

	Input interface {
		TransactionID() types.Bytes
		N() int
		UnlockScript() ScriptSig
		Sequence() int
		Coinbase() types.Bytes
	}

	ScriptSig interface {
		Bytes() types.Bytes
		Asm() string
	}

	Output interface {
		Value() types.Amount
		N() int
		LockScript() ScriptPubKey
	}

	ScriptPubKey interface {
		Bytes() types.Bytes
		Asm() string
		RequiredSignatures() int
		Type() string
		Addresses() []string
	}

	TxStateBased interface {
		Value() types.Amount
		To() string
	}
)
