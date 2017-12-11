package euler

// https://projecteuler.net/problem=6

// PE0006 100迄の和の二乘と、二乘の和の差を計算する。
func PE0006(n int) int {
	sum0, sum1 := 0, 0
	for i := 1; i <= n; i++ {
		sum0 = sum0 + i*i
		sum1 = sum1 + i
	}
	return sum1*sum1 - sum0
}
