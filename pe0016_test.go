package project_euler

import "testing"

func TestPE0016(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{15, 26},
		{1000, 1366},
	}
	for _, tc := range cases {
		if actual := PE0016(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
