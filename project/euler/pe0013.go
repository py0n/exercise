package euler

import (
	"math/big"
	"strconv"
)

// PE0013 与へられたn個の整數の和の最初の10桁を計算する。
func PE0013(digits []string) int {
	sum := big.NewInt(0)
	for _, s := range digits {
		n := new(big.Int)
		n.SetString(s, 10)
		sum.Add(sum, n)
	}
	s := sum.String()
	m, _ := strconv.Atoi(s[0:10])
	return m
}
