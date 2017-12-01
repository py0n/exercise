package project_euler

import "testing"

func Test_PE0017(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{1000, 21124},
	}
	for _, tc := range cases {
		actual, err := PE0017(tc.Input)
		if err != nil {
			t.Fatal(err)
		} else if actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
