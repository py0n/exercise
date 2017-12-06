package euler

// PE0019 20世紀の月初が日曜日になる回数を計算
func PE0019(s, e int) int {
	sum := 0
	for y := s; y < e+1; y++ {
		for m := 1; m <= 12; m++ {
			if ZellerWeekday(y, m, 1) != Sunday {
				continue
			}
			sum++
		}
	}
	return sum
}
