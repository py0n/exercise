package euler

import "testing"

var pe0024Cases = []struct {
	Input    int
	Expected string
}{
	{1000000, "2783915460"},
}

func TestPE0024a(t *testing.T) {
	for _, tc := range pe0024Cases {
		if actual := PE0024a(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

func TestPE0024b(t *testing.T) {
	for _, tc := range pe0024Cases {
		if actual := PE0024b(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}
