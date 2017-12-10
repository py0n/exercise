package euler

import "testing"

var pe0016Cases = []struct {
	Input    int
	Expected int
}{
	{15, 26},
	{1000, 1366},
}

func TestPE0016(t *testing.T) {
	for _, tc := range pe0016Cases {
		if actual := PE0016(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
