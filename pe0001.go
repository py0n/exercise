package project_euler

// https://projecteuler.net/problem=1
func PE0001(n int) int {
	result := 0
	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			result = result + i
		}
	}
	return result
}
