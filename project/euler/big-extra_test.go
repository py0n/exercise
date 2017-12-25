package euler

import (
	"math/big"
	"testing"
)

func TestBigNumberLength(t *testing.T) {
	cases := []struct {
		Input    *big.Int
		Expected int
	}{
		{big.NewInt(-1), 1},
		{big.NewInt(-10), 2},
		{big.NewInt(-100), 3},
		{big.NewInt(0), 1},
		{big.NewInt(1), 1},
		{big.NewInt(10), 2},
		{big.NewInt(100), 3},
		{big.NewInt(1000), 4},
	}
	for _, tc := range cases {
		if actual := BigNumberLength(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}
