package euler

import "testing"

var pe0025Cases = []struct {
	Input    int
	Expected int
}{
	{1000, 1000},
}

func TestPE0025(t *testing.T) {
	for _, tc := range pe0025Cases {
		if actual := PE0025(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}
