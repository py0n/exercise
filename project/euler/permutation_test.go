package euler

import (
	"reflect"
	"testing"
)

func TestPermutation(t *testing.T) {
	cases := []struct {
		InputString string
		InputK      int
		Expected    []string
	}{
		{"abc", 2, []string{"ab", "ac", "ba", "bc", "ca", "cb"}},
		{"abc", 3, []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{"abcd", 2, []string{"ab", "ac", "ad", "ba", "bc", "bd", "ca", "cb", "cd", "da", "db", "dc"}},
		{"abcd", 3, []string{
			"abc", "abd", "acb", "acd", "adb", "adc",
			"bac", "bad", "bca", "bcd", "bda", "bdc",
			"cab", "cad", "cba", "cbd", "cda", "cdb",
			"dab", "dac", "dba", "dbc", "dca", "dcb",
		}},
	}
	for _, tc := range cases {
		if actual := Permutation(tc.InputString, tc.InputK); !reflect.DeepEqual(actual, tc.Expected) {
			t.Errorf(
				"InputString:%v\nInputK:%v\nExpected:%v\nActual:%v",
				tc.InputString,
				tc.InputK,
				tc.Expected,
				actual,
			)
		}
	}
}
