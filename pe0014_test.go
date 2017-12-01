package project_euler

import "testing"

func Test_PE0014(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{10000000, 686},
	}
	for _, tc := range cases {
		if actual := PE0014(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Benchmark_PE0014(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0014(b.N)
	}
}

/*
func Test_PE0014_2(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{10000000, 686},
	}
	for _, tc := range cases {
		if actual := PE0014_2(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
*/

func Benchmark_PE0014_2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0014_2(b.N)
	}
}

func Test_countCollatzSequenceLength(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{8, 4},
		{13, 10},
	}
	for _, tc := range cases {
		if actual := countCollatzSequenceLength(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
