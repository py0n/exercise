package project_euler

import (
	"strconv"
)

// https://projecteuler.net/problem=18

// PE0018Memoization Problem 18をメモ化を使用して計算
func PE0018Memoization(size int, data []string) int {

	var routeSum func(int, int, int, []string, []int) int
	routeSum = func(i, j, size int, data []string, memo []int) int {
		if i < 0 || j < 0 {
			return 0
		} else if i > j {
			return 0
		}
		p := j*(j+1)/2 + i
		m, _ := strconv.Atoi(data[p])
		if memo[p] > 0 {
			return memo[p]
		}
		if routeSum(i-1, j-1, size, data, memo) > routeSum(i, j-1, size, data, memo) {
			memo[p] = m + routeSum(i-1, j-1, size, data, memo)
		} else {
			memo[p] = m + routeSum(i, j-1, size, data, memo)
		}
		return memo[p]
	}

	memo := make([]int, size*(size+1)/2)
	nMax := 0
	for i := 0; i < size; i++ {
		if routeSum(i, size-1, size, data, memo) > nMax {
			nMax = routeSum(i, size-1, size, data, memo)
		}
	}
	return nMax
}
