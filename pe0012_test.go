package project_euler

import "testing"

func Test_PE0012(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{5, 28},
		{500, 76576500},
	}
	for _, tc := range cases {
		if actual := PE0012(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Benchmark_PE0012(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0012(500)
	}
}

func Test_PE0012_2(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{5, 28},
		{500, 76576500},
	}
	for _, tc := range cases {
		if actual, _ := PE0012_2(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Benchmark_PE0012_2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0012_2(500)
	}
}
