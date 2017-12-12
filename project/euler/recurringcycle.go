package euler

/*
循環少數關係

* [循環小数 (Wikipedia)](https://ja.wikipedia.org/wiki/%E5%BE%AA%E7%92%B0%E5%B0%8F%E6%95%B0)
* [循環小数(1): フェルマーの小定理](http://tsujimotter.hatenablog.com/entry/2014/04/08/212954)
* [循環小数の不思議](http://www004.upp.so-net.ne.jp/s_honma/mathbun/mathbun36.htm)
* [循環小数の話](http://www.highflyer2.com/math/junkan.html)
* [循環小数はおもしろい](http://www2.math.kindai.ac.jp/~chinen/junkan_f/junkan.html)
*/

// CalculateRecurringCycle 1/nの循環節を計算する
func CalculateRecurringCycle(n int) string {
	if n < 1 {
		return ""
	}

	qs := []int{} // 小數點以下第(i+1)位の數字
	rs := []int{} // 小數點以下第(i+1)位を計算した時の剩餘を10倍した數値

	for q, r := 0, 10; r > 0; q, r = r/n, (r%n)*10 {
		for i := 0; i < len(qs); i++ {
			if qs[i] != q || rs[i] != r {
				continue
			}
			rc := make([]rune, len(qs[i:]))
			for k, v := range qs[i:] {
				rc[k] = rune(v) + '0'
			}
			return string(rc)
		}
		qs = append(qs, q)
		rs = append(rs, r)
	}
	return ""
}

// CalculateRecurringCycleLength 1/nの循環節の長さを計算する
func CalculateRecurringCycleLength(n int) int {
	if n < 1 {
		return 0
	}

	qs := []int{} // 小數點以下第(i+1)位の數字
	rs := []int{} // 小數點以下第(i+1)位を計算した時の剩餘を10倍した數値

	for q, r := 0, 10; r > 0; q, r = r/n, (r%n)*10 {
		for i := 0; i < len(qs); i++ {
			if qs[i] != q || rs[i] != r {
				continue
			}
			return len(qs[i:])
		}
		qs = append(qs, q)
		rs = append(rs, r)
	}
	return 0
}
