package project_euler

import (
	"strconv"
)

// https://projecteuler.net/problem=18

// PE0018Memoization Problem 18をメモ化を使用して計算
func PE0018Memoization(size int, data []string) int {
	memo := make([]int, size*(size+1)/2) // メモ

	var routeSum func(int, int) int
	routeSum = func(i, j int) int {
		if i < 0 || j < 0 || i > j {
			return 0
		}

		p := j*(j+1)/2 + i
		if memo[p] > 0 {
			return memo[p]
		}

		m, _ := strconv.Atoi(data[p])

		lv := routeSum(i-1, j-1)
		rv := routeSum(i, j-1)

		if lv > rv {
			memo[p] = m + lv
		} else {
			memo[p] = m + rv
		}
		return memo[p]
	}

	nMax := 0
	for i := 0; i < size; i++ {
		if routeSum(i, size-1) > nMax {
			nMax = routeSum(i, size-1)
		}
	}
	return nMax
}
