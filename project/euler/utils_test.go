package euler

import (
	"testing"
)

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
