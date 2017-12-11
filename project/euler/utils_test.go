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

func TestCalculateRecurringCycle(t *testing.T) {
	cases := []struct {
		Input    int
		Expected string
	}{
		{1, ""},
		{2, ""},
		{3, "3"},
		{4, ""},
		{5, ""},
		{6, "6"},
		{7, "142857"},
		{8, ""},
		{9, "1"},
		{10, ""},
		{11, "09"},
		{13, "076923"},
		{61, "016393442622950819672131147540983606557377049180327868852459"},
		{97, "010309278350515463917525773195876288659793814432989690721649484536082474226804123711340206185567"},
	}
	for _, tc := range cases {
		if actual := CalculateRecurringCycle(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

func TestCalculateRecurringCycleLength(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{1, 0},
		{2, 0},
		{3, 1},
		{4, 0},
		{5, 0},
		{6, 1},
		{7, 6},
		{8, 0},
		{9, 1},
		{10, 0},
		{11, 2},
		{13, 6},
		{61, 60},
		{97, 96},
	}
	for _, tc := range cases {
		if actual := CalculateRecurringCycleLength(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

func TestCollatzLength(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		// 異常系: < 2
		{-1, 0},
		{0, 0},
		{1, 0},
		// 正常系
		{4, 3},
		{13, 10},
	}
	for _, tc := range cases {
		if actual := CollatzLength(tc.Input); actual != tc.Expected {
			t.Errorf("expected:%v, actual:%v", tc.Expected, actual)
		}
	}
}

func TestFibonacci(t *testing.T) {
	cases := []struct {
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
	for _, tc := range cases {
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

func TestSumSlice(t *testing.T) {
	cases := []struct {
		InputSlice     []int
		InputConvertor func(n int) int
		Expected       int
	}{
		{[]int{1, 2, 3, 4}, func(n int) int { return 2 * n }, 20},
	}
	for _, tc := range cases {
		if actual := SumSlice(tc.InputSlice, tc.InputConvertor); actual != tc.Expected {
			t.Errorf(
				"InputSlice:%v\nInputConvertor:%v\nExpected:%v\nActual%v",
				tc.InputSlice,
				tc.InputConvertor,
				tc.Expected,
				actual,
			)
		}
	}
}

func TestmultiplyDigits(t *testing.T) {
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
