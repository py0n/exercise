package euler

import "testing"

// PE0026 {{{
var pe0026Cases = []struct {
	Input    int
	Expected int
}{
	{1000, 983},
}

func TestPE0026(t *testing.T) {
	for _, tc := range pe0026Cases {
		if actual := PE0026(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

/*
Project Euler謚慕ｨｿ逕ｨ
func TestPE0026b(t *testing.T) {

	recurringCycleLength := func(n int) int {
		qs := []int{}
		rs := []int{}
		for q, r := 0, 10; r > 0; q, r = r/n, (r%n)*10 {
			for i := 0; i < len(qs); i++ {
				if qs[i] != q || rs[i] != r {
					continue
				}
				return len(qs[i:])
			}
			qs = append(qs, q)
			rs = append(rs, r)
		}
		return 0
	}

	result, maxLength := 0, 0

	for n := 1; n < 1000; n++ {
		if l := recurringCycleLength(n); l > maxLength {
			result, maxLength = n, l
		}
	}

	if result != 983 {
		t.Errorf("Expected:%v, Result:%v", 983, result)
	}
}
*/
// }}}

// PE0027 {{{
var pe0027Cases = []struct {
	InputA   int
	InputB   int
	Expected int
}{
	{1000, 1000, -59231},
}

func TestPE0027a(t *testing.T) {
	for _, tc := range pe0027Cases {
		if actual := PE0027a(tc.InputA, tc.InputB); actual != tc.Expected {
			t.Errorf(
				"InputA:%v\nInputB:%v\nExpected:%v\nActual:%v",
				tc.InputA,
				tc.InputB,
				tc.Expected,
				actual,
			)
		}
	}
}

// }}}

// vim:set foldmethod=marker: