package types

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type InvalidAmountError string

func (e InvalidAmountError) Error() string {
	return fmt.Sprintf("invalid decimal amount: %s\n", string(e))
}

type Amount string

func splitAmount(amt string) (uint64, uint64, error) {
	p := strings.Split(amt, ".")
	if len(p) > 2 {
		return 0, 0, InvalidAmountError(amt)
	}
	a, err := strconv.Atoi(p[0])
	if err != nil {
		return 0, 0, InvalidAmountError(amt)
	}
	var b int
	if len(p) > 1 {
		if b, err = strconv.Atoi(p[1]); err != nil {
			return 0, 0, InvalidAmountError(amt)
		}
	}
	return uint64(a), uint64(b), nil
}

func NewAmount(a uint64, p uint64) Amount {
	d := uint64(math.Pow10(int(p)))
	aa, bb := a/d, a%d
	return Amount(fmt.Sprintf("%[1]d.%0.[3]*[2]d", aa, bb, p))
}

func ParseAmount(amt string) (Amount, error) {
	if _, _, err := splitAmount(amt); err != nil {
		return "", err
	}
	return Amount(amt), nil
}

func (a Amount) String() string               { return string(a) }
func (a Amount) MarshalJSON() ([]byte, error) { return []byte(a), nil }

func (a *Amount) UnmarshalJSON(b []byte) error {
	if _, _, err := splitAmount(string(b)); err != nil {
		return err
	}
	*a = Amount(b)
	return nil
}

func (a Amount) Prec() uint64 {
	p := strings.Split(string(a), ".")
	if len(p) < 2 {
		return 0
	}
	return uint64(len(p[1]))
}

func (a Amount) WithPrec(n int) Amount {
	p := strings.Split(string(a), ".")
	if len(p) == 1 {
		p = append(p, "")
	}
	if p1sz := len(p[1]); p1sz > n {
		p[1] = p[1][:n]
	} else {
		p[1] += strings.Repeat("0", n-p1sz)
	}
	return Amount(p[0] + "." + p[1])
}

func (a Amount) UInt64(p int) uint64 {
	p1, p2, err := splitAmount(string(a.WithPrec(p)))
	if err != nil {
		panic(InvalidAmountError(a))
	}
	return p1*uint64(math.Pow10(p)) + p2
}

func (a Amount) Valid() bool {
	if _, _, err := splitAmount(string(a)); err != nil {
		return false
	}
	return true
}
