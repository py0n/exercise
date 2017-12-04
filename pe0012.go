package project_euler

// http://hanatsubaki.shiseidogroup.jp/comic2/2368/
//
// n が a^x * b^y * c^z と素因数分解できる時、
// 約数の個数は(x+1)(y+1)(z+1)個である。

// PE0012 約數の個數が初めてnを越える三角數を計算
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

// PE0012a 約數の個數が初めてnを越える三角數を計算(2)
func PE0012a(n int) int {
	pm := []int{}             // 素數リスト
	pg := NewPrimeGenerator() // 素數ジェネレータ
	for i := 1; ; i++ {
		t := i * (i + 1) / 2 // 三角数
		for {
			pm = append(pm, pg.Next())
			if pm[len(pm)-1] > t {
				break
			}
		}

		// 素因数分解
		pf := map[int]int{} // 素因数
		for m, j := t, 0; m >= pm[j]; j++ {
			p := pm[j]
			for ; m%p == 0; m = m / p {
				pf[p]++
			}
		}

		d := 1 // 約数の個数
		for _, v := range pf {
			d *= (v + 1)
		}
		if d > n {
			return t
		}
	}
}
