package euler

import (
	"strconv"

	"github.com/pkg/errors"
)

// https://projecteuler.net/problem=8

// PE0008 13の隣接する數字の積の内、最大の物を計算する。
func PE0008(n int, grid string) (int, error) {
	maxNum := 0
	for i := 0; i < len(grid)-n-1; i++ {
		s := grid[i : i+n]
		m, err := strconv.Atoi(s)
		if err != nil {
			return 0, errors.Wrap(err, s+" is invalid digits")
		}
		product := multiplyDigits(m)
		if product > maxNum {
			maxNum = product
		}
	}
	return maxNum, nil
}
