package euler

// CollatzLength Collatz Sequenceの長さを計算
func CollatzLength(n int) int {
	if n < 2 {
		return 0
	}

	l := 0 // Collatz Sequenceの長さ
	for ; n > 1; l++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	return l + 1
}
