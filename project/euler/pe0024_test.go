package euler

import "testing"

var pe0024Cases = []struct {
	Input    int
	Expected string
}{
	{1000000, "2783915460"},
}

func TestPE0024(t *testing.T) {
	for _, tc := range pe0024Cases {
		if actual := PE0024(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}
