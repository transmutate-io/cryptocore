package block

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlockBTC(t *testing.T) {
	b := &BlockBTC{}
	err := json.Unmarshal([]byte(`
{
	"hash": "aabbccddeeff00112233445566778899",
	"confirmations": 42,
	"height": 150,
	"tx": [ "00112233445566778899aabbccddeeff" ],
	"time": 1590572411,
	"previousblockhash": "ccddeeff00112233445566778899aabb",
	"nextblockhash": "5566778899aabbccddeeff0011223344"
}`), b)
	require.NoError(t, err, "can't unmarshal block")
	t.Logf("%#v\n", b)
}
