package euler

// CalculateDigitNumber 桁数を計算する
func CalculateDigitNumber(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n *= -1
	}
	m := 0
	for i := 1; n/i > 0; i *= 10 {
		m++
	}
	return m
}

// CirculateDigit 末尾の桁の數字を先頭に移動する
func CirculateDigit(n int) int {
	minus := false
	if n < 0 {
		minus = true
		n *= -1
	}

	m := 1
	for n/(10*m) > 0 {
		m *= 10
	}
	h := n / 10
	t := n % 10

	l := t*m + h
	if minus {
		return -1 * l
	}
	return l
}
