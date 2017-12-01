package project_euler

import (
	"math/big"
)

// https://projecteuler.net/problem=15

/*
	+-+-+-+
	| | | |
	+-+-+-+
	| | | |
	+-+-+-+
	| | | |
	+-+-+-+

	|  | 1| 2| 3| 4| 5|
	| 1| 2| 3| 4| 5| 6|
	| 2| 3| 6|10|15|
	| 3| 4|10|20|  |
	| 4| 5|15|  |  |
	| 5| 6|

	n * m -> (n+m)Cn
*/
func PE0015(n int64) int64 {
	numerator := big.NewInt(1) // 分子
	for i := int64(0); i < n; i++ {
		numerator.Mul(numerator, big.NewInt(int64(2*n-i)))
	}
	denominator := big.NewInt(1) //分母
	for i := int64(0); i < n; i++ {
		denominator.Mul(denominator, big.NewInt(int64(n-i)))
	}
	r := big.NewInt(0)
	r.Quo(numerator, denominator)
	return r.Int64()
}
