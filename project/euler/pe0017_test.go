package euler

import "testing"

var pe0017Cases = []struct {
	Input    int
	Expected int
}{
	{1000, 21124},
}

func TestPE0017SortSort(t *testing.T) {
	for _, tc := range pe0017Cases {
		actual, err := PE0017SortSort(tc.Input)
		if err != nil {
			t.Fatal(err)
		} else if actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func BenchmarkPE0017SortSort(b *testing.B) {
	b.ResetTimer()
	PE0017SortSort(b.N)
}

func TestPE0017SortSlice(t *testing.T) {
	for _, tc := range pe0017Cases {
		actual, err := PE0017SortSort(tc.Input)
		if err != nil {
			t.Fatal(err)
		} else if actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func BenchmarkPE0017SortSlice(b *testing.B) {
	b.ResetTimer()
	PE0017SortSlice(b.N)
}
