package btccore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	Address  string
	Username string
	Password string
	UseHTTPS bool
}

func (c *Client) do(method string, params interface{}, r interface{}) error {
	b := bytes.NewBuffer(make([]byte, 0, 1024))
	err := json.NewEncoder(b).Encode(&rpcRequest{
		JsonRPC: "1.0",
		ID:      "go-btcorigins",
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
	err = json.NewDecoder(resp.Body).Decode(rr)
	if err != nil {
		return err
	}
	if rr.Error != nil {
		return rr.Error
	}
	return nil
}

type rpcRequest struct {
	JsonRPC string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type rpcResponse struct {
	ID     string       `json:"id"`
	Error  *walletError `json:"error"`
	Result interface{}  `json:"result"`
}

type walletError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (we *walletError) Error() string { return we.Message }

func (c *Client) GetBlockHash(n int) (HexBytes, error) {
	r := HexBytes{}
	err := c.do("getblockhash", []interface{}{n}, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) getBlock(hash HexBytes, verboseLevel int) (interface{}, error) {
	var r interface{}
	switch verboseLevel {
	case 0:
		r = &HexBytes{}
	case 1, 2:
		r = &Block{}
	default:
		panic("invalid getblock verbose level")
	}
	err := c.do("getblock", []interface{}{hash, verboseLevel}, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) GetBlockHex(hash HexBytes) (HexBytes, error) {
	r, err := c.getBlock(hash, 0)
	if err != nil {
		return nil, err
	}
	return *(r.(*HexBytes)), nil
}

func (c *Client) GetBlock(hash HexBytes, verbose bool) (*Block, error) {
	verboseLevel := 1
	if verbose {
		verboseLevel = 2
	}
	r, err := c.getBlock(hash, verboseLevel)
	if err != nil {
		return nil, err
	}
	return r.(*Block), nil
}

func (c *Client) GetRawTransactionHex(hash HexBytes) (HexBytes, error) {
	r, err := c.getRawTransaction(hash, false)
	if err != nil {
		return nil, err
	}
	return r.(HexBytes), nil
}

func (c *Client) GetRawTransaction(hash HexBytes) (*RawTransaction, error) {
	r, err := c.getRawTransaction(hash, true)
	if err != nil {
		return nil, err
	}
	return r.(*RawTransaction), nil
}

func (c *Client) getRawTransaction(hash HexBytes, decode bool) (interface{}, error) {
	var r interface{}
	if decode {
		r = &RawTransaction{}
	} else {
		r = &HexBytes{}
	}
	if err := c.do("getrawtransaction", []interface{}{hash, decode}, r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) GetBlockCount() (uint64, error) {
	var r uint64
	if err := c.do("getblockcount", nil, &r); err != nil {
		return 0, err
	}
	return r, nil
}

func (c *Client) GenerateToAddress(nBlocks int, addr string) ([]string, error) {
	r := make([]string, 0, nBlocks)
	if err := c.do("generatetoaddress", []interface{}{nBlocks, addr}, &r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) GetNewAddress() (string, error) {
	var r string
	if err := c.do("getnewaddress", nil, &r); err != nil {
		return "", err
	}
	return r, nil
}

func (c *Client) DumpPrivKey(addr string) (string, error) {
	var r string
	if err := c.do("dumpprivkey", []interface{}{addr}, &r); err != nil {
		return "", err
	}
	return r, nil
}

func (c *Client) ListUnspent(minConf, maxConf int, addrs []string) ([]*UnspentOutput, error) {
	var r []*UnspentOutput
	args := make(map[string]interface{}, 3)
	if minConf > -1 {
		args["minconf"] = minConf
	}
	if maxConf > -1 {
		args["maxconf"] = maxConf
	}
	if addrs != nil {
		args["addresses"] = addrs
	}
	if err := c.do("listunspent", args, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// {
//     "txid": "ee99c2be6a111c5ee04e17b2b848b05e0e19cd291d8b4553b9bb0789720dbaff",
//     "vout": 0,
//     "address": "2MuQ4m4vyadPoTPxJVfZnASuutDQWTn2BVE",
//     "label": "",
//     "redeemScript": "0014fe67e5782a3e33f8da397b5e973d76098f35c9f7",
//     "scriptPubKey": "a914179c4acdb6bbb9ee3be5abc1281b1c35931c99c287",
//     "amount": 50.00000000,
//     "confirmations": 107,
//     "spendable": true,
//     "solvable": true,
//     "desc": "sh(wpkh([6e32d818/0'/0'/0']03bbe66ff0657083287222a71a3282b79da837494b49575553c6c997085d3d1203))#r6s8dqa7",
//     "safe": true
//   }

// listunspent
