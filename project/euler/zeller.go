package euler

// DayOfWeek 曜日を表す型
type DayOfWeek int

// 曜日
const (
	None      DayOfWeek = iota // 該當無し
	Monday                     // 月曜日
	Tuesday                    // 火曜日
	Wednesday                  // 水曜日
	Thursday                   // 木曜日
	Friday                     // 金曜日
	Saturday                   // 土曜日
	Sunday                     // 日曜日
)

// ZellerWeekday ツェラーの公式
func ZellerWeekday(y, m, d int) DayOfWeek {
	if y < 1582 || (y == 1582 && m < 10) || (y == 1582 && m == 10 && d < 15) {
		return None
	}
	if m < 3 {
		y, m = y-1, m+12
	}
	C := int(y / 100)
	G := 5*C + int(C/4)
	Y := y % 100
	return DayOfWeek(((d + int(26*(m+1)/10) + Y + int(Y/4) + G + 5) % 7) + 1)
}
