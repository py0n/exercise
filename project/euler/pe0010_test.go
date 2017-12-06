package euler

import "testing"

func TestPE0010(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{10, 17},
		{2000000, 142913828922},
	}
	for _, tc := range cases {
		if actual := PE0010(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
