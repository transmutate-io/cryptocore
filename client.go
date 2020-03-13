package btccore

import (
	"transmutate.io/pkg/btccore/types"
)

type Client interface {
	do(method string, params interface{}, r interface{}) error
	NewAddress() (string, error)
	DumpPrivateKey(addr string) (string, error)
	GenerateToAddress(nBlocks int, addr string) ([]types.Bytes, error)
	SendToAddress(addr string, value types.Amount) (types.Bytes, error)
	Balance(minConf int64) (types.Amount, error)
	BlockCount() (uint64, error)
	BlockHash(height uint64) (types.Bytes, error)
	RawBlock(hash types.Bytes) (types.Bytes, error)
	Block(hash types.Bytes) (*types.Block, error)
	SendRawTransaction(tx types.Bytes) (types.Bytes, error)
	RawTransaction(hash types.Bytes) (types.Bytes, error)
	Transaction(hash types.Bytes) (*types.Transaction, error)
	// Unspent(minConf, maxConf int, addrs []string) ([]*btc.UnspentOutput, error)
	ReceivedByAddress(minConf, includeEmpty, includeWatchOnly interface{}) ([]*types.AddressFunds, error)
}

// func (c *baseClient) ListUnspent(minConf, maxConf int, addrs []string) ([]*btc.UnspentOutput, error) {
// 	var r []*btc.UnspentOutput
// 	args := make(map[string]interface{}, 3)
// 	if minConf > -1 {
// 		args["minconf"] = minConf
// 	}
// 	if maxConf > -1 {
// 		args["maxconf"] = maxConf
// 	}
// 	if addrs != nil {
// 		args["addresses"] = addrs
// 	}
// 	if err := c.Do("listunspent", args, &r); err != nil {
// 		return nil, err
// 	}
// 	return r, nil
// }

// func (c *baseClient) ListReceivedByAddress(minConf, includeEmpty, includeWatchOnly, addrFilter interface{}) ([]*btc.AddressFunds, error) {
// 	r := make([]*btc.AddressFunds, 0, 64)
// 	args := make(map[string]interface{}, 4)
// 	if minConf != nil {
// 		args["minconf"] = minConf
// 	}
// 	if includeEmpty != nil {
// 		args["include_empty"] = includeEmpty
// 	}
// 	if includeWatchOnly != nil {
// 		args["include_watchonly"] = includeWatchOnly
// 	}
// 	if addrFilter != nil {
// 		args["address_filter"] = addrFilter
// 	}
// 	if err := c.Do("listreceivedbyaddress", args, &r); err != nil {
// 		return nil, err
// 	}
// 	return r, nil
// }
