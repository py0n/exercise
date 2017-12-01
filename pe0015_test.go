package project_euler

import "testing"

func Test_PE0015(t *testing.T) {
	cases := []struct {
		Input    int64
		Expected int64
	}{
		{2, 6},
		{3, 20},
		{20, 137846528820},
	}
	for _, tc := range cases {
		if actual := PE0015(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
