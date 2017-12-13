package euler

import "testing"

func TestProductIntSlice(t *testing.T) {
	cases := []struct {
		InputSlice     []int
		InputConvertor func(n int) int
		Expected       int
	}{
		{[]int{1, 2, 3, 4}, func(n int) int { return 2 * n }, 384},
	}
	for _, tc := range cases {
		if actual := ProductIntSlice(tc.InputSlice, tc.InputConvertor); actual != tc.Expected {
			t.Errorf(
				"InputSlice:%v\nInputConvertor:%v\nExpected:%v\nActual%v",
				tc.InputSlice,
				tc.InputConvertor,
				tc.Expected,
				actual,
			)
		}
	}
}

func TestSumIntSlice(t *testing.T) {
	cases := []struct {
		InputSlice     []int
		InputConvertor func(n int) int
		Expected       int
	}{
		{[]int{1, 2, 3, 4}, func(n int) int { return 2 * n }, 20},
	}
	for _, tc := range cases {
		if actual := SumIntSlice(tc.InputSlice, tc.InputConvertor); actual != tc.Expected {
			t.Errorf(
				"InputSlice:%v\nInputConvertor:%v\nExpected:%v\nActual%v",
				tc.InputSlice,
				tc.InputConvertor,
				tc.Expected,
				actual,
			)
		}
	}
}
