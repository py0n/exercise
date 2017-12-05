package project_euler

// https://projecteuler.net/problem=14

// PE0014a 1,000,000以下の数字でCollatz Sequenceが最も長いものを計算
func PE0014a(n int) int {
	maxLength := 0      // Collatz Sequenceの最大の長さ
	startingNumber := 0 // その時の開始数値
	for i := 1; i < n; i++ {
		if l := CollatzLength(i); l > maxLength {
			maxLength = l
			startingNumber = i
		}
	}
	return startingNumber
}

/*
計算済の数字をマップで記録したら寧ろ遅くなった。
メモリ確保の為。
=== RUN   Test_PE0014
--- PASS: Test_PE0014 (0.38s)
=== RUN   Test_PE0014_2
--- PASS: Test_PE0014_2 (2.84s)
Benchmark_PE0014-2         10000           2470412 ns/op               0 B/op          0 allocs/op
Benchmark_PE0014_2-2       10000          14004803 ns/op          854254 B/op       1177 allocs/op
PASS
ok      github.com/py0n/project_euler   167.988s
*/

// PE0014b 1,000,000以下の数字でCollatz Sequenceが最も長いものを計算(2)
func PE0014b(n int) int {
	trush := map[int]bool{}
	maxLength := 0
	startingNumber := 0
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
			startingNumber = i
		}
	}
	return startingNumber
}
