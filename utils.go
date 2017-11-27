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

type PrimeGenerator2 struct {
	Next      func() int
	Size      func() int
	multiples map[int]int
	num       int
	primeChan chan int
}

func NewPrimeGenerator2() *PrimeGenerator2 {
	p := PrimeGenerator2{
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

func CreatePrimeArray(n int) []int {
	primeNumbers := []int{}
	primeGenerator := NewPrimeGenerator0()
	for {
		m := primeGenerator.Next()
		if m > n {
			break
		}
		primeNumbers = append(primeNumbers, m)
	}
	return primeNumbers
}
