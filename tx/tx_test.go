package tx

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTxCommon(t *testing.T) {
	tx := &TxBTC{}
	err := json.Unmarshal([]byte(`{
	"txid": "00112233445566778899aabbccddeeff",
	"hash": "66778899aabbccddeeff001122334455",
	"blockhash": "bbccddeeff00112233445566778899aa",
	"confirmations": 120,
	"blocktime": 1590572411,
	"locktime": 1590572411,
	"vin": [
		{
			"sequence": 4294967295,
			"coinbase": "00002f646372642f"
		},
		{
			"txid": "526b2de1ea84e31f6d7541b6aee3a90af6310c1ed0ada2eca485ac62b241550d",
			"vout": 3,
			"scriptSig": {
				"asm": "001413db5f88272ef4e6c05ecc7a4008e2a9a1f2574f",
				"hex": "16001413db5f88272ef4e6c05ecc7a4008e2a9a1f2574f"
			},
			"sequence": 4294967294
		}
	],
	"vout": [
		{
			"value": 48.9999668,
			"n": 0,
			"scriptPubKey": {
			  	"asm": "OP_HASH160 062ff2a72f2bb4f3dfed50ef5233bf3e68344f46 OP_EQUAL",
			  	"hex": "a914062ff2a72f2bb4f3dfed50ef5233bf3e68344f4687",
			  	"reqSigs": 1,
			  	"type": "scripthash",
			  	"addresses": [
					"2MsowUunWhc7UtwtgTPgSpgmhKyyRPCJ9Sz"
				]
			}
		},
		{
			"value": 1,
			"n": 1,
			"scriptPubKey": {
				"asm": "OP_HASH160 fb7d29ccd6171fb8f0d89098186ae45799e43cf9 OP_EQUAL",
				"hex": "a914fb7d29ccd6171fb8f0d89098186ae45799e43cf987",
				"reqSigs": 1,
				"type": "scripthash",
				"addresses": [
					"2NGAyXvw1zqqFeWL6MWkJoSkFXafJTq2BSZ"
			  	]
			}
		}
	]
}`), tx)
	outputs := tx.Outputs()
	require.NoError(t, err, "can't unmarshal")
	require.Len(t, outputs, 2)
	t.Logf("%#v\n", outputs[1])
}
