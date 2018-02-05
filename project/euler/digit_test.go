package euler

import "testing"

func TestCalculateDigitNumber(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{-10, 2},
		{-1, 1},
		{0, 1},
		{1, 1},
		{3, 1},
		{10, 2},
		{41, 2},
		{100, 3},
		{213, 3},
	}
	for _, tc := range cases {
		if actual := CalculateDigitNumber(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

func TestCirculateDigit(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		// 負數
		{-1, -1}, {-10, -1}, {-12, -21},
		// 0
		{0, 0},
		// 正數
		{1, 1}, {2, 2},
		{10, 1}, {12, 21},
		{100, 10}, {102, 210}, {120, 12}, {123, 312},
		{1000, 100}, {1002, 2100}, {1020, 102}, {1200, 120},
		{1023, 3102}, {1203, 3120}, {1230, 123}, {1234, 4123},
	}
	for _, tc := range cases {
		if actual := CirculateDigit(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}
