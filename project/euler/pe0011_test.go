package euler

import "testing"

var pe0011Cases = []struct {
	Expected int
}{
	{70600674},
}

func TestPE0011(t *testing.T) {
	for _, tc := range pe0011Cases {
		if actual := PE0011(); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
