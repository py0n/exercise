package euler

import "testing"

var pe0001Cases = []struct {
	Input    int
	Expected int
}{
	{10, 23},
	{1000, 233168},
}

func TestPE0001(t *testing.T) {
	for _, tc := range pe0001Cases {
		if actual := PE0001(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
