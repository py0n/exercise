package project_euler

import (
	"errors"
	"math"
	"sort"
)

// https://projecteuler.net/problem=3
func PE0003(n int) (int, error) {
	primeNumbers := []int{}

	uLimit := int(math.Sqrt(float64(n)))

	primeGenerator := NewPrimeGenerator()
	for {
		m := primeGenerator.Next()
		if m > uLimit {
			break
		}
		primeNumbers = append(primeNumbers, m)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(primeNumbers)))

	for _, p := range primeNumbers {
		if n%p == 0 {
			return p, nil
		}
	}

	return 0, errors.New("no solution")
}
