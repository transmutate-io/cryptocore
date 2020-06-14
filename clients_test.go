package cryptocore

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"transmutate.io/pkg/cryptocore/types"
)

const amt = "1"

var clients = []struct {
	name    string
	nBlocks int
	amount  types.Amount
	cl      Client
}{
	{
		"BTC",
		101,
		types.Amount(amt),
		NewClientBTC("bitcoin-core-regtest.docker:4444", "admin", "pass", nil),
	},
	{
		"LTC",
		101,
		types.Amount(amt),
		NewClientLTC("litecoin-regtest.docker:4444", "admin", "pass", nil),
	},
	{
		"DOGE",
		101,
		types.Amount(amt),
		NewClientDOGE("dogecoin-regtest.docker:4444", "admin", "pass", nil),
	},
	{
		"BCH",
		101,
		types.Amount(amt),
		NewClientBCH("bitcoin-cash-regtest.docker:4444", "admin", "pass", nil),
	},
	{
		"DCR",
		101,
		types.Amount(amt),
		NewClientDCR("decred-wallet-simnet.docker:4444", "admin", "pass", &TLSConfig{SkipVerify: true}),
	},
}

func generateBlocks(cl Client, nBlocks int, addr string) ([]types.Bytes, error) {
	if cl.CanGenerateBlocksToAddress() {
		return cl.GenerateBlocksToAddress(nBlocks, addr)
	} else {
		return cl.GenerateBlocks(nBlocks)
	}
}

func TestClient(t *testing.T) {
	for _, i := range clients {
		t.Run(i.name, func(t *testing.T) {
			// generate a new address
			addr, err := i.cl.NewAddress()
			require.NoError(t, err, "can't generate address")
			// generate blocks
			_, err = generateBlocks(i.cl, 101, addr)
			require.NoError(t, err, "can't generate blocks")
			// get block count
			bc, err := i.cl.BlockCount()
			require.NoError(t, err, "can't get block count")
			// get block hash
			bh, err := i.cl.BlockHash(bc)
			require.NoError(t, err, "can't get block hash")
			// get raw block
			_, err = i.cl.RawBlock(bh)
			require.NoError(t, err, "can't get raw block")
			// get last block
			blk, err := i.cl.Block(bh)
			require.NoError(t, err, "can't get block")
			// get transaction
			txs := blk.Transactions()
			lastTx, err := i.cl.Transaction(txs[0])
			require.NoError(t, err, "can't find transaction")
			// iterate all blocks
			blkIter := NewBlockIterator(i.cl, bc-5)
			for {
				_, err = blkIter.Next()
				if err == ErrNoBlock {
					break
				}
				require.NoError(t, err, "can't iterate blocks")
			}
			// iterate transactions
			txIter := NewTransactionIterator(i.cl, bc-5)
			for {
				tx, err := txIter.Next()
				require.NoError(t, err, "can't iterate transactions")
				if bytes.Equal(tx.ID(), lastTx.ID()) {
					break
				}
			}
		})
	}
}

func TestBalance(t *testing.T) {
	for _, i := range clients {
		t.Run(i.name, func(t *testing.T) {
			bal, err := i.cl.Balance(0)
			require.NoError(t, err, "can't get balance")
			t.Logf("::: %#v\n", bal)
		})
	}
}
