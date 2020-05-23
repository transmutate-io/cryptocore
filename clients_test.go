package cryptocore

import (
	"bytes"
	"encoding/hex"
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
}

func TestClient(t *testing.T) {
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
			// get block count
			bc, err := i.cl.BlockCount()
			require.NoError(t, err, "can't get block count")
			// get block count
			bh, err := i.cl.BlockHash(bc)
			require.NoError(t, err, "can't get block hash")
			// get raw block
			_, err = i.cl.RawBlock(bh)
			require.NoError(t, err, "can't get raw block")
			// get last block
			_, err = i.cl.Block(bh)
			require.NoError(t, err, "can't get block")
			// get transaction
			_, err = i.cl.Transaction(txID)
			require.NoError(t, err, "can't find transaction")
			// iterate all blocks
			nextblock, closeIter := NewBlockIterator(i.cl, 1)
			for i := uint64(1); i < bc; i++ {
				_, err = nextblock()
				require.NoError(t, err, "can't iterate blocks")
			}
			closeIter()
			// iterate transactions
			nextTx, closeIter := NewTransactionIterator(i.cl, 1)
			for {
				tx, err := nextTx()
				require.NoError(t, err, "can't iterate transactions")
				if bytes.Equal(tx.ID, txID) {
					break
				}
			}
			closeIter()
		})
	}
}

func TestSendRawTransaction(t *testing.T) {
	for _, i := range clients {
		t.Run(i.name, func(t *testing.T) {
			b, _ := hex.DecodeString("0100000002a61146522a305aac6e2ef94ecaf5f4a5798bb01ca43c6a018f9db291ab221ed600000000e147304402205984f9d7dab5924c7d20362aa63064117c52cc3c97c6ff3fcb66ee2db25642ac022036046cc0315b17a3526652b83f087fe33380e52bcfa78f4a1b58868c299362de01210295b87234904bc7e5e55570e3a14eccb0528d56525884d00ae6045dc28c3d375920a0db52560bae8d8159a4aed02d1c863439bd83ea9eb77099e80c07d464cb56e7004c53630480ee8f5eb17576a9142991e7d62bcdc636e3855b8583d127f8a7e35dcc88ac67a9142a28f32a4435e3f673f6279f1d8c029d4f1c234a8876a914e8132a43fa082a44474848b3be59fdb1910ef5dd88ac68ffffffff050cf032abcb9448d81abf96b4bdb166c878e31ae1808098ea77ec028196ae5d01000000e2483045022100d630403335cb70d15609369d98dd60d35c2223c72705f290e916940340eb42e2022007af80ffac8ab16324ab4ea66faa2d8739fa070b1e1cd570b4f7ffe9e2e2f6b901210295b87234904bc7e5e55570e3a14eccb0528d56525884d00ae6045dc28c3d375920a0db52560bae8d8159a4aed02d1c863439bd83ea9eb77099e80c07d464cb56e7004c53630480ee8f5eb17576a9142991e7d62bcdc636e3855b8583d127f8a7e35dcc88ac67a9142a28f32a4435e3f673f6279f1d8c029d4f1c234a8876a914e8132a43fa082a44474848b3be59fdb1910ef5dd88ac68ffffffff01b6b6eb0b000000001976a91495fce05fe7c10e1f0fd9374fba32b85f94e8677788ac00000000")
			txid, err := i.cl.SendRawTransaction(b)
			t.Logf("txid: %v\n", txid)
			t.Logf("error: %v\n", err)
		})
	}
}
