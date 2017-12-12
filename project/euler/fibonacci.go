package euler

import "math/big"

// Fibonacci n番目のFibonacci數を計算
func Fibonacci(nth int, fibonacciMap map[int]*big.Int) *big.Int {
	if nth == 0 || nth == 1 {
		return big.NewInt(int64(1))
	}
	if fibonacciMap == nil {
		fibonacciMap = make(map[int]*big.Int)
	}
	if _, ok := fibonacciMap[nth]; ok {
		return fibonacciMap[nth]
	}
	fibonacciMap[nth] = big.NewInt(int64(0)).Add(
		Fibonacci(nth-1, fibonacciMap),
		Fibonacci(nth-2, fibonacciMap),
	)
	return fibonacciMap[nth]
}
