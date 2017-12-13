package euler

import "testing"

func TestProductIntSlice(t *testing.T) {
	cases := []struct {
		InputSlice []int
		Expected   int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{1, 2, 3, 4}, 24},
	}
	for _, tc := range cases {
		if actual := ProductIntSlice(tc.InputSlice); actual != tc.Expected {
			t.Errorf(
				"InputSlice:%v\nExpected:%v\nActual%v",
				tc.InputSlice,
				tc.Expected,
				actual,
			)
		}
	}
}

func TestSumIntSlice(t *testing.T) {
	cases := []struct {
		InputSlice []int
		Expected   int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{1, 2, 3, 4}, 10},
	}
	for _, tc := range cases {
		if actual := SumIntSlice(tc.InputSlice); actual != tc.Expected {
			t.Errorf(
				"InputSlice:%v\nExpected:%v\nActual%v",
				tc.InputSlice,
				tc.Expected,
				actual,
			)
		}
	}
}
