package types

type Transaction struct {
	ID            Bytes     `json:"txid"`
	Hash          Bytes     `json:"hash"`
	BlockHash     Bytes     `json:"blockhash"`
	Confirmations int       `json:"confirmations"`
	BlockTime     UnixTime  `json:"blocktime"`
	LockTime      UnixTime  `json:"locktime"`
	Inputs        []*Input  `json:"vin"`
	Outputs       []*Output `json:"vout"`
}

type ScriptSig struct {
	Hex Bytes  `json:"hex"`
	Asm string `json:"asm"`
}

type Input struct {
	TransactionID Bytes      `json:"txid"`
	N             int        `json:"vout"`
	LockScript    *ScriptSig `json:"scriptSig"`
	Sequence      int        `json:"sequence"`
	Coinbase      Bytes      `json:"coinbase"`
}

type Output struct {
	Value        Amount        `json:"value"`
	N            int           `json:"n"`
	UnlockScript *ScriptPubKey `json:"scriptPubKey"`
}

type ScriptPubKey struct {
	Hex                Bytes    `json:"hex"`
	Asm                string   `json:"asm"`
	RequiredSignatures int      `json:"reqSigs"`
	Type               string   `json:"type"`
	Addresses          []string `json:"addresses"`
}
