package euler

import "testing"

var pe0007Cases = []struct {
	Input    int
	Expected int
}{
	{1, 2},
	{2, 3},
	{3, 5},
	{4, 7},
	{10001, 104743},
}

func TestPE0007(t *testing.T) {
	for _, tc := range pe0007Cases {
		if actual := PE0007(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
