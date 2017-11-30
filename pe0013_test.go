package project_euler

import "testing"

func Test_PE0013(t *testing.T) {
	cases := []struct {
		Expected int
	}{
		{5537376230},
	}
	for _, tc := range cases {
		if actual := PE0013(); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
