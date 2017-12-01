package project_euler

// https://projecteuler.net/problem=14
func PE0014(n int) int {
	maxLength := 0
	for i := 1; i < n; i++ {
		if l := countCollatzSequenceLength(i); l > maxLength {
			maxLength = l
		}
	}
	return maxLength
}

func countCollatzSequenceLength(n int) int {
	c := 0
	for ; n > 1; c++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	return c + 1
}

/*
計算済の数字をマップで記録したら寧ろ遅くなった。
メモリ確保の為。
Benchmark_PE0014-2         10000           2291253 ns/op               0 B/op          0 allocs/op
Benchmark_PE0014_2-2       10000          14047469 ns/op          854232 B/op       1177 allocs/op
PASS
ok      github.com/py0n/project_euler   210.552s
*/
func PE0014_2(n int) int {
	trush := map[int]bool{}
	maxLength := 0
	for i := 1; i < n; i++ {
		if trush[i] {
			continue
		}
		c := 0
		for m := i; m > 1; c++ {
			if m%2 == 0 {
				m /= 2
			} else {
				m = 3*m + 1
			}
			trush[m] = true
		}
		if c+1 > maxLength {
			maxLength = c + 1
		}
	}
	return maxLength
}
