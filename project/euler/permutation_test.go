package euler

// 順列生成のテスト

import (
	"reflect"
	"testing"
)

func TestPermutation(t *testing.T) {
	for _, tc := range testPermutationCases2 {
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

func TestPermutationA(t *testing.T) {
	for _, tc := range testPermutationCases {
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

func BenchmarkPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Permutation([]rune("123456789"), 9)
	}
}

func BenchmarkPermutationA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PermutationA("123456789", 9)
	}
}

func BenchmarkPermutationInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PermutationInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9)
	}
}

func TestPermutationInt(t *testing.T) {
	for _, tc := range testPermutationIntCases {
		if actual := PermutationInt(tc.InputIntSlice, tc.InputK); !reflect.DeepEqual(actual, tc.Expected) {
			t.Errorf(
				"InputString:%v\nInputK:%v\nExpected:%v\nActual:%v",
				tc.InputIntSlice,
				tc.InputK,
				tc.Expected,
				actual,
			)
		}
	}
}

func TestPermutationNumber(t *testing.T) {
	for _, tc := range testPermutationNumberCases {
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

var testPermutationCases = []struct {
	InputString string
	InputK      int
	Expected    []string
}{
	{"abc", 2, []string{
		"ab", "ac",
		"ba", "bc",
		"ca", "cb",
	}},
	{"abc", 3, []string{
		"abc", "acb",
		"bac", "bca",
		"cab", "cba",
	}},
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

var testPermutationCases2 = []struct {
	InputString []rune
	InputK      int
	Expected    [][]rune
}{
	{[]rune("abc"), 2, [][]rune{
		[]rune("ab"), []rune("ac"),
		[]rune("ba"), []rune("bc"),
		[]rune("ca"), []rune("cb"),
	}},
	{[]rune("abc"), 3, [][]rune{
		[]rune("abc"), []rune("acb"),
		[]rune("bac"), []rune("bca"),
		[]rune("cab"), []rune("cba"),
	}},
	{[]rune("abcd"), 2, [][]rune{
		[]rune("ab"), []rune("ac"), []rune("ad"),
		[]rune("ba"), []rune("bc"), []rune("bd"),
		[]rune("ca"), []rune("cb"), []rune("cd"),
		[]rune("da"), []rune("db"), []rune("dc"),
	}},
	{[]rune("abcd"), 3, [][]rune{
		[]rune("abc"), []rune("abd"), []rune("acb"), []rune("acd"), []rune("adb"), []rune("adc"),
		[]rune("bac"), []rune("bad"), []rune("bca"), []rune("bcd"), []rune("bda"), []rune("bdc"),
		[]rune("cab"), []rune("cad"), []rune("cba"), []rune("cbd"), []rune("cda"), []rune("cdb"),
		[]rune("dab"), []rune("dac"), []rune("dba"), []rune("dbc"), []rune("dca"), []rune("dcb"),
	}},
	// 異常系
	{[]rune("abc"), 0, nil},  // K<1
	{[]rune("abc"), -1, nil}, // K<1
	{[]rune("abc"), 4, nil},  // K>len(String)
	{[]rune("abcd"), 5, nil}, // K>len(String)
}

var testPermutationIntCases = []struct {
	InputIntSlice []int
	InputK        int
	Expected      [][]int
}{
	{[]int{1, 2, 3}, 2, [][]int{
		[]int{1, 2}, []int{1, 3},
		[]int{2, 1}, []int{2, 3},
		[]int{3, 1}, []int{3, 2},
	}},
	{[]int{1, 2, 3}, 3, [][]int{
		[]int{1, 2, 3}, []int{1, 3, 2},
		[]int{2, 1, 3}, []int{2, 3, 1},
		[]int{3, 1, 2}, []int{3, 2, 1},
	}},
	{[]int{1, 2, 3, 4}, 2, [][]int{
		[]int{1, 2}, []int{1, 3}, []int{1, 4},
		[]int{2, 1}, []int{2, 3}, []int{2, 4},
		[]int{3, 1}, []int{3, 2}, []int{3, 4},
		[]int{4, 1}, []int{4, 2}, []int{4, 3},
	}},
	{[]int{1, 2, 3, 4}, 3, [][]int{
		[]int{1, 2, 3}, []int{1, 2, 4}, []int{1, 3, 2}, []int{1, 3, 4}, []int{1, 4, 2}, []int{1, 4, 3},
		[]int{2, 1, 3}, []int{2, 1, 4}, []int{2, 3, 1}, []int{2, 3, 4}, []int{2, 4, 1}, []int{2, 4, 3},
		[]int{3, 1, 2}, []int{3, 1, 4}, []int{3, 2, 1}, []int{3, 2, 4}, []int{3, 4, 1}, []int{3, 4, 2},
		[]int{4, 1, 2}, []int{4, 1, 3}, []int{4, 2, 1}, []int{4, 2, 3}, []int{4, 3, 1}, []int{4, 3, 2},
	}},
	// 異常系
	{[]int{1, 2, 3}, 0, nil},    // K<1
	{[]int{1, 2, 3}, -1, nil},   // K<1
	{[]int{1, 2, 3}, 4, nil},    // K>len{Int}
	{[]int{1, 2, 3, 4}, 5, nil}, // K>len{Int}
}

var testPermutationNumberCases = []struct {
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
