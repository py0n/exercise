package euler

import (
	"errors"
)

// https://projecteuler.net/problem=2

// PE0002 4,000,000を越えないFibonacci數の内、偶數である物の總和を計算
func PE0002(n int) (int, error) {
	if n < 3 {
		return 0, errors.New("n must be >= 3")
	}
	i, j, result := 1, 1, 0
	for ; j < n; i, j = j, i+j {
		if j%2 == 0 {
			result = result + j
		}
	}
	return result, nil
}
