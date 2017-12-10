package euler

// https://projecteuler.net/problem=10

// PE0010 2,000,000以下の素數の和を計算する。
func PE0010(n int) int {
	sum := 0

	primeGenerator := NewPrimeGenerator()
	for p := 0; p < n; p = primeGenerator.Next() {
		sum += p
	}

	return sum
}
