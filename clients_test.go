package btccore

import (
	"testing"

	"github.com/stretchr/testify/require"
	"transmutate.io/pkg/btccore/types"
)

func TestClient(t *testing.T) {
	const amt = "1"
	clients := []struct {
		name    string
		nBlocks int
		amount  types.Amount
		cl      Client
	}{
		{
			"BTC",
			101,
			types.Amount(amt),
			NewClientBTC("bitcoin-core-testnet.docker:4444", "admin", "pass", false),
		},
		{
			"LTC",
			101,
			types.Amount(amt),
			NewClientLTC("litecoin-testnet.docker:4444", "admin", "pass", false),
		},
		{
			"DOGE",
			101,
			types.Amount(amt),
			NewClientDOGE("dogecoin-testnet.docker:4444", "admin", "pass", false),
		},
	}
	for _, i := range clients {
		t.Run(i.name, func(t *testing.T) {
			// generate a new address
			addr, err := i.cl.NewAddress()
			require.NoError(t, err, "can't generate address")
			// dump the private key
			_, err = i.cl.DumpPrivateKey(addr)
			require.NoError(t, err, "can't dump private key")
			// generate blocks
			_, err = i.cl.GenerateToAddress(i.nBlocks, addr)
			require.NoError(t, err, "can't generate blocks")
			// generate a new address
			addr2, err := i.cl.NewAddress()
			require.NoError(t, err, "can't generate address")
			// send to new address
			txID, err := i.cl.SendToAddress(addr2, i.amount)
			require.NoError(t, err, "can't send to address")
			// generate blocks
			_, err = i.cl.GenerateToAddress(i.nBlocks, addr)
			require.NoError(t, err, "can't generate blocks")
			bal, err := i.cl.Balance(0)
			require.NoError(t, err, "can't get balance")
			_ = bal
			bc, err := i.cl.BlockCount()
			require.NoError(t, err, "can't get block count")
			bh, err := i.cl.BlockHash(bc)
			require.NoError(t, err, "can't get block hash")
			_, err = i.cl.RawBlock(bh)
			require.NoError(t, err, "can't get raw block")
			_, err = i.cl.Block(bh)
			require.NoError(t, err, "can't get block")
			_, err = i.cl.Transaction(txID)
			require.NoError(t, err, "can't find transaction")
		})
	}
}
