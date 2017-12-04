package project_euler

import "testing"

var pe0012Cases = []struct {
	Input    int
	Expected int
}{
	{5, 28},
	{500, 76576500},
}

func Test_PE0012(t *testing.T) {
	for _, tc := range pe0012Cases {
		if actual := PE0012(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Benchmark_PE0012(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 10; i++ {
		PE0012(500)
	}
}

func Test_PE0012a(t *testing.T) {
	for _, tc := range pe0012Cases {
		if actual := PE0012a(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Benchmark_PE0012a(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 10; i++ {
		PE0012a(500)
	}
}
