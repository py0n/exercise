package project_euler

import "testing"

func TestPE0011(t *testing.T) {
	cases := []struct {
		Expected int
	}{
		{70600674},
	}
	for _, tc := range cases {
		if actual := PE0011(); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
