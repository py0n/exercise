package euler

// https://projecteuler.net/problem=9

// PE0009 a+b+c=1000を滿たすピタゴラス數の積を計算する。
//
// a < b < c and a + b + c = 1000 の時の abc の値
// a: a < 1000/3
// c: c = 1000 - a - b
// b: b < c = 1000 - a - b → 2b < 1000 - a → b < (1000 - a)/2
//    a < b < (1000 - a) / 2
func PE0009() int {
	maxNum := 0
	for a := 1; a <= 1000/3; a++ {
		for b := a + 1; b <= (1000-a)/2; b++ {
			c := 1000 - a - b
			if c <= b {
				continue
			}
			if a*a+b*b != c*c {
				continue
			}
			m := a * b * c
			if m > maxNum {
				maxNum = m
			}
		}
	}
	return maxNum
}
