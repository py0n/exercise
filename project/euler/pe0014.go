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
=== RUN   TestPE0014a
--- PASS: TestPE0014a (0.36s)
=== RUN   TestPE0014b
--- PASS: TestPE0014b (0.77s)
BenchmarkPE0014a-2        10000           2284030 ns/op               0 B/op          0 allocs/op
BenchmarkPE0014b-2        10000           4302648 ns/op          684764 B/op        627 allocs/op
PASS
ok      github.com/py0n/project_euler   67.015s
*/

// PE0014b 1,000,000以下の数字でCollatz Sequenceが最も長いものを計算(2)
func PE0014b(n int) int {

	cm := map[int]int{} // その数のCollatz Length

	collatzLength := func(n int) int {
		if n < 2 {
			return 0
		}
		cm[n] = 0
		for m := n; m > 1; cm[n]++ {
			if m%2 == 0 {
				m = m / 2
			} else {
				m = 3*m + 1
			}
			if _, ok := cm[m]; ok {
				cm[n] += cm[m]
				break
			}
		}
		return cm[n]
	}

	maxLength := 0
	startingNumber := 0

	for i := 1; i < n; i++ {
		if l := collatzLength(i); l+1 > maxLength {
			maxLength = l + 1
			startingNumber = i
		}
	}
	return startingNumber
}
