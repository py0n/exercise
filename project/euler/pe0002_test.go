package euler

import "testing"

var pe0002Cases = []struct {
	Input    int
	Expected int
}{
	{3, 2},
	{10, 10},
	{4000000, 4613732},
}

func TestPE0002(t *testing.T) {
	for _, tc := range pe0002Cases {
		if actual, _ := PE0002(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
