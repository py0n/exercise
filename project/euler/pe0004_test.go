package project_euler

import (
	"testing"
)

func TestPE0004(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{2, 9009},
		{3, 906609},
	}
	for _, tc := range cases {
		actual := PE0004(tc.Input)
		if actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
