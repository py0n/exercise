package project_euler

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
