package euler

import (
	"testing"
)

var pe0022Cases = []struct {
	Input    string
	Expected int
}{
	{"p022_test.txt", 6*1 + 2*9 + 12*3},
	{"p022_names.txt", 871198282},
}

func TestPE0022(t *testing.T) {
	t.Skip("skip... test file is not included in repository")
	for _, tc := range pe0022Cases {
		if actual, _ := PE0022(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

func Test_wordWorth(t *testing.T) {
	cases := []struct {
		Input    string
		Expected int
	}{
		{"COLIN", 53},
	}
	for _, tc := range cases {
		if actual := wordWorth(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input: %v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}
