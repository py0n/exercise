package euler

// https://projecteuler.net/problem=7

// PE0007 10001番目の素數を計算する。
func PE0007(n int) int {
	primeGenerator := NewPrimeGenerator()
	for i := 0; i < n-1; i++ {
		primeGenerator.Next()
	}
	return primeGenerator.Next()
}
