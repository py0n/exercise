package project_euler

import "testing"

func Test_PE0007(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{1, 2},
		{2, 3},
		{3, 5},
		{4, 7},
		{10001, 104743},
	}
	for _, tc := range cases {
		if actual := PE0007(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
