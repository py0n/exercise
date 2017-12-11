package euler

import "testing"

var pe0026Cases = []struct {
	Input    int
	Expected int
}{
	{1000, 983},
}

func TestPE0026(t *testing.T) {
	for _, tc := range pe0026Cases {
		if actual := PE0026(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}
