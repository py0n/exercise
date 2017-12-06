package euler

import "testing"

var pe0014Cases = []struct {
	Input    int
	Expected int
}{
	{1000000, 837799},
}

func TestPE0014a(t *testing.T) {
	for _, tc := range pe0014Cases {
		if actual := PE0014a(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func BenchmarkPE0014a(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0014a(b.N)
	}
}

func TestPE0014b(t *testing.T) {
	for _, tc := range pe0014Cases {
		if actual := PE0014b(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func BenchmarkPE0014b(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0014b(b.N)
	}
}
