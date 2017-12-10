package euler

import (
	"testing"
)

var pe0004Cases = []struct {
	Input    int
	Expected int
}{
	{2, 9009},
	{3, 906609},
}

func TestPE0004(t *testing.T) {
	for _, tc := range pe0004Cases {
		actual := PE0004(tc.Input)
		if actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
