package euler

import "testing"

func TestIsPandigital10(t *testing.T) {
	cases := []struct {
		Input    int
		Expected bool
	}{
		{12345, false},
		{98765, false},
		{11223344, false},
		{123456789, true},
		{987654321, true},
	}
	for _, tc := range cases {
		if actual := IsPandigital10(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}
