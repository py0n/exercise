package project_euler

// https://projecteuler.net/problem=6
func PE0006(n int) int {
	sum0, sum1 := 0, 0
	for i := 1; i <= n; i++ {
		sum0 = sum0 + i*i
		sum1 = sum1 + i
	}
	return sum1*sum1 - sum0
}
