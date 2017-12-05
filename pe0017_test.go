package project_euler

import "testing"

var pe0017Cases = []struct {
	Input    int
	Expected int
}{
	{1000, 21124},
}

func Test_PE0017SortSort(t *testing.T) {
	for _, tc := range pe0017Cases {
		actual, err := PE0017SortSort(tc.Input)
		if err != nil {
			t.Fatal(err)
		} else if actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Benchmark_PE0017SortSort(b *testing.B) {
	PE0017SortSort(b.N)
}

func Test_PE0017SortSlice(t *testing.T) {
	for _, tc := range pe0017Cases {
		actual, err := PE0017SortSort(tc.Input)
		if err != nil {
			t.Fatal(err)
		} else if actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func Benchmark_PE0017SortSlice(b *testing.B) {
	PE0017SortSlice(b.N)
}
