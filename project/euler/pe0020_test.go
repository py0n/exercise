package project_euler

import "testing"

var pe0020Cases = []struct {
	Input    int
	Expected int
}{
	{10, 27},
	{100, 648},
}

func TestPE0020(t *testing.T) {
	for _, tc := range pe0020Cases {
		if actual := PE0020(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
