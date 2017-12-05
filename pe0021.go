package project_euler

// 約数の総和
//
// 36の約数(自分自身を除く) 1 + 2 + 3 + 4 + 6 + 9 + 12 + 18 = 55
//
// 36を素因数分解:
// 36 = 2^2 * 3^2
// (2^0 + 2^1 + 2^2) * (3^0 + 3^1 + 3^2) - 36 = 7 * 13 - 36 = 91 - 36 = 55

// PE0021a n以下の友愛数の總和
func PE0021a(n int) int {
	sumMap := map[int]int{} // m, mの約數の總和-m
	for i := 2; i <= n; i++ {
		m := SumFactors(i) - i
		if m == 1 {
			continue
		}
		sumMap[i] = m
	}
	sum := 0
	for k, v := range sumMap {
		if k == v {
			continue
		} else if _, ok := sumMap[v]; ok && sumMap[v] == k {
			sum += k
		}
	}
	return sum
}

// PE0021Memoization n以下の友愛数の總和(2)
func PE0021Memoization(n int) int {
	pm := []int{} // n以下の素數
	pg := NewPrimeGenerator()
	for p := pg.Next(); p <= n; p = pg.Next() {
		pm = append(pm, p)
	}

	sumMap := map[int]int{} // m, mの約數の總和-m
	for i := 2; i <= n; i++ {
		m := i
		pf := map[int]int{} // mの素因數
		for j := 0; pm[j] <= i/2; j++ {
			p := pm[j]
			if m%p == 0 {
				for ; m%p == 0; m = m / p {
					pf[p]++
				}
			}
		}
		// 素數は友愛數にならないので省く
		if len(pf) > 0 {
			sumMap[i] = 1
			for k, v := range pf {
				sumMap[i] *= (Pow(k, (v+1)) - 1) / (k - 1)
			}
			sumMap[i] -= i
		}
	}
	sum := 0
	for k, v := range sumMap {
		if k == v {
			continue
		} else if _, ok := sumMap[v]; ok && sumMap[v] == k {
			sum += k
		}
	}
	return sum
}
