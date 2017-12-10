package euler

import "testing"

var pe0010Cases = []struct {
	Input    int
	Expected int
}{
	{10, 17},
	{2000000, 142913828922},
}

func TestPE0010(t *testing.T) {
	for _, tc := range pe0010Cases {
		if actual := PE0010(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
