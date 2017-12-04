package project_euler

import (
	"testing"
)

const N = 1000000

func Benchmark_PrimeGenerator0(b *testing.B) {
	PrimeGenerator := NewPrimeGenerator0()
	for i := 0; i < N; i++ {
		PrimeGenerator.Next()
	}
}

func Benchmark_PrimeGenerator1(b *testing.B) {
	PrimeGenerator := NewPrimeGenerator1()
	for i := 0; i < N; i++ {
		PrimeGenerator.Next()
	}
}

func Benchmark_PrimeGenerator(b *testing.B) {
	PrimeGenerator := NewPrimeGenerator()
	for i := 0; i < N; i++ {
		PrimeGenerator.Next()
	}
}

func Test_PrimeGenerator(t *testing.T) {
	PrimeGenerator := NewPrimeGenerator()

	for _, expected := range []int{2, 3, 5, 7, 11, 13, 17, 19, 23} {
		if actual := PrimeGenerator.Next(); actual != expected {
			t.Errorf("expected=%v, actual=%v", expected, actual)
		}
	}
}

func Test_Gcm(t *testing.T) {
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

func Test_isPalrindrome(t *testing.T) {
	cases := []struct {
		Input    string
		Expected bool
	}{
		{"abcde", false},
		{"abcba", true},
		{"abccba", true},
	}
	for _, tc := range cases {
		if actual := isPalindrome(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Test_Lcm(t *testing.T) {
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

func Test_pow10(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{0, 1},
		{1, 10},
		{2, 100},
		{3, 1000},
	}
	for _, tc := range cases {
		if actual := pow10(tc.Input); actual != tc.Expected {
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

func Test_zellerWeekday(t *testing.T) {
	cases := []struct {
		InputY   int
		InputM   int
		InputD   int
		Expected int
	}{
		{1904, 1, 4, 1},
		{1904, 4, 5, 2},
	}
	for _, tc := range cases {
		if actual := zellerWeekday(tc.InputY, tc.InputM, tc.InputD); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
