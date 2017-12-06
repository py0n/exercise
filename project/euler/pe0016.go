package euler

import (
	"math/big"
	"strconv"
)

/*
	https://projecteuler.net/problem=16

	2^1000 の各桁の和

	2^1000 は大体300桁(log10 2 = 0.30....)
*/
func PE0016(n int) int {
	// 2^n
	m := big.NewInt(2)
	for i := 0; i < n-1; i++ {
		m.Mul(m, big.NewInt(2))
	}

	s := m.String()

	sum := 0
	for i := 0; i < len(s); i++ {
		j, _ := strconv.Atoi(s[i : i+1])
		sum += j
	}
	return sum
}
