package euler

// Lcm 最小公倍數を計算する
func Lcm(n, m int) int {
	g := Gcm(n, m)
	if g == 0 {
		return 0
	} else if g == 1 {
		return n * m
	}
	return n * m / g
}
