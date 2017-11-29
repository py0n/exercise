package project_euler

// https://projecteuler.net/problem=10
func PE0010(n int) int {
	sum := 0

	primeGenerator := NewPrimeGenerator()
	for p := 0; p < n; p = primeGenerator.Next() {
		sum += p
	}

	return sum
}
