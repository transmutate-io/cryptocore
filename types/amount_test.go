package types

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAmount(t *testing.T) {
	stringTests := []struct {
		a   uint64
		d   int
		e   string
		ed  int
		eds Amount
	}{
		{100001000, 8, "1.00001", 5, "1.00001000"},
		{100000001, 8, "1.00000001", 8, "1.00000001"},
		{100000000, 8, "1", 0, "1.00000000"},
	}
	for _, i := range stringTests {
		amt := NewAmountUInt64(i.a, i.d)
		require.Equal(t, i.e, amt.String())
		require.Equal(t, i.ed, amt.Decimals())
		require.Equal(t, i.eds, amt.WithDecimals(i.d))
		ui, err := amt.UInt64(i.d)
		require.NoError(t, err)
		require.Equal(t, i.a, ui)
		bi, err := amt.BigInt(i.d)
		require.NoError(t, err)
		require.Zero(t, bi.Cmp(big.NewInt(int64(i.a))))
		require.True(t, amt.Valid())
		b, err := json.Marshal(amt)
		require.NoError(t, err)
		amt2 := NewAmountBig(big.NewInt(int64(i.a)), i.d)
		bi, err = amt.BigInt(i.d)
		require.NoError(t, err)
		bi2, err := amt2.BigInt(i.d)
		require.NoError(t, err)
		require.Zero(t, bi.Cmp(bi2))
		err = json.Unmarshal(b, &amt2)
		require.NoError(t, err)
		require.Equal(t, amt.Clean(), amt2.Clean())
	}
}
