package btc

// type RawTransaction struct {
// 	RawTransactionInfo
// 	BlockHash     types.HexBytes `json:"blockhash"`
// 	Confirmations int            `json:"confirmations"`
// 	BlockTime     types.UnixTime `json:"blocktime"`
// 	Time          types.UnixTime `json:"time"`
// }

// type RawTransactionInfo struct {
// 	ID          types.HexBytes `json:"txid"`
// 	Hash        types.HexBytes `json:"hash"`
// 	Size        int            `json:"size"`
// 	VirtualSize int            `json:"vsize"`
// 	Weight      int            `json:"weight"`
// 	Version     int            `json:"version"`
// 	LockTime    types.UnixTime `json:"locktime"`
// 	VIn         []*VIn         `json:"vin"`
// 	VOut        []*VOut        `json:"vout"`
// 	Hex         types.HexBytes `json:"hex"`
// }

// type ScriptSig struct {
// 	Hex types.HexBytes `json:"hex"`
// 	Asm string         `json:"asm"`
// }

// type VIn struct {
// 	TransactionID      types.HexBytes   `json:"txid"`
// 	VOut               int              `json:"vout"`
// 	ScriptSig          *ScriptSig       `json:"scriptSig"`
// 	TransactionWitness []types.HexBytes `json:"txinwitness"`
// 	Sequence           int              `json:"sequence"`
// 	Coinbase           types.HexBytes   `json:"coinbase"`
// }

// type ScriptPubKey struct {
// 	Hex                types.HexBytes `json:"hex"`
// 	Asm                string         `json:"asm"`
// 	RequiredSignatures int            `json:"reqSigs"`
// 	Type               string         `json:"type"`
// 	Addresses          []string       `json:"addresses"`
// }

// type VOut struct {
// 	Value        Amount        `json:"value"`
// 	N            int           `json:"n"`
// 	ScriptPubKey *ScriptPubKey `json:"scriptPubKey"`
// }

// type UnspentOutput struct {
// 	TxID          string `json:"txid"`
// 	VOut          int    `json:"vout"`
// 	Address       string `json:"address"`
// 	ScriptPubKey  string `json:"scriptPubKey"`
// 	Amount        Amount `json:"amount"`
// 	Confirmations int    `json:"confirmations"`
// 	RedeemScript  string `json:"redeemScript"`
// 	WitnessScript string `json:"witnessScript"`
// 	Spendable     bool   `json:"spendable"`
// 	Solvable      bool   `json:"solvable"`
// 	Reused        bool   `json:"reused"`
// 	Description   string `json:"desc"`
// 	Safe          bool   `json:"safe"`
// }

// type Block struct {
// 	Hash              types.HexBytes    `json:"hash"`
// 	Confirmations     int               `json:"confirmations"`
// 	Size              int               `json:"size"`
// 	StrippedSize      int               `json:"strippedsize"`
// 	Weight            int               `json:"weight"`
// 	Height            int               `json:"height"`
// 	Version           int               `json:"version"`
// 	VersionHex        types.HexBytes    `json:"versionHex"`
// 	MerkleRoot        types.HexBytes    `json:"merkleroot"`
// 	Transactions      BlockTransactions `json:"tx"`
// 	Time              types.UnixTime    `json:"time"`
// 	MedianTime        types.UnixTime    `json:"mediantime"`
// 	Nonce             int               `json:"nonce"`
// 	Bits              types.HexBytes    `json:"bits"`
// 	Difficulty        float64           `json:"difficulty"`
// 	ChainWork         types.HexBytes    `json:"chainwork"`
// 	PreviousBlockHash types.HexBytes    `json:"previousblockhash"`
// 	NextBlockHash     types.HexBytes    `json:"nextblockhash"`
// }

// type BlockTransactions []interface{}

// var ErrInvalidTransactionData = errors.New("invalid transaction data")

// func (bt BlockTransactions) MarshalJSON() ([]byte, error) {
// 	return json.Marshal([]interface{}(bt))
// }

// func (bt *BlockTransactions) UnmarshalJSON(b []byte) error {
// 	var (
// 		getTx func(int) interface{}
// 		sz    int
// 	)
// 	var (
// 		hashesOnly []types.HexBytes
// 		rawTxs     []*RawTransactionInfo
// 	)
// 	err := json.Unmarshal(b, &hashesOnly)
// 	if err == nil {
// 		sz = len(hashesOnly)
// 		getTx = func(i int) interface{} { return hashesOnly[i] }
// 	} else if err = json.Unmarshal(b, &rawTxs); err == nil {
// 		sz = len(rawTxs)
// 		getTx = func(i int) interface{} { return rawTxs[i] }
// 	} else {
// 		return ErrInvalidTransactionData
// 	}
// 	*bt = make(BlockTransactions, 0, sz)
// 	for i := 0; i < sz; i++ {
// 		*bt = append(*bt, getTx(i))
// 	}
// 	return nil
// }

// func (bt BlockTransactions) IsHashes() bool {
// 	if len(bt) < 0 {
// 		return false
// 	}
// 	_, ok := bt[0].(types.HexBytes)
// 	return ok
// }

// func (bt BlockTransactions) IsRaw() bool {
// 	if len(bt) < 0 {
// 		return false
// 	}
// 	_, ok := bt[0].(*RawTransactionInfo)
// 	return ok
// }

// func (bt BlockTransactions) Hashes() []types.HexBytes {
// 	r := make([]types.HexBytes, 0, len(bt))
// 	if bt.IsHashes() {
// 		for _, t := range bt {
// 			r = append(r, t.(types.HexBytes))
// 		}
// 	} else if bt.IsRaw() {
// 		for _, t := range bt {
// 			r = append(r, t.(*RawTransactionInfo).Hash)
// 		}
// 	}
// 	return r
// }

// func (bt BlockTransactions) Raw() []*RawTransactionInfo {
// 	if !bt.IsRaw() {
// 		panic("not raw tx info")
// 	}
// 	r := make([]*RawTransactionInfo, 0, len(bt))
// 	for _, t := range bt {
// 		r = append(r, t.(*RawTransactionInfo))
// 	}
// 	return r
// }

// type AddressFunds struct {
// 	WatchOnly     bool     `json:"involvesWatchonly"`
// 	Address       string   `json:"address"`
// 	Amount        Amount   `json:"amount"`
// 	Confirmations int      `json:"confirmations"`
// 	Label         string   `json:"label"`
// 	TxIDs         []string `json:"txids"`
// }
