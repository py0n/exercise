package project_euler

import "testing"

func Test_PE0014a(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{1000000, 837799},
	}
	for _, tc := range cases {
		if actual := PE0014a(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Benchmark_PE0014a(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0014a(b.N)
	}
}

func Test_PE0014b(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{1000000, 837799},
	}
	for _, tc := range cases {
		if actual := PE0014b(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Benchmark_PE0014b(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0014b(b.N)
	}
}
