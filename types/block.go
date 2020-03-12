package types

type Block struct {
	Hash              Bytes    `json:"hash"`
	Confirmations     int      `json:"confirmations"`
	Height            int      `json:"height"`
	Transactions      []Bytes  `json:"tx"`
	Time              UnixTime `json:"time"`
	PreviousBlockHash Bytes    `json:"previousblockhash"`
	NextBlockHash     Bytes    `json:"nextblockhash"`
}
