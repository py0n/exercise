package euler

// PE0026 d<nの内、1/dの循環節が最も長いものを計算する
func PE0026(n int) int {
	result, maxLength := 0, 0
	for i := 1; i < n; i++ {
		if l := CalculateRecurringCycleLength(i); l > maxLength {
			maxLength = l
			result = i
		}
	}
	return result
}
