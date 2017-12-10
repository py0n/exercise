package euler

// https://projecteuler.net/problem=23

// Pe0023Limit 此より大きい整數は二つの過剰數の和として表せる
var Pe0023Limit = 28123

// PE0023 28123以下の數で二つの過剰數の和で書き表せない正の整數の總和を計算する
func PE0023() int {
	abundants := []int{} // 28123以下の過剰數

	pg := NewPrimeGenerator() // 素數ジェネレータ
	pm := []int{pg.Next()}    // 素數配列

	for i := 12; i <= Pe0023Limit; i++ {
		m, pf := i, map[int]int{}
		for j := 0; m >= pm[j]; j++ {
			if j == len(pm)-1 {
				pm = append(pm, pg.Next())
			}
			for p := pm[j]; m%p == 0; m = m / p {
				pf[p]++
			}
		}
		// 約數の和
		n := 1
		for k, v := range pf {
			n *= (Pow(k, (v+1)) - 1) / (k - 1)
		}
		// 過剰數 (眞の約數の和が元の數より大きい)
		if n > 2*i {
			abundants = append(abundants, i)
		}
	}

	sumList := map[int]int{}
	for i := 0; i < len(abundants); i++ {
		for j := 0; j < len(abundants) && abundants[i]+abundants[j] <= Pe0023Limit; j++ {
			sumList[abundants[i]+abundants[j]]++
		}
	}

	sum := 0
	for i := 1; i < Pe0023Limit; i++ {
		if _, ok := sumList[i]; !ok {
			sum += i
		}
	}

	return sum
}
