package project_euler

func PE0007(n int) int {
	primeGenerator := NewPrimeGenerator()
	for i := 0; i < n-1; i++ {
		primeGenerator.Next()
	}
	return primeGenerator.Next()
}
