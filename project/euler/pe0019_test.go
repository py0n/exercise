package euler

import "testing"

var pe0019Cases = []struct {
	Input0    int
	Input1    int
	Expeceted int
}{
	{1904, 1904, 1},
	{1901, 2000, 171},
}

func TestPE0019(t *testing.T) {
	for _, tc := range pe0019Cases {
		if actual := PE0019(tc.Input0, tc.Input1); actual != tc.Expeceted {
			t.Errorf("expected=%v, actual=%v", tc.Expeceted, actual)
		}
	}
}
