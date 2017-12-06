package euler

// http://hanatsubaki.shiseidogroup.jp/comic2/2368/
//
// n が a^x * b^y * c^z と素因数分解できる時、
// 約数の個数は(x+1)(y+1)(z+1)個である。

// PE0012a 約數の個數が初めてnを越える三角數を計算
func PE0012a(n int) int {
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
	pg := NewPrimeGenerator()
	pf := map[int]int{}
	for m, p := n, pg.Next(); m >= p; p = pg.Next() {
		for ; m%p == 0; m = m / p {
			pf[p]++
		}
	}
	return pf
}

// PE0012b 約數の個數が初めてnを越える三角數を計算(2)
func PE0012b(n int) int {
	pg := NewPrimeGenerator() // 素數ジェネレータ
	pm := []int{pg.Next()}    // 素數リスト
	for i := 1; ; i++ {
		t := i * (i + 1) / 2 // 三角数
		pf := map[int]int{}  // 素因数
		for m, j := t, 0; m >= pm[j]; j++ {
			for ; m%pm[j] == 0; m = m / pm[j] {
				pf[pm[j]]++
			}
			if j == len(pm)-1 {
				pm = append(pm, pg.Next())
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

/*
=== RUN   TestPE0012a
--- PASS: TestPE0012a (3.13s)
=== RUN   TestPE0012b
--- PASS: TestPE0012b (0.06s)
BenchmarkPE0012a-2            1        3138566396 ns/op        308812312 B/op    784230 allocs/op
BenchmarkPE0012b-2           20          60763821 ns/op         3117858 B/op      49740 allocs/op
PASS
ok      github.com/py0n/project/euler   7.615s
*/
