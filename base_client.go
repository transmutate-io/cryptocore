package cryptocore

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"transmutate.io/pkg/cryptocore/types"
)

type baseClient struct {
	Address   string
	Username  string
	Password  string
	TLSConfig *TLSConfig
}

func newBaseClient(addr, user, pass string, tlsConf *TLSConfig) *baseClient {
	return &baseClient{
		Address:   addr,
		Username:  user,
		Password:  pass,
		TLSConfig: tlsConf,
	}
}

func (c *baseClient) do(method string, params interface{}, r interface{}) error {
	b := bytes.NewBuffer(make([]byte, 0, 1024))
	err := json.NewEncoder(b).Encode(&rpcRequest{
		JsonRPC: "1.0",
		ID:      "go-cryptocore",
		Method:  method,
		Params:  params,
	})
	if err != nil {
		return err
	}
	cl := &http.Client{}
	callURL := append(make([]string, 0, 16), "http")
	if c.TLSConfig != nil {
		callURL = append(callURL, "s")
		tlsConf := &tls.Config{InsecureSkipVerify: c.TLSConfig.SkipVerify}
		if c.TLSConfig.CA != "" {
			b, err := ioutil.ReadFile(c.TLSConfig.CA)
			if err != nil {
				return err
			}
			cert, err := x509.ParseCertificate(b)
			if err != nil {
				return err
			}
			rootCAs := x509.NewCertPool()
			rootCAs.AddCert(cert)
			tlsConf.RootCAs = rootCAs
		}
		if c.TLSConfig.ClientCertificate != "" && c.TLSConfig.ClientKey != "" {
			cert, err := tls.LoadX509KeyPair(c.TLSConfig.ClientCertificate, c.TLSConfig.ClientKey)
			if err != nil {
				return err
			}
			tlsConf.Certificates = append(tlsConf.Certificates, cert)
		}
		cl.Transport = &http.Transport{TLSClientConfig: tlsConf}
	}
	callURL = append(callURL, "://")
	if c.Username != "" {
		callURL = append(callURL, c.Username, ":", c.Password)
	}
	callURL = append(callURL, "@", c.Address, "/")
	resp, err := cl.Post(strings.Join(callURL, ""), "application/json", b)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
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

// func (c *baseClient) DumpPrivateKey(addr string) (string, error) {
// 	return c.doString("dumpprivkey", args(addr))
// }

func (c *baseClient) CanGenerateBlocksToAddress() bool { return true }

func (c *baseClient) CanGenerateBlocks() bool { return true }

func (c *baseClient) GenerateBlocks(nBlocks int) ([]types.Bytes, error) {
	return c.doSliceBytes("generate", args(nBlocks))
}

func (c *baseClient) GenerateBlocksToAddress(nBlocks int, addr string) ([]types.Bytes, error) {
	return c.doSliceBytes("generatetoaddress", args(nBlocks, addr))
}

func (c *baseClient) SendToAddress(addr string, value types.Amount) (types.Bytes, error) {
	return c.doBytes("sendtoaddress", args(addr, value))
}

func (c *baseClient) BlockCount() (uint64, error) { return c.doUint64("getblockcount", nil) }

func (c *baseClient) BlockHash(height uint64) (types.Bytes, error) {
	r, err := c.doBytes("getblockhash", args(height))
	if err != nil {
		if e, ok := err.(*ClientError); ok && e.Code == -8 {
			return nil, ErrNoBlock
		}
		return nil, err
	}
	return r, nil
}

func (c *baseClient) SendRawTransaction(tx types.Bytes) (types.Bytes, error) {
	return c.doBytes("sendrawtransaction", args(tx.Hex()))
}

func (c *baseClient) RawTransaction(hash types.Bytes) (types.Bytes, error) {
	return c.doBytes("getrawtransaction", args(hash.Hex(), false))
}

func (c *baseClient) transaction(t interface{}, args interface{}) error {
	return c.do("getrawtransaction", args, t)
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

func (c *baseClient) block(b interface{}, args interface{}) error {
	return c.do("getblock", args, b)
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
