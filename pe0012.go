package project_euler

import "github.com/pkg/errors"

// http://hanatsubaki.shiseidogroup.jp/comic2/2368/
//
// n が a^x * b^y * c^z と素因数分解できる時、
// 約数の個数は(x+1)(y+1)(z+1)個である。
func PE0012(n int) int {
	for i := 1; ; i++ {
		triangleNumber := i * (i + 1) / 2 // 三角数
		primeMap := factorizePrime(triangleNumber)
		solution := 1
		for _, v := range primeMap {
			solution *= (v + 1)
		}
		if solution > n {
			return triangleNumber
		}
	}
}

func factorizePrime(n int) map[int]int {
	g := NewPrimeGenerator()
	primeMap := map[int]int{}
	for m, p := n, g.Next(); m >= p; p = g.Next() {
		for ; m%p == 0; m = m / p {
			primeMap[p]++
		}
	}
	return primeMap
}

func PE0012_2(n int) (int, error) {
	for i := 1; ; i++ {
		t := i * (i + 1) / 2     // 三角数
		g := NewPrimeGenerator() // 素数ジェネレータ
		dn := 1                  // 約数の個数
		for m, p := t, g.Next(); m >= p; p = g.Next() {
			c := 0
			for ; m%p == 0; m = m / p {
				c++
			}
			if c == 0 {
				continue
			}
			dn *= c + 1
		}
		if dn > n {
			return t, nil
		}
	}
	return 0, errors.New("No solution exist")
}
