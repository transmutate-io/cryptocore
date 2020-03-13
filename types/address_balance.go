package types

type AddressFunds struct {
	WatchOnly     bool    `json:"involvesWatchonly"`
	Address       string  `json:"address"`
	Amount        Amount  `json:"amount"`
	Confirmations int     `json:"confirmations"`
	Label         string  `json:"label"`
	TxIDs         []Bytes `json:"txids"`
}
