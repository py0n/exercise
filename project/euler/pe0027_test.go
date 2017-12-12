package euler

import "testing"

var pe0027Cases = []struct {
	InputA   int
	InputB   int
	Expected int
}{
	{1000, 1000, -59231},
}

func TestPE0027a(t *testing.T) {
	for _, tc := range pe0027Cases {
		if actual := PE0027a(tc.InputA, tc.InputB); actual != tc.Expected {
			t.Errorf(
				"InputA:%v\nInputB:%v\nExpected:%v\nActual:%v",
				tc.InputA,
				tc.InputB,
				tc.Expected,
				actual,
			)
		}
	}
}
