package euler

import (
	"math/big"
	"testing"
)

var testFibonacciCases = []struct {
	Input    int
	Expected *big.Int
}{
	{0, big.NewInt(1)},
	{1, big.NewInt(1)},
	{2, big.NewInt(2)},
	{3, big.NewInt(3)},
	{4, big.NewInt(5)},
	{5, big.NewInt(8)},
	{6, big.NewInt(13)},
	{7, big.NewInt(21)},
	{8, big.NewInt(34)},
	{9, big.NewInt(55)},
	{10, big.NewInt(89)},
	{11, big.NewInt(144)},
}

func TestFibonacci(t *testing.T) {
	for _, tc := range testFibonacciCases {
		if actual := Fibonacci(tc.Input, nil); actual.Cmp(tc.Expected) != 0 {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}
