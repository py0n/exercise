package euler

import (
	"math/big"
	"testing"
)

var pe0025Cases = []struct {
	Input    int
	Expected int
}{
	{1000, 4782},
}

func TestPE0025(t *testing.T) {
	for _, tc := range pe0025Cases {
		if actual := PE0025(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

func Test_calculateDigitLength(t *testing.T) {
	cases := []struct {
		Input    *big.Int
		Expected int
	}{
		{big.NewInt(1), 1},
		{big.NewInt(10), 2},
		{big.NewInt(100), 3},
		{big.NewInt(1000), 4},
	}
	for _, tc := range cases {
		if actual := calculateDigitsLength(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}
