package euler

// PE0027a 素数生成式
//
// |a| < 1000, |b| < 1000 の時, n^2+an+b で n=0 から初めて連続する整数で素数を生成したときに,
// 最も多く素数を生成できるa, bを計算する.
//
// a, b の条件を考える
//     1. n=0 の時, 0^2+a*0+b=b なので, b は素数
func PE0027a(aMax, bMax int) int {
	// 1000未満の素数を生成
	pg := NewPrimeGenerator()
	pm := map[int]bool{} // 素数マップ
	bs := []int{}        // bの取りうる値(bMax未満の素数)
	for p := pg.Next(); p < bMax; p = pg.Next() {
		pm[p] = true
		bs = append(bs, p)
	}

	aResult, bResult, nMax := 0, -1*aMax, -1*bMax

	// a, bを動かす
	for _, b := range bs {
		for a := -1 * (aMax - 1); a <= aMax-1; a++ {
			for n := 0; ; n++ {
				m := n*n + a*n + b
				for p := pg.Next(); p <= m; p = pg.Next() {
					pm[p] = true
				}
				// n^2+a*n+b が素数で在る間は継続
				if _, ok := pm[m]; ok {
					continue
				}
				if n > nMax {
					aResult, bResult, nMax = a, b, n
				}
				break
			}
		}
	}

	return aResult * bResult
}
