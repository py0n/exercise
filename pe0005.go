package project_euler

// https://projecteuler.net/problem=5
func PE0005(n int) int {
	if n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result = lcm(result, i)
	}
	return result
}
