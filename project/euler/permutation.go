package euler

// Permutation 文字列の順列を生成
func Permutation(s string, k int) []string {
	rs := []rune(s)

	// 引數檢査
	if k < 1 || len(rs) < k {
		return nil
	}

	// 結果を格納する配列
	ss := make([]string, 0, PermutationNumber(len(rs), k))

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

	perm(rs, 0)

	return ss
}

// PermutationA 文字列の順列を生成
func PermutationA(s string, k int) []string {
	rs := []rune(s)

	// 引數檢査
	if k < 1 || len(rs) < k {
		return nil
	}

	// 結果を格納する配列
	ss := make([]string, 0, PermutationNumber(len(rs), k))

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

	perm(rs, 0)

	return ss
}

// PermutationNumber 順列の數P(m,n)
func PermutationNumber(m, n int) int {
	if m < 1 || n < 1 || m < n {
		return 0
	}

	p := 1
	for i := m - n + 1; i <= m; i++ {
		p *= i
	}
	return p
}
