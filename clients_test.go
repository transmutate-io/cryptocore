package cryptocore

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/transmutate-io/cryptocore/block"
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
		nBlocks: 1,
		amount:  types.Amount(amt),
		cl:      cl,
	})
	if cl, err = NewClientLTC("litecoin-localnet.docker:4444", "admin", "pass", nil); err != nil {
		panic(err)
	}
	clients = append(clients, clientConfig{
		name:    "LTC",
		nBlocks: 1,
		amount:  types.Amount(amt),
		cl:      cl,
	})
	if cl, err = NewClientDOGE("dogecoin-localnet.docker:4444", "admin", "pass", nil); err != nil {
		panic(err)
	}
	clients = append(clients, clientConfig{
		name:    "DOGE",
		nBlocks: 1,
		amount:  types.Amount(amt),
		cl:      cl,
	})
	if cl, err = NewClientBCH("bitcoin-cash-localnet.docker:4444", "admin", "pass", nil); err != nil {
		panic(err)
	}
	clients = append(clients, clientConfig{
		name:    "BCH",
		nBlocks: 1,
		amount:  types.Amount(amt),
		cl:      cl,
	})
	if cl, err = NewClientDCR("decred-wallet-localnet.docker:4444", "admin", "pass", &TLSConfig{SkipVerify: true}); err != nil {
		panic(err)
	}
	clients = append(clients, clientConfig{
		name:    "DCR",
		nBlocks: 1,
		amount:  types.Amount(amt),
		cl:      cl,
	})
	if cl, err = NewClientETH("ethereum-localnet.docker:4444", "", "", nil); err != nil {
		panic(err)
	}
	clients = append(clients, clientConfig{
		name:    "ETH",
		nBlocks: 1,
		amount:  types.Amount(amt),
		cl:      cl,
	})
}

var errNotABlockGenerator = errors.New("not a block generator")

func generateBlocks(cl Client, nBlocks int, addr string) ([]types.Bytes, error) {
	if bg, ok := cl.(TargetedBlockGenerator); ok {
		return bg.GenerateToAddress(nBlocks, addr)
	} else if bg, ok := cl.(BlockGenerator); ok {
		return bg.Generate(nBlocks)
	}
	return nil, errNotABlockGenerator
}

func TestClient(t *testing.T) {
	for _, i := range clients {
		t.Run(i.name, func(t *testing.T) {
			// generate a new address
			var (
				addr string
				err  error
			)
			if addrGen, ok := i.cl.(AddressGenerator); ok {
				addr, err = addrGen.NewAddress()
				require.NoError(t, err, "can't generate address")
			}
			// generate blocks
			_, err = generateBlocks(i.cl, i.nBlocks, addr)
			if err != errNotABlockGenerator {
				require.NoError(t, err, "can't generate blocks")
			}
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
			if blkTx, ok := blk.(block.TransactionsLister); ok {
				txs := blkTx.TransactionsHashes()
				_, err = i.cl.Transaction(txs[0])
				require.NoError(t, err, "can't find transaction")
			}
		})
	}
}

func TestBalance(t *testing.T) {
	for _, i := range clients {
		t.Run(i.name, func(t *testing.T) {
			if balancer, ok := i.cl.(Balancer); ok {
				bal, err := balancer.Balance(0)
				require.NoError(t, err, "can't get balance")
				t.Logf("::: %#v\n", bal)
			}
		})
	}
}
