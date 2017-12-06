package project_euler

// https://projecteuler.net/problem=5

// PE0005 n迄の數の最小公倍數を計算する
func PE0005(n int) int {
	if n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result = Lcm(result, i)
	}
	return result
}
