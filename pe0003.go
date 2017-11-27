package project_euler

import (
	"math"
	"sort"
)

// https://projecteuler.net/problem=3
func PE0003(n int) int {
	primeNumbers := CreatePrimeArray(int(math.Sqrt(float64(n))))

	sort.Sort(sort.Reverse(sort.IntSlice(primeNumbers)))

	for _, p := range primeNumbers {
		if n%p == 0 {
			return p
		}
	}

	return 0
}
