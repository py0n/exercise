package project_euler

import "testing"

func TestPE0005(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{10, 2520},
		{20, 232792560},
	}
	for _, tc := range cases {
		if actual := PE0005(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
