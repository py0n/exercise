package project_euler

import "testing"

func TestPE0002(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{3, 2},
		{10, 10},
		{4000000, 4613732},
	}
	for _, tc := range cases {
		if actual, _ := PE0002(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
