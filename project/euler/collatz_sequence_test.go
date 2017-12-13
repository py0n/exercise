package euler

import "testing"

func TestCollatzLength(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		// 逡ｰ蟶ｸ邉ｻ: < 2
		{-1, 0},
		{0, 0},
		{1, 0},
		// 豁｣蟶ｸ邉ｻ
		{4, 3},
		{13, 10},
	}
	for _, tc := range cases {
		if actual := CollatzLength(tc.Input); actual != tc.Expected {
			t.Errorf("expected:%v, actual:%v", tc.Expected, actual)
		}
	}
}
