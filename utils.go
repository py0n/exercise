package project_euler

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
type PrimeGenerator struct {
	Next      func() int
	Size      func() int
	multiples map[int]int
	num       int
	primeChan chan int
}

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

func isPalindrome(s string) bool {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
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

// Lcm 最小公倍數を計算する
func Lcm(n, m int) int {
	g := Gcm(n, m)
	if g == 1 {
		return n * m
	}
	return n * m / g
}

func pow10(n int) int {
	m := 1
	for i := 0; i < n; i++ {
		m *= 10
	}
	return m
}

func multiplyDigits(n int) int {
	product := 1
	for n > 0 {
		product *= n % 10
		n = n / 10
	}
	return product
}

// zellerWeekday ツェラーの公式
func zellerWeekday(y, m, d int) int {
	y_, m_ := y, m
	if m < 3 {
		y_, m_ = y-1, m+12
	}
	C := int(y_ / 100)
	G := 5*C + int(C/4)
	Y := y_ % 100
	return ((d + int(26*(m_+1)/10) + Y + int(Y/4) + G + 5) % 7) + 1
}
