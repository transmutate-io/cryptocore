package cryptocore

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/transmutate-io/cryptocore/types"
)

const amt = "1"

type clientConfig struct {
	name    string
	nBlocks int
	amount  types.Amount
	cl      Client
}

var clients = make([]clientConfig, 0, 8)

func init() {
	var (
		cl  Client
		err error
	)
	if cl, err = NewClientBTC("bitcoin-core-localnet.docker:4444", "admin", "pass", nil); err != nil {
		panic(err)
	}
	clients = append(clients, clientConfig{
		name:    "BTC",
		nBlocks: 101,
		amount:  types.Amount(amt),
		cl:      cl,
	})
	if cl, err = NewClientLTC("litecoin-localnet.docker:4444", "admin", "pass", nil); err != nil {
		panic(err)
	}
	clients = append(clients, clientConfig{
		name:    "LTC",
		nBlocks: 101,
		amount:  types.Amount(amt),
		cl:      cl,
	})
	if cl, err = NewClientDOGE("dogecoin-localnet.docker:4444", "admin", "pass", nil); err != nil {
		panic(err)
	}
	clients = append(clients, clientConfig{
		name:    "DOGE",
		nBlocks: 101,
		amount:  types.Amount(amt),
		cl:      cl,
	})
	if cl, err = NewClientBCH("bitcoin-cash-localnet.docker:4444", "admin", "pass", nil); err != nil {
		panic(err)
	}
	clients = append(clients, clientConfig{
		name:    "BCH",
		nBlocks: 101,
		amount:  types.Amount(amt),
		cl:      cl,
	})
	if cl, err = NewClientDCR("decred-wallet-localnet.docker:4444", "admin", "pass", &TLSConfig{SkipVerify: true}); err != nil {
		panic(err)
	}
	clients = append(clients, clientConfig{
		name:    "DCR",
		nBlocks: 101,
		amount:  types.Amount(amt),
		cl:      cl,
	})
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
			// lastTx, err := i.cl.Transaction(txs[0])
			_, err = i.cl.Transaction(txs[0])
			require.NoError(t, err, "can't find transaction")
			// // iterate all blocks
			// blkIter := NewBlockIterator(i.cl, bc-5)
			// for {
			// 	_, err = blkIter.Next()
			// 	if err == ErrNoBlock {
			// 		break
			// 	}
			// 	require.NoError(t, err, "can't iterate blocks")
			// }
			// // iterate transactions
			// txIter := NewTransactionIterator(i.cl, bc-5)
			// for {
			// 	tx, err := txIter.Next()
			// 	require.NoError(t, err, "can't iterate transactions")
			// 	if bytes.Equal(tx.ID(), lastTx.ID()) {
			// 		break
			// 	}
			// }
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
