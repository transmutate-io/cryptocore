package types

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

type InvalidAmountError string

func (e InvalidAmountError) Error() string {
	return fmt.Sprintf("invalid decimal amount: %s\n", string(e))
}

type Amount string

func (a Amount) Clean() Amount {
	r := []rune(a)
	for {
		rsz := len(r)
		if rsz == 0 || r[0] != '0' {
			break
		}
		if rsz > 1 && r[1] == '.' {
			break
		}
		r = r[1:]
	}
	for {
		rsz := len(r)
		if rsz == 0 || r[rsz-1] != '0' {
			break
		}
		r = r[:rsz-1]
	}
	if rsz := len(r); rsz > 0 && r[rsz-1] == '.' {
		r = r[:rsz-1]
	}
	return Amount(r)
}

func NewAmountUInt64(a uint64, d int) Amount {
	dd := uint64(math.Pow10(d))
	aa, bb := a/dd, a%dd
	return Amount(fmt.Sprintf("%[1]d.%0.[3]*[2]d", aa, bb, d)).Clean()
}

func NewAmountBigInt(a *big.Int, d int) Amount {
	s := a.String()
	ssz := len(s)
	if ssz > d {
		intPart, decPart := s[:ssz-d], s[ssz-d:]
		return Amount(intPart + "." + decPart).Clean()
	}
	return Amount("0." + strings.Repeat("0", d-ssz) + s)
}

func (a Amount) String() string { return string(a.Clean()) }

func (a Amount) MarshalJSON() ([]byte, error) { return []byte(a), nil }

func (a *Amount) UnmarshalJSON(b []byte) error {
	v := Amount(b)
	if !v.Valid() {
		return InvalidAmountError(string(b))
	}
	*a = v
	return nil
}

func (a Amount) Decimals() int {
	p := strings.Split(string(a), ".")
	if len(p) < 2 {
		return 0
	}
	return len(p[1])
}

func (a Amount) WithDecimals(d int) Amount {
	p := strings.Split(string(a), ".")
	if len(p) == 1 {
		p = append(p, "")
	}
	if p1sz := len(p[1]); p1sz > d {
		p[1] = p[1][:d]
	} else {
		p[1] += strings.Repeat("0", d-p1sz)
	}
	return Amount(p[0] + "." + p[1])
}

func (a Amount) UInt64(d int) (uint64, error) {
	b, err := a.BigInt(d)
	if err != nil {
		return 0, err
	}
	if !b.IsUint64() {
		return 0, InvalidAmountError(a)
	}
	return b.Uint64(), nil
}

func (a Amount) BigInt(d int) (*big.Int, error) {
	parts := strings.Split(string(a.Clean()), ".")
	if len(parts) > 2 {
		return nil, InvalidAmountError(a)
	}
	if len(parts) == 1 {
		parts = append(parts, strings.Repeat("0", d))
	}
	dsz := len(parts[1])
	if dsz != d {
		parts[1] += strings.Repeat("0", d-dsz)
	}
	r, ok := big.NewInt(0).SetString(strings.Join(parts, ""), 10)
	if !ok {
		return nil, InvalidAmountError(a)
	}
	return r, nil
}

func (a Amount) Valid() bool {
	parts := strings.Split(string(a), ".")
	if len(parts) > 2 {
		return false
	}
	for _, i := range parts {
		if _, err := strconv.Atoi(i); err != nil {
			return false
		}
	}
	return true
}
