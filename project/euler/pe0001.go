package euler

// https://projecteuler.net/problem=1

// PE0001 1000以下の3と5の倍數の總和を計算
func PE0001(n int) int {
	result := 0
	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			result = result + i
		}
	}
	return result
}
