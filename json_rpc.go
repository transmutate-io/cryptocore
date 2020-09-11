package cryptocore

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/transmutate-io/cryptocore/types"
)

type jsonRPCClient struct {
	Address         string
	Username        string
	Password        string
	tlsConfig       *TLSConfig
	cachedTLSConfig *tls.Config
}

func newJsonRpcClient(addr, user, pass string, tlsConf *TLSConfig) (*jsonRPCClient, error) {
	r := &jsonRPCClient{
		Address:  addr,
		Username: user,
		Password: pass,
	}
	if err := r.SetTLSConfig(tlsConf); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *jsonRPCClient) SetTLSConfig(cfg *TLSConfig) error {
	c.tlsConfig = cfg
	nc, err := c.newTLSConfig()
	if err != nil {
		return err
	}
	c.cachedTLSConfig = nc
	return nil
}

func (c *jsonRPCClient) newTLSConfig() (*tls.Config, error) {
	if c.tlsConfig == nil {
		return nil, nil
	}
	tlsConf := &tls.Config{InsecureSkipVerify: c.tlsConfig.SkipVerify}
	if c.tlsConfig.CA != "" {
		b, err := ioutil.ReadFile(c.tlsConfig.CA)
		if err != nil {
			return nil, err
		}
		cert, err := x509.ParseCertificate(b)
		if err != nil {
			return nil, err
		}
		rootCAs := x509.NewCertPool()
		rootCAs.AddCert(cert)
		tlsConf.RootCAs = rootCAs
	}
	if c.tlsConfig.ClientCertificate != "" && c.tlsConfig.ClientKey != "" {
		cert, err := tls.LoadX509KeyPair(c.tlsConfig.ClientCertificate, c.tlsConfig.ClientKey)
		if err != nil {
			return nil, err
		}
		tlsConf.Certificates = append(tlsConf.Certificates, cert)
	}
	return tlsConf, nil
}

func clientURL(addr string, username string, password string, tlsConf *TLSConfig) string {
	callURL := append(make([]string, 0, 16), "http")
	if tlsConf != nil {
		callURL = append(callURL, "s")
	}
	// use cached config
	callURL = append(callURL, "://")
	if username != "" {
		callURL = append(callURL, username, ":", password, "@")
	}
	return strings.Join(append(callURL, addr, "/"), "")
}

func (c *jsonRPCClient) clientURL() string {
	return clientURL(c.Address, c.Username, c.Password, c.tlsConfig)
}

func (c *jsonRPCClient) do(method string, params interface{}, r interface{}) error {
	b := make([]byte, 0, 1024)
	bb := bytes.NewBuffer(b)
	err := json.NewEncoder(bb).Encode(&rpcRequest{
		JsonRPC: "1.0",
		ID:      "go-cryptocore",
		Method:  method,
		Params:  params,
	})
	if err != nil {
		return err
	}
	cl := &http.Client{Transport: &http.Transport{TLSClientConfig: c.cachedTLSConfig}}
	resp, err := cl.Post(c.clientURL(), "application/json", bb)
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

func (c *jsonRPCClient) doString(method string, args interface{}) (string, error) {
	var r string
	if err := c.do(method, args, &r); err != nil {
		return "", err
	}
	return r, nil
}

func (c *jsonRPCClient) doUint64(method string, args interface{}) (uint64, error) {
	var r uint64
	if err := c.do(method, args, &r); err != nil {
		return 0, err
	}
	return r, nil
}

func (c *jsonRPCClient) doBytes(method string, args interface{}) (types.Bytes, error) {
	var r types.Bytes
	if err := c.do(method, args, &r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *jsonRPCClient) doSliceBytes(method string, args interface{}) ([]types.Bytes, error) {
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

func (ce *ClientError) Error() string { return ce.Message }
