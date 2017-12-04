package project_euler

import "testing"

var pe0015Cases = []struct {
	Input    int
	Expected int64
}{
	{2, 6},
	{3, 20},
	{20, 137846528820},
}

func Test_PE0015(t *testing.T) {
	for _, tc := range pe0015Cases {
		if actual := PE0015(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Test_PE0015Dp(t *testing.T) {
	for _, tc := range pe0015Cases {
		if actual := PE0015Dp(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Test_PE0015Memoization(t *testing.T) {
	for _, tc := range pe0015Cases {
		if actual := PE0015Memoization(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Benchmark_PE0015(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE0015(20)
	}
}

func Benchmark_PE0015Dp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE0015Dp(20)
	}
}

func Benchmark_PE0015Memoization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE0015Memoization(20)
	}
}