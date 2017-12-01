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
