package project_euler

import (
	"math/big"
)

// https://projecteuler.net/problem=15
//
// Benchmark_PE0015-4                        300000              9816 ns/op            3440 B/op         86 allocs/op
// Benchmark_PE0015Dp-4                      100000             12500 ns/op            4096 B/op          1 allocs/op
// Benchmark_PE0015Memoization-4             300000              5701 ns/op            4096 B/op          1 allocs/op
// PASS
// ok      github.com/py0n/project_euler   6.166s

// PE0015 組み合はせを使用して計算 (bigも使用)
//
// +-+-+-+
// | | | |
// +-+-+-+
// | | | |
// +-+-+-+
// | | | |
// +-+-+-+
//
// |  | 1| 2| 3| 4| 5|
// | 1| 2| 3| 4| 5| 6|
// | 2| 3| 6|10|15|
// | 3| 4|10|20|  |
// | 4| 5|15|  |  |
// | 5| 6|
//
// n * m -> (n+m)Cn
func PE0015(n int) int64 {
	numerator := big.NewInt(1) // 分子
	for i := 0; i < n; i++ {
		numerator.Mul(numerator, big.NewInt(int64(2*n-i)))
	}
	denominator := big.NewInt(1) //分母
	for i := 0; i < n; i++ {
		denominator.Mul(denominator, big.NewInt(int64(n-i)))
	}
	r := big.NewInt(0)
	r.Quo(numerator, denominator)
	return r.Int64()
}

// PE0015Dp DPを使用して計算
//
// + - 1 - 1 - 1
// |   |   |   |
// 1 - 2 - 3 - 4
// |   |   |   |
// 1 - 3 - 6 -10
// |   |   |   |
// 1 - 4 -10 -20
func PE0015Dp(n int) int64 {
	m := n + 1 // 一辺の點の數
	grid := make([]int64, m*m)
	grid[0] = 1 // 始點
	for i := 1; i < m*m; i++ {
		// 左の點が存在するなら、其處迄の經路數を加算
		if j := i % m; j > 0 {
			grid[i] += grid[i-1]
		}
		// 上の點が存在するなら、其處までの經路數を加算
		if j := i - m; j >= 0 {
			grid[i] += grid[i-m]
		}
	}
	return grid[m*m-1]
}

// PE0015Memoization メモ化を利用して計算 (再歸を試したが計算が終了しなかつた)
func PE0015Memoization(n int) int64 {
	memo := make([]int64, (n+1)*(n+1))

	var routeNumber func(int, int) int64
	routeNumber = func(i, j int) int64 {
		if i < 0 || j < 0 {
			return 0
		} else if i == 0 && j == 0 {
			return 1
		}

		p0 := j*(n+1) + i
		if memo[p0] > 0 {
			return memo[p0]
		}
		p1 := i*(n+1) + j // 對稱性から
		if memo[p1] > 0 {
			return memo[p1]
		}

		memo[p0] = routeNumber(i-1, j) + routeNumber(i, j-1)

		return memo[p0]
	}

	return routeNumber(n, n)
}
