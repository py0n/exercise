package project_euler

import "testing"

func TestPE0003(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{13195, 29},
		{600851475143, 6857},
	}
	for _, tc := range cases {
		if actual := PE0003(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
