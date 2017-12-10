package euler

import "testing"

var pe0006Cases = []struct {
	Input    int
	Expected int
}{
	{10, 2640},
	{100, 25164150},
}

func TestPE0006(t *testing.T) {
	for _, tc := range pe0006Cases {
		if actual := PE0006(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}
