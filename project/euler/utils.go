package euler

import (
	"math/big"
)

type PrimeGenerator0 struct {
	primeChan chan int
	multiples map[int]int
}

func NewPrimeGenerator0() *PrimeGenerator0 {
	primeGenerator := PrimeGenerator0{
		primeChan: make(chan int),
		multiples: map[int]int{}, // Key: 元の素数の奇倍数, Value: 元の素数の2倍の数
	}

	go primeGenerator.start()

	return &primeGenerator
}

func (p *PrimeGenerator0) start() {
	// 2は偶素数なので例外扱い
	p.primeChan <- 2

	// 3以上の奇数の中で素数であるものを探す
	for num := int(3); ; num += 2 {
		factor, hasFactor := p.multiples[num]

		// hasFactorがtrueならば、numはfactor/2の倍数なので素数ではない
		if hasFactor {
			delete(p.multiples, num)
		} else {
			// 後に追加する倍数を奇数に限定するための調節
			factor = num * 2
		}

		// 新たな倍数を追加
		for newNum := num + factor; ; newNum += factor {
			_, hasNewFactor := p.multiples[newNum]

			// hasNewFactorがtrueならば、newNumは既に登録された倍数なのでスルー
			if !hasNewFactor {
				p.multiples[newNum] = factor
				break
			}
		}

		if !hasFactor {
			p.primeChan <- num
		}
	}
}

func (p *PrimeGenerator0) Next() int {
	return <-p.primeChan
}

func (p *PrimeGenerator0) Size() int {
	return len(p.multiples)
}

type PrimeGenerator1 struct {
	primeChan chan int
	multiples map[int]int
}

func NewPrimeGenerator1() *PrimeGenerator1 {
	primeGenerator := PrimeGenerator1{
		primeChan: make(chan int),
		multiples: map[int]int{}, // Key: 元の素数の奇倍数, Value: 元の素数の2倍の数
	}

	go primeGenerator.start()

	return &primeGenerator
}
func (p *PrimeGenerator1) start() {
	// 2は偶素数なので例外扱い
	p.primeChan <- 2

	// 3以上の奇数の中で素数であるものを探す
	for num := int(3); ; num += 2 {
		factor, hasFactor := p.multiples[num]
		if !hasFactor {
			// 後に追加する倍数を奇数に限定するための調節
			factor = num * 2
		}

		// 新たな倍数を追加
		for newNum := num + factor; ; newNum += factor {
			_, hasNewFactor := p.multiples[newNum]

			// hasNewFactorがtrueならば、newNumは既に登録された倍数なのでスルー
			if !hasNewFactor {
				p.multiples[newNum] = factor
				break
			}
		}

		if !hasFactor {
			p.primeChan <- num
		}
	}
}

func (p *PrimeGenerator1) Next() int {
	return <-p.primeChan
}

func (p *PrimeGenerator1) Size() int {
	return len(p.multiples)
}

// https://qiita.com/cia_rana/items/2a878181da41033ec1d8
// https://qiita.com/antimon2/items/cada59fb3b1f028f4fb3

// PrimeGenerator 素數生成器
type PrimeGenerator struct {
	Next      func() int
	Size      func() int
	multiples map[int]int
	num       int
	primeChan chan int
}

// NewPrimeGenerator 素數生成器を作成する
func NewPrimeGenerator() *PrimeGenerator {
	p := PrimeGenerator{
		primeChan: make(chan int),
		multiples: map[int]int{},
	}

	p.Next = func() int {
		if p.num == 0 {
			p.num = 1
			return 2
		}

		for {
			p.num += 2

			factor, hasFactor := p.multiples[p.num]
			if hasFactor {
				delete(p.multiples, p.num)
			} else {
				factor = p.num * 2
			}

			for newNum := p.num + factor; ; newNum += factor {
				_, hasNewFactor := p.multiples[newNum]
				if !hasNewFactor {
					p.multiples[newNum] = factor
					break
				}
			}

			if !hasFactor {
				return p.num
			}
		}
	}

	p.Size = func() int {
		return len(p.multiples)
	}

	return &p
}

// BigNumberLength 桁数を計算
func BigNumberLength(n *big.Int) int {
	m := big.NewInt(0).Set(n)
	m.Abs(m)
	return len(m.String())
}

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

// Fibonacci n番目のFibonacci數を計算
func Fibonacci(nth int, fibonacciMap map[int]*big.Int) *big.Int {
	if nth == 0 || nth == 1 {
		return big.NewInt(int64(1))
	}
	if fibonacciMap == nil {
		fibonacciMap = make(map[int]*big.Int)
	}
	if _, ok := fibonacciMap[nth]; ok {
		return fibonacciMap[nth]
	}
	fibonacciMap[nth] = big.NewInt(int64(0)).Add(
		Fibonacci(nth-1, fibonacciMap),
		Fibonacci(nth-2, fibonacciMap),
	)
	return fibonacciMap[nth]
}

// Gcm 最大公約數を計算する
func Gcm(n, m int) int {
	if n < 1 || m < 1 {
		return 0
	}
	if n < m {
		n, m = m, n
	}
	for {
		r := n % m
		if r > 0 {
			n, m = m, r
		}
		if r == 0 {
			return m
		}
	}
	return 0
}

// IsPalindrome 囘文？
func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

// Lcm 最小公倍數を計算する
func Lcm(n, m int) int {
	g := Gcm(n, m)
	if g == 0 {
		return 0
	} else if g == 1 {
		return n * m
	}
	return n * m / g
}

// PrimeFactorize 素因數分解する
func PrimeFactorize(n int) map[int]int {
	if n < 2 {
		return nil
	}
	pm := map[int]int{} // 素因數
	pg := NewPrimeGenerator()
	for m, p := n, pg.Next(); p <= n/2 || m > 1; p = pg.Next() {
		if m%p > 0 {
			continue
		}
		for ; m%p == 0; m = m / p {
			pm[p]++
		}
	}
	if len(pm) == 0 {
		return map[int]int{n: 1}
	}
	return pm
}

// Pow 累乘を計算
func Pow(m, n int) int {
	if n < 0 {
		return 0
	}
	p := 1
	for i := 0; i < n; i++ {
		p *= m
	}
	return p
}

// Pow10 10の累乘を計算
func Pow10(n int) int {
	return Pow(10, n)
}

// SumFactors 約數の総和を計算
func SumFactors(n int) int {
	if n < 2 {
		return 0
	}
	pm := PrimeFactorize(n)
	sum := 1
	for k, v := range pm {
		sum *= (Pow(k, (v+1)) - 1) / (k - 1)
	}
	return sum
}

// SumSlice スライスの要素の總和を計算
func SumSlice(slice []int, convertor func(n int) int) int {
	sum := 0
	for i := 0; i < len(slice); i++ {
		sum += convertor(slice[i])
	}
	return sum
}

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
	y_, m_ := y, m
	if m < 3 {
		y_, m_ = y-1, m+12
	}
	C := int(y_ / 100)
	G := 5*C + int(C/4)
	Y := y_ % 100
	return DayOfWeek(((d + int(26*(m_+1)/10) + Y + int(Y/4) + G + 5) % 7) + 1)
}

func multiplyDigits(n int) int {
	product := 1
	for n > 0 {
		product *= n % 10
		n = n / 10
	}
	return product
}
