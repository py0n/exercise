package euler

import "testing"

var testGcmCases = []struct {
	InputX   int
	InputY   int
	Expected int
}{
	// 異常系: 0を含む
	{0, 2, 0},
	{3, 0, 0},
	// 異常系: 負數を含む
	{-2, 2, 0},
	{10, -5, 0},
	// 正常系
	{2, 3, 1},
	{12, 16, 4},
	{24, 36, 12},
}

func TestGcm(t *testing.T) {
	for _, tc := range testGcmCases {
		if actual := Gcm(tc.InputX, tc.InputY); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
