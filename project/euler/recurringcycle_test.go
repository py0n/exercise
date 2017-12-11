package euler

import "testing"

func TestCalculateRecurringCycle(t *testing.T) {
	cases := []struct {
		Input    int
		Expected string
	}{
		{1, ""},
		{2, ""},
		{3, "3"},
		{4, ""},
		{5, ""},
		{6, "6"},
		{7, "142857"},
		{8, ""},
		{9, "1"},
		{10, ""},
		{11, "09"},
		{13, "076923"},
		{61, "016393442622950819672131147540983606557377049180327868852459"},
		{97, "010309278350515463917525773195876288659793814432989690721649484536082474226804123711340206185567"},
	}
	for _, tc := range cases {
		if actual := CalculateRecurringCycle(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

func TestCalculateRecurringCycleLength(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{1, 0},
		{2, 0},
		{3, 1},
		{4, 0},
		{5, 0},
		{6, 1},
		{7, 6},
		{8, 0},
		{9, 1},
		{10, 0},
		{11, 2},
		{13, 6},
		{61, 60},
		{97, 96},
	}
	for _, tc := range cases {
		if actual := CalculateRecurringCycleLength(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}
