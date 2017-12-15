package euler

import "testing"

func TestZellerWeekday(t *testing.T) {
	cases := []struct {
		InputY   int
		InputM   int
		InputD   int
		Expected DayOfWeek
	}{
		// 異常系: グレゴリオ歴以前
		{1582, 10, 14, None},
		{1580, 1, 1, None},
		// 正常系
		{1582, 10, 15, Friday},
		{1904, 1, 4, Monday},
		{1904, 4, 5, Tuesday},
	}
	for _, tc := range cases {
		if actual := ZellerWeekday(tc.InputY, tc.InputM, tc.InputD); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
