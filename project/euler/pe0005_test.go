package euler

import "testing"

var pe0005Cases = []struct {
	Input    int
	Expected int
}{
	{10, 2520},
	{20, 232792560},
}

func TestPE0005(t *testing.T) {
	for _, tc := range pe0005Cases {
		if actual := PE0005(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
