package euler

// Permutation 文字列の順列を生成
func Permutation(s string, k int) []string {
	ss := []string{} // 結果を格納

	var perm func([]rune, int)

	perm = func(rs []rune, i int) {
		if i == k {
			ss = append(ss, string(rs[:i]))
		} else {
			rt := make([]rune, len(rs))
			copy(rt, rs)
			perm(rt, i+1)
			for j := i + 1; j < len(rs); j++ {
				rt[i], rt[j] = rt[j], rt[i]
				perm(rt, i+1)
			}
		}
	}

	perm([]rune(s), 0)

	return ss
}
