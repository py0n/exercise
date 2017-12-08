package euler

import (
	"math/big"
)

// https://projecteuler.net/problem=25

// PE0025a 初めて1000桁になるFibonacci数の順番を計算する
func PE0025a(n int) int {

	fibonacciMap := map[int]*big.Int{} // n番目のFibonacci数

	var fibonacci func(int) *big.Int

	fibonacci = func(nth int) *big.Int {
		if nth == 0 || nth == 1 {
			return big.NewInt(int64(1))
		} else if _, ok := fibonacciMap[nth]; ok {
			return fibonacciMap[nth]
		}
		fibonacciMap[nth] = big.NewInt(int64(0)).Add(
			fibonacci(nth-1),
			fibonacci(nth-2),
		)
		return fibonacciMap[nth]
	}

	m := 1
	for ; len(fibonacci(m).String()) < n; m++ {
	}
	return m + 1
}

// PE0025b クロージャを使用せず、初めて1000桁になるFibonacci數を計算
func PE0025b(n int) int {
	fibonacciMap := map[int]*big.Int{} // n番目のFibonacci数

	m := 1
	for len(Fibonacci(m, fibonacciMap).String()) < n {
		m++
	}
	return m + 1
}

// PE0025c 更に共有マップを使用しない
func PE0025c(n int) int {
	m := 1
	for len(Fibonacci(m, nil).String()) < n {
		m++
	}
	return m + 1
}
