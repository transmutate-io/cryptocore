package block

import "github.com/transmutate-io/cryptocore/types"

type baseBTCBlock struct {
	Hash              types.Bytes    `json:"hash"`
	Confirmations     int            `json:"confirmations"`
	Height            int            `json:"height"`
	Transactions      []types.Bytes  `json:"tx"`
	Time              types.UnixTime `json:"time"`
	PreviousBlockHash types.Bytes    `json:"previousblockhash,omitempty"`
	NextBlockHash     types.Bytes    `json:"nextblockhash,omitempty"`
}
