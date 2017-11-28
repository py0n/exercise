package project_euler

import (
	"fmt"
	"testing"
)

const N = 1000000

func Benchmark_PrimeGenerator0(b *testing.B) {
	PrimeGenerator := NewPrimeGenerator0()
	for i := 0; i < N; i++ {
		PrimeGenerator.Next()
	}
	fmt.Println(PrimeGenerator.Size())
}

func Benchmark_PrimeGenerator1(b *testing.B) {
	PrimeGenerator := NewPrimeGenerator1()
	for i := 0; i < N; i++ {
		PrimeGenerator.Next()
	}
	fmt.Println(PrimeGenerator.Size())
}

func Benchmark_PrimeGenerator(b *testing.B) {
	PrimeGenerator := NewPrimeGenerator()
	for i := 0; i < N; i++ {
		PrimeGenerator.Next()
	}
	fmt.Println(PrimeGenerator.Size())
}

func Test_isPalrindrome(t *testing.T) {
	cases := []struct {
		Input    string
		Expected bool
	}{
		{"abcde", false},
		{"abcba", true},
		{"abccba", true},
	}
	for _, tc := range cases {
		if actual := isPalindrome(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Test_pow10(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{0, 1},
		{1, 10},
		{2, 100},
		{3, 1000},
	}
	for _, tc := range cases {
		if actual := pow10(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
