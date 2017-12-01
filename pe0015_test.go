package project_euler

import "testing"

func Test_PE0015(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{0, 0},
	}
	for _, tc := range cases {
		if actual := PE0015(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
