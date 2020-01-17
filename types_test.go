package btccore

import (
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMarshalUnmarshalHexBytes(t *testing.T) {
	hb := HexBytes{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	b, err := json.Marshal(hb)
	require.NoError(t, err, "can't marshal")
	newHb := make(HexBytes, 0, 16)
	err = json.Unmarshal(b, &newHb)
	require.NoError(t, err, "can't unmarshal")
	require.Equal(t, hb, newHb, "bytes mismatch")
}

func TestMarshalUnmarshalUnixTime(t *testing.T) {
	ut := NewUnixTime(time.Now())
	b, err := json.Marshal(ut)
	require.NoError(t, err, "can't marshal")
	var newUt UnixTime
	err = json.Unmarshal(b, &newUt)
	require.NoError(t, err, "can't unmarshal")
	require.Equal(t, ut, newUt, "time mismatch")
}

func TestMarshalUnmarshalBlock(t *testing.T) {
	bt := BlockTransactions{
		HexBytes{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		HexBytes{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0},
		HexBytes{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0, 1},
		HexBytes{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0, 1, 2},
		HexBytes{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0, 1, 2, 3},
		HexBytes{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0, 1, 2, 3, 4},
		HexBytes{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0, 1, 2, 3, 4, 5},
		HexBytes{7, 8, 9, 10, 11, 12, 13, 14, 15, 0, 1, 2, 3, 4, 5, 6},
	}
	b, err := json.Marshal(bt)
	require.NoError(t, err, "can't marshal tx hashes")
	newBt := BlockTransactions{}
	err = json.Unmarshal(b, &newBt)
	require.NoError(t, err, "can't unmarshal tx hashes")
	require.Equal(t, bt, newBt)
	bt = BlockTransactions{&RawTransactionInfo{
		Hex:  HexBytes{},
		Hash: HexBytes{},
		ID:   HexBytes{},
	}}
	b, err = json.Marshal(bt)
	require.NoError(t, err, "can't marshal raw tx")
	newBt = BlockTransactions{}
	err = json.Unmarshal(b, &newBt)
	require.NoError(t, err, "can't unmarshal raw tx")
	require.Len(t, newBt, 1, "expecting 1 empty raw tx")
	require.Equal(t, bt[0], newBt[0])
}

func TestMarshalUnmarshalAmount(t *testing.T) {
	a := (*Amount)(big.NewInt(100001000))
	b, err := json.Marshal(a)
	require.NoError(t, err, "can't marshal amount")
	newA := &Amount{}
	err = json.Unmarshal(b, newA)
	require.NoError(t, err, "can't unmarshal amount")
	require.Zero(t, (*big.Int)(a).Cmp((*big.Int)(newA)), "values mismatch")
}
