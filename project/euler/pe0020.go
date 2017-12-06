package euler

import (
	"math/big"
	"strconv"
)

// PE0020 階乗の各桁の和を計算
func PE0020(n int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		return 1
	}

	u := big.NewInt(int64(1))
	m := big.NewInt(int64(n))
	f := big.NewInt(int64(1))
	for m.Cmp(u) > 0 {
		f.Mul(f, m)
		m.Sub(m, u)
	}

	sum := 0
	s := f.String()
	for i := 0; i < len(s)-1; i++ {
		m, _ := strconv.Atoi(s[i : i+1])
		sum += m
	}
	return sum
}
