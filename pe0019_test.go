package project_euler

import "testing"

var pe0019Cases = []struct {
	Expeceted int
}{
	{10},
}

func Test_PE0019(t *testing.T) {
	for _, tc := range pe0019Cases {
		if actual := PE0019(); actual != tc.Expeceted {
			t.Errorf("expected=%v, actual=%v", tc.Expeceted, actual)
		}
	}
}
