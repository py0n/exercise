package project_euler

import "testing"

func TestPE0001(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{10, 23},
		{1000, 233168},
	}
	for _, tc := range cases {
		if actual := PE0001(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
