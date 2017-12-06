package project_euler

import "testing"

var pe0012Cases = []struct {
	Input    int
	Expected int
}{
	{5, 28},
	{500, 76576500},
}

func TestPE0012a(t *testing.T) {
	for _, tc := range pe0012Cases {
		if actual := PE0012a(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func BenchmarkPE0012a(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0012a(500)
	}
}

func TestPE0012b(t *testing.T) {
	for _, tc := range pe0012Cases {
		if actual := PE0012b(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func BenchmarkPE0012b(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0012b(500)
	}
}
