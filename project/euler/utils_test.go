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

func TestGcm(t *testing.T) {
	cases := []struct {
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
	for _, tc := range cases {
		if actual := Gcm(tc.InputX, tc.InputY); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func TestIsPalrindrome(t *testing.T) {
	cases := []struct {
		Input    string
		Expected bool
	}{
		// 異常系: 囘文で無い
		{"abcde", false},
		{"あいうえお", false},
		// 正常系: 囘文
		{"", true},
		{"abcba", true},
		{"あいういあ", true},
	}
	for _, tc := range cases {
		if actual := IsPalindrome(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func TestLcm(t *testing.T) {
	cases := []struct {
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
	for _, tc := range cases {
		if actual := Lcm(tc.InputX, tc.InputY); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func TestPow(t *testing.T) {
	cases := []struct {
		InputX   int
		InputY   int
		Expected int
	}{
		// 異常系: 指數が負數
		{2, -1, 0},
		// 正常系
		{0, 2, 0},
		{0, 0, 1},
		{2, 3, 8},
		{2, 0, 1},
	}
	for _, tc := range cases {
		if actual := Pow(tc.InputX, tc.InputY); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func TestPow10(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		// 異常系: 指數が負數
		{-1, 0},
		// 正常系
		{0, 1},
		{1, 10},
		{2, 100},
		{3, 1000},
	}
	for _, tc := range cases {
		if actual := Pow10(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func TestSumFactors(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		// 異常系: < 2
		{-1, 0},
		{0, 0},
		{1, 0},
		// 正常系
		{2, 3},
		{3, 4},
		{4, 7},
		{5, 6},
		{220, 504},
		{284, 504},
	}
	for _, tc := range cases {
		if actual := SumFactors(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func TestZellerWeekday(t *testing.T) {
	cases := []struct {
		InputY   int
		InputM   int
		InputD   int
		Expected DayOfWeek
	}{
		// 異常系: グレゴリオ歴以前
		{1582, 10, 14, None},
		{1580, 1, 1, None},
		// 正常系
		{1582, 10, 15, Friday},
		{1904, 1, 4, Monday},
		{1904, 4, 5, Tuesday},
	}
	for _, tc := range cases {
		if actual := ZellerWeekday(tc.InputY, tc.InputM, tc.InputD); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Test_multiplyDigits(t *testing.T) {
	cases := []struct {
		Input0   int
		Expected int
	}{
		{1234, 24},
		{9876, 3024},
	}
	for _, tc := range cases {
		if actual := multiplyDigits(tc.Input0); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
