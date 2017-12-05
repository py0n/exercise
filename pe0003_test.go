package project_euler

import "testing"

var pe0003Cases = []struct {
	Input    int
	Expected int
}{
	{13195, 29},
	//{600851475143, 6857},
}

func TestPE0003SortSort(t *testing.T) {
	for _, tc := range pe0003Cases {
		if actual, _ := PE0003SortSort(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Benchmark_PE0003SortSort(b *testing.B) {
	b.ResetTimer()
	PE0003SortSort(b.N)
}

func TestPE0003SortSlice(t *testing.T) {
	for _, tc := range pe0003Cases {
		if actual, _ := PE0003SortSlice(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Benchmark_PE0003SortSlice(b *testing.B) {
	b.ResetTimer()
	PE0003SortSlice(b.N)
}
