package euler

/*
メモリの確保は矢張り重い
% go test -v ./project/euler -run PE0025 -bench PE0025 -benchmem
=== RUN   TestPE0025a
--- PASS: TestPE0025a (0.07s)
=== RUN   TestPE0025b
--- PASS: TestPE0025b (0.09s)
=== RUN   TestPE0025c
--- PASS: TestPE0025c (10.16s)
BenchmarkPE0025a-4            20          89606950 ns/op        11923335 B/op      46874 allocs/op
BenchmarkPE0025b-4            20          89584965 ns/op        11923821 B/op      46877 allocs/op
BenchmarkPE0025c-4             1        10095828900 ns/op       3396665360 B/op 23628161 allocs/op
PASS
ok      github.com/py0n/exercise/project/euler  25.186s
*/

import (
	"testing"
)

var pe0025Cases = []struct {
	Input    int
	Expected int
}{
	{1000, 4782},
}

func TestPE0025a(t *testing.T) {
	for _, tc := range pe0025Cases {
		if actual := PE0025a(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

func BenchmarkPE0025a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE0025a(1000)
	}
}

func TestPE0025b(t *testing.T) {
	for _, tc := range pe0025Cases {
		if actual := PE0025b(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

func BenchmarkPE0025b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE0025b(1000)
	}
}

func TestPE0025c(t *testing.T) {
	t.Skip()
	for _, tc := range pe0025Cases {
		if actual := PE0025c(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

func BenchmarkPE0025c(b *testing.B) {
	b.Skip()
	for i := 0; i < b.N; i++ {
		PE0025c(1000)
	}
}
