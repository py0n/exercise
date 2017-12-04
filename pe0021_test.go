package project_euler

import "testing"

var pe0021Cases = []struct {
	Input    int
	Expected []int
}{
	{300, []int{220, 284}},
}

func Test_PE0021(t *testing.T) {
	for _, tc := range pe0021Cases {
		if actual := PE0021(tc.Input); !loopEqualInt(actual, tc.Expected) {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func loopEqualInt(l, r []int) bool {
	if len(l) != len(r) {
		return false
	}
	min := len(l)
	if len(r) < len(l) {
		min = len(r)
	}
	for i := 0; i < min; i++ {
		if l[i] != r[i] {
			return false
		}
	}
	return true
}
