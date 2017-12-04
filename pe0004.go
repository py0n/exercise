package project_euler

import (
	"strconv"
)

// https://projecteuler.net/problem=4

// PE0004 二つの三桁の数値の積で囘文と成る物の内、最大の物を計算する
func PE0004(n int) int {

	imax := Pow10(n) - Pow10(n-1) - 1

	max := Pow10(n) - 1

	mMax, x0, y0 := 0, 0, 0
	for i := 0; i <= imax; i++ {
		for j := 0; j <= i; j++ {
			x := max - i
			y := max - j
			if x < x0 && x < y0 && y < x0 && y < y0 {
				continue
			}
			m := x * y
			if m <= mMax {
				continue
			}
			if IsPalindrome(strconv.Itoa(m)) {
				mMax, x0, y0 = m, x, y
			}
		}
	}
	return mMax
}
