package euler

import "testing"

var pe0008Cases = []struct {
	Input    int
	Expected int
}{
	{4, 5832},
	{13, 23514624000},
}

func TestPE0008(t *testing.T) {
	for _, tc := range pe0008Cases {
		if actual, _ := PE0008(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
