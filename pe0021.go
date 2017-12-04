package project_euler

import "log"

// 約数の総和
//
// 36の約数(自分自身を除く) 1 + 2 + 3 + 4 + 6 + 9 + 12 + 18 = 55
//
// 36を素因数分解:
// 36 = 2^2 * 3^2
// (2^0 + 2^1 + 2^2) * (3^0 + 3^1 + 3^2) - 36 = 7 * 13 - 36 = 91 - 36 = 55

// PE0021 10000以下の友愛数を計算する
func PE0021(n int) []int {
	dsm := map[int]int{} // 数とその約数の総和
	for i := 0; i < n; i++ {
		m := i + 2
		pg := NewPrimeGenerator()
		pm := map[int]int{}
		for p := pg.Next(); p <= m/2; p = pg.Next() {
			if m%p == 0 {
				pm[p]++
			}
		}
		log.Println(m, pm)
		dsm[m] = 1
		for k, v := range pm {
			dsm[m] *= (pow(k, (v+1)) - 1) / (k - 1)
		}
		dsm[m] -= m
	}
	log.Println(dsm[220], dsm[284])
	ans := make([]int, 0, n)
	return ans
}
