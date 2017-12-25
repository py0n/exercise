package euler

import "testing"

func TestIsPandigital10(t *testing.T) {
	cases := []struct {
		Input    int
		Expected bool
	}{
		{123456789, true},
		{135792468, true},
		{987654321, true},
		{123456788, false}, // < 123456789
		{987654322, false}, // > 987654321
		{977654321, false}, // duplicated
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
