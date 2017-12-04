package project_euler

import "testing"

var pe0021Cases = []struct {
	Input    int
	Expected int
}{
	{-1, 0},
	{0, 0},
	{1, 0},
	{2, 0},
	{300, 504},
	{10000, 31626},
}

func Test_PE0021(t *testing.T) {
	for _, tc := range pe0021Cases {
		if actual := PE0021(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Benchmark_PE0021(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE0021(1000)
	}
}

func Test_PE0021a(t *testing.T) {
	for _, tc := range pe0021Cases {
		if actual := PE0021a(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Benchmark_PE0021a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE0021a(1000)
	}
}
