package euler

// 順列生成のテスト

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

func BenchmarkPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Permutation("123456789", 9)
	}
}

func TestPermutationA(t *testing.T) {
	cases := []struct {
		InputString string
		InputK      int
		Expected    []string
	}{
		{"abc", 2, []string{"ab", "ac", "ba", "bc", "ca", "cb"}},
		{"abc", 3, []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{"abcd", 2, []string{
			"ab", "ac", "ad",
			"ba", "bc", "bd",
			"ca", "cb", "cd",
			"da", "db", "dc",
		}},
		{"abcd", 3, []string{
			"abc", "abd", "acb", "acd", "adb", "adc",
			"bac", "bad", "bca", "bcd", "bda", "bdc",
			"cab", "cad", "cba", "cbd", "cda", "cdb",
			"dab", "dac", "dba", "dbc", "dca", "dcb",
		}},
		// 異常系
		{"abc", 0, nil},  // K<1
		{"abc", -1, nil}, // K<1
		{"abc", 4, nil},  // K>len(String)
		{"abcd", 5, nil}, // K>len(String)
	}
	for _, tc := range cases {
		if actual := PermutationA(tc.InputString, tc.InputK); !reflect.DeepEqual(actual, tc.Expected) {
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

func BenchmarkPermutationA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PermutationA("123456789", 9)
	}
}

func TestPermutationNumber(t *testing.T) {
	cases := []struct {
		InputM   int
		InputN   int
		Expected int
	}{
		{1, 1, 1},
		{10, 1, 10},
		{10, 4, 10 * 9 * 8 * 7},
		{10, 10, 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1},
		// 異常系
		{0, 2, 0}, // m < 1
		{2, 0, 0}, // n < 1
		{2, 3, 0}, // m < n
	}
	for _, tc := range cases {
		if actual := PermutationNumber(tc.InputM, tc.InputN); actual != tc.Expected {
			t.Errorf(
				"InputM:%v\nInputN:%v\nExpected:%v\nActual:%v",
				tc.InputM,
				tc.InputN,
				tc.Expected,
				actual,
			)
		}
	}
}
