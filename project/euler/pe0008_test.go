package euler

import "testing"

func TestPE0008(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{4, 5832},
		{13, 23514624000},
	}
	for _, tc := range cases {
		if actual, _ := PE0008(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
