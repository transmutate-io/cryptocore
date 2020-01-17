package btccore

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
)

type HexBytes []byte

func (hb HexBytes) String() string               { return hex.EncodeToString(hb) }
func (hb HexBytes) MarshalJSON() ([]byte, error) { return json.Marshal(hb.String()) }

func (hb *HexBytes) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	*hb = b
	return nil
}

type UnixTime int

func NewUnixTime(t time.Time) UnixTime           { return UnixTime(t.UTC().Unix()) }
func (ut UnixTime) Time() time.Time              { return time.Unix(int64(ut), 0).UTC() }
func (ut UnixTime) String() string               { return ut.Time().Format(time.RFC3339) }
func (ut UnixTime) MarshalJSON() ([]byte, error) { return json.Marshal(int(ut)) }

func (ut *UnixTime) UnmarshalJSON(b []byte) error {
	var t int
	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}
	*ut = UnixTime(t)
	return nil
}

type Block struct {
	Hash              HexBytes          `json:"hash"`
	Confirmations     int               `json:"confirmations"`
	Size              int               `json:"size"`
	StrippedSize      int               `json:"strippedsize"`
	Weight            int               `json:"weight"`
	Height            int               `json:"height"`
	Version           int               `json:"version"`
	VersionHex        HexBytes          `json:"versionHex"`
	MerkleRoot        HexBytes          `json:"merkleroot"`
	Transactions      BlockTransactions `json:"tx"`
	Time              UnixTime          `json:"time"`
	MedianTime        UnixTime          `json:"mediantime"`
	Nonce             int               `json:"nonce"`
	Bits              HexBytes          `json:"bits"`
	Difficulty        float64           `json:"difficulty"`
	ChainWork         HexBytes          `json:"chainwork"`
	PreviousBlockHash HexBytes          `json:"previousblockhash"`
	NextBlockHash     HexBytes          `json:"nextblockhash"`
}

type BlockTransactions []interface{}

var ErrInvalidTransactionData = errors.New("invalid transaction data")

func (bt BlockTransactions) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}(bt))
}

func (bt *BlockTransactions) UnmarshalJSON(b []byte) error {
	var (
		getTx func(int) interface{}
		sz    int
	)
	var (
		hashesOnly []HexBytes
		rawTxs     []*RawTransactionInfo
	)
	err := json.Unmarshal(b, &hashesOnly)
	if err == nil {
		sz = len(hashesOnly)
		getTx = func(i int) interface{} { return hashesOnly[i] }
	} else if err = json.Unmarshal(b, &rawTxs); err == nil {
		sz = len(rawTxs)
		getTx = func(i int) interface{} { return rawTxs[i] }
	} else {
		return ErrInvalidTransactionData
	}
	*bt = make(BlockTransactions, 0, sz)
	for i := 0; i < sz; i++ {
		*bt = append(*bt, getTx(i))
	}
	return nil
}

func (bt BlockTransactions) IsHashes() bool {
	if len(bt) < 0 {
		return false
	}
	_, ok := bt[0].(HexBytes)
	return ok
}

func (bt BlockTransactions) IsRaw() bool {
	if len(bt) < 0 {
		return false
	}
	_, ok := bt[0].(*RawTransactionInfo)
	return ok
}

func (bt BlockTransactions) Hashes() []HexBytes {
	r := make([]HexBytes, 0, len(bt))
	if bt.IsHashes() {
		for _, t := range bt {
			r = append(r, t.(HexBytes))
		}
	} else if bt.IsRaw() {
		for _, t := range bt {
			r = append(r, t.(*RawTransactionInfo).Hash)
		}
	}
	return r
}

func (bt BlockTransactions) Raw() []*RawTransactionInfo {
	if !bt.IsRaw() {
		panic("not raw tx info")
	}
	r := make([]*RawTransactionInfo, 0, len(bt))
	for _, t := range bt {
		r = append(r, t.(*RawTransactionInfo))
	}
	return r
}

type RawTransaction struct {
	RawTransactionInfo
	BlockHash     HexBytes `json:"blockhash"`
	Confirmations int      `json:"confirmations"`
	BlockTime     UnixTime `json:"blocktime"`
	Time          UnixTime `json:"time"`
}

type RawTransactionInfo struct {
	ID          HexBytes `json:"txid"`
	Hash        HexBytes `json:"hash"`
	Size        int      `json:"size"`
	VirtualSize int      `json:"vsize"`
	Weight      int      `json:"weight"`
	Version     int      `json:"version"`
	LockTime    UnixTime `json:"locktime"`
	VIn         []*VIn   `json:"vin"`
	VOut        []*VOut  `json:"vout"`
	Hex         HexBytes `json:"hex"`
}

type ScriptSig struct {
	Hex HexBytes `json:"hex"`
	Asm string   `json:"asm"`
}

type VIn struct {
	TransactionID      HexBytes   `json:"txid"`
	VOut               int        `json:"vout"`
	ScriptSig          *ScriptSig `json:"scriptSig"`
	TransactionWitness []HexBytes `json:"txinwitness"`
	Sequence           int        `json:"sequence"`
	Coinbase           HexBytes   `json:"coinbase"`
}

type ScriptPubKey struct {
	Hex                HexBytes `json:"hex"`
	Asm                string   `json:"asm"`
	RequiredSignatures int      `json:"reqSigs"`
	Type               string   `json:"type"`
	Addresses          []string `json:"addresses"`
}

type VOut struct {
	Value        *Amount       `json:"value"`
	N            int           `json:"n"`
	ScriptPubKey *ScriptPubKey `json:"scriptPubKey"`
}

type Amount big.Int

var oneBTC = big.NewInt(100000000)

func splitSats(sats *big.Int) (*big.Int, *big.Int) {
	b, s := new(big.Int).Set(sats), new(big.Int)
	b.DivMod(b, oneBTC, s)
	return b, s
}

func (a Amount) String() string {
	bi := big.Int(a)
	b, s := splitSats(&bi)
	r := fmt.Sprintf("%s.%.8d", b.String(), s.Int64())
	return strings.TrimRight(r, "0.")
}

func (a Amount) MarshalJSON() ([]byte, error) { return []byte(a.String()), nil }

func newNumError(funcName, num string) *strconv.NumError {
	return &strconv.NumError{
		Func: funcName,
		Num:  num,
		Err:  strconv.ErrSyntax,
	}
}

func (a *Amount) UnmarshalJSON(b []byte) error {
	parts := bytes.SplitN(b, []byte{'.'}, 2)
	part := string(parts[0])
	r, ok := new(big.Int).SetString(part, 10)
	if !ok {
		return newNumError("Amount.UnmarshalJSON", part)
	}
	r.Mul(r, oneBTC)
	if len(parts) > 1 {
		for len(parts[1]) < 8 {
			parts[1] = append(parts[1], '0')
		}
		part = string(parts[1])
		rs, ok := new(big.Int).SetString(part, 10)
		if !ok {
			return newNumError("Amount.UnmarshalJSON", part)
		}
		r.Add(r, rs)
	}
	(*big.Int)(a).Set(r)
	return nil
}

type UnspentOutput struct {
	TxID          string  `json:"txid"`
	VOut          int     `json:"vout"`
	Address       string  `json:"address"`
	ScriptPubKey  string  `json:"scriptPubKey"`
	Amount        *Amount `json:"amount"`
	Confirmations int     `json:"confirmations"`
	RedeemScript  string  `json:"redeemScript"`
	WitnessScript string  `json:"witnessScript"`
	Spendable     bool    `json:"spendable"`
	Solvable      bool    `json:"solvable"`
	Reused        bool    `json:"reused"`
	Description   string  `json:"desc"`
	Safe          bool    `json:"safe"`
}
