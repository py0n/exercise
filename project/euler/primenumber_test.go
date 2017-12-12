package euler

import (
	"reflect"
	"testing"
)

func BenchmarkPrimeGenerator0(b *testing.B) {
	PrimeGenerator := NewPrimeGenerator0()
	for i := 0; i < b.N; i++ {
		PrimeGenerator.Next()
	}
}

func BenchmarkPrimeGenerator1(b *testing.B) {
	PrimeGenerator := NewPrimeGenerator1()
	for i := 0; i < b.N; i++ {
		PrimeGenerator.Next()
	}
}

func BenchmarkPrimeGenerator(b *testing.B) {
	PrimeGenerator := NewPrimeGenerator()
	for i := 0; i < b.N; i++ {
		PrimeGenerator.Next()
	}
}

func TestPrimeGenerator(t *testing.T) {
	PrimeGenerator := NewPrimeGenerator()

	for _, expected := range []int{2, 3, 5, 7, 11, 13, 17, 19, 23} {
		if actual := PrimeGenerator.Next(); actual != expected {
			t.Errorf("expected=%v, actual=%v", expected, actual)
		}
	}
}

func TestPrimeFactorize(t *testing.T) {
	invalidCases := []struct {
		Input int
	}{
		// 異常系: < 2
		{-1},
		{0},
		{1},
	}
	for _, tc := range invalidCases {
		if actual := PrimeFactorize(tc.Input); actual != nil {
			t.Errorf("expected=%v, actual=%v", nil, actual)
		}
	}
	validCases := []struct {
		Input    int
		Expected map[int]int
	}{
		// 正常系
		{2, map[int]int{2: 1}},
		{3, map[int]int{3: 1}},
		{4, map[int]int{2: 2}},
		{5, map[int]int{5: 1}},
		{17, map[int]int{17: 1}},
		{60, map[int]int{2: 2, 3: 1, 5: 1}},
	}
	for _, tc := range validCases {
		if actual := PrimeFactorize(tc.Input); !reflect.DeepEqual(actual, tc.Expected) {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
