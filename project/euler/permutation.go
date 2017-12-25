package euler

// Permutation 文字列の順列を生成
func Permutation(s string, k int) []string {
	ss := []string{} // 結果を格納

	var perm func([]rune, int, int)

	perm = func(rs []rune, i, k int) {
		if k == 0 {
			ss = append(ss, string(rs[:i]))
		} else {
			rt := make([]rune, len(rs))
			copy(rt, rs)
			perm(rt, i+1, k-1)
			for j := i + 1; j < len(rs); j++ {
				rt[i], rt[j] = rt[j], rt[i]
				perm(rt, i+1, k-1)
			}
		}
	}

	perm([]rune(s), 0, k)

	return ss
}
