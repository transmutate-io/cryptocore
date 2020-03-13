package btccore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"transmutate.io/pkg/btccore/types"
)

type baseClient struct {
	Address  string
	Username string
	Password string
	UseHTTPS bool
}

func newBaseClient(addr, user, pass string, useHTTPS bool) *baseClient {
	return &baseClient{
		Address:  addr,
		Username: user,
		Password: pass,
		UseHTTPS: useHTTPS,
	}
}

func (c *baseClient) do(method string, params interface{}, r interface{}) error {
	b := bytes.NewBuffer(make([]byte, 0, 1024))
	err := json.NewEncoder(b).Encode(&rpcRequest{
		JsonRPC: "1.0",
		ID:      "go-btccore",
		Method:  method,
		Params:  params,
	})
	if err != nil {
		return err
	}
	var useHTTPS string
	if c.UseHTTPS {
		useHTTPS = "s"
	}
	resp, err := http.Post(fmt.Sprintf("http%s://%s:%s@%s/", useHTTPS, c.Username, c.Password, c.Address), "application/json", b)
	if err != nil {
		return err
	}
	rr := &rpcResponse{Result: r}
	if err = json.NewDecoder(resp.Body).Decode(rr); err != nil {
		return err
	}
	if rr.Error != nil {
		return rr.Error
	}
	return nil
}

func (c *baseClient) doString(method string, args interface{}) (string, error) {
	var r string
	if err := c.do(method, args, &r); err != nil {
		return "", err
	}
	return r, nil
}

func (c *baseClient) doUint64(method string, args interface{}) (uint64, error) {
	var r uint64
	if err := c.do(method, args, &r); err != nil {
		return 0, err
	}
	return r, nil
}

func (c *baseClient) doBytes(method string, args interface{}) (types.Bytes, error) {
	var r types.Bytes
	if err := c.do(method, args, &r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *baseClient) doSliceBytes(method string, args interface{}) ([]types.Bytes, error) {
	var r []types.Bytes
	if err := c.do(method, args, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type rpcRequest struct {
	JsonRPC string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type rpcResponse struct {
	ID     string       `json:"id"`
	Error  *ClientError `json:"error"`
	Result interface{}  `json:"result"`
}

type ClientError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (we *ClientError) Error() string { return we.Message }

func (c *baseClient) NewAddress() (string, error) { return c.doString("getnewaddress", nil) }

func mkArgs(n int, args ...interface{}) []interface{} {
	var sz int
	if argsSz := len(args); n > argsSz {
		sz = n
	} else {
		sz = argsSz
	}
	return append(make([]interface{}, 0, sz), args...)
}

func args(a ...interface{}) []interface{} { return a }

func (c *baseClient) DumpPrivateKey(addr string) (string, error) {
	return c.doString("dumpprivkey", args(addr))
}

func (c *baseClient) GenerateToAddress(nBlocks int, addr string) ([]types.Bytes, error) {
	return c.doSliceBytes("generatetoaddress", args(nBlocks, addr))
}

func (c *baseClient) SendToAddress(addr string, value types.Amount) (types.Bytes, error) {
	return c.doBytes("sendtoaddress", args(addr, value))
}

func (c *baseClient) BlockCount() (uint64, error) { return c.doUint64("getblockcount", nil) }

func (c *baseClient) BlockHash(height uint64) (types.Bytes, error) {
	return c.doBytes("getblockhash", args(height))
}

func (c *baseClient) SendRawTransaction(tx types.Bytes) (types.Bytes, error) {
	return c.doBytes("sendrawtransaction", args(tx.Hex()))
}

func (c *baseClient) RawTransaction(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getrawtransaction", args(hash.Hex(), false))
}

func (c *baseClient) Transaction(hash types.Bytes) (*types.Transaction, error) {
	r := &types.Transaction{}
	if err := c.do("getrawtransaction", args(hash, true), r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *baseClient) Balance(minConf int64) (types.Amount, error) {
	var (
		a = mkArgs(2, "*")
		r types.Amount
	)
	if minConf >= 0 {
		a = append(a, minConf)
	}
	if err := c.do("getbalance", a, &r); err != nil {
		return "", err
	}
	return r, nil
}

func (c *baseClient) block(args ...interface{}) (*types.Block, error) {
	r := &types.Block{}
	if err := c.do("getblock", args, r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *baseClient) ReceivedByAddress(minConf, includeEmpty, includeWatchOnly interface{}) ([]*types.AddressFunds, error) {
	args := mkArgs(3)
	if minConf != nil {
		args = append(args, minConf)
	}
	if includeEmpty != nil {
		args = append(args, includeEmpty)
	}
	if includeWatchOnly != nil {
		args = append(args, includeWatchOnly)
	}
	r := []*types.AddressFunds{}
	if err := c.do("listreceivedbyaddress", args, &r); err != nil {
		return nil, err
	}
	return r, nil
}
