package btccore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func newClient() *Client {
	return &Client{
		Address:  "127.0.0.1:3333",
		Username: "admin",
		Password: "pass",
	}
}

func TestDo(t *testing.T) {
	cl := newClient()
	err := cl.do("non-existing-method", nil, nil)
	require.Error(t, err, "expecting an error")
	var bh string
	err = cl.do("getblockhash", []interface{}{1}, &bh)
	require.NoError(t, err, "not expecting an error")
}

func TestGetBlock(t *testing.T) {
	cl := newClient()
	hash, err := cl.GetBlockHash(100)
	require.NoError(t, err, "not expecting an error")
	_, err = cl.GetBlock(hash, false)
	require.NoError(t, err, "not expecting an error")
}
