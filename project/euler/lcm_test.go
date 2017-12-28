package euler

import "testing"

var testLcmCases = []struct {
	InputX   int
	InputY   int
	Expected int
}{
	// 異常系:0を含む
	{0, 3, 0},
	{2, 0, 0},
	// 異常系:負數を含む
	{-2, 3, 0},
	{12, -16, 0},
	// 正常系
	{2, 3, 6},
	{12, 16, 48},
	{24, 36, 72},
}

func TestLcm(t *testing.T) {
	for _, tc := range testLcmCases {
		if actual := Lcm(tc.InputX, tc.InputY); actual != tc.Expected {
			t.Errorf(
				"InputX:%v\nInputY:%v\nExpected:%v\nActual:%v",
				tc.InputX,
				tc.InputY,
				tc.Expected,
				actual,
			)
		}
	}
}
