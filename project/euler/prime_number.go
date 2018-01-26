package euler

/*
素數關係

* [Golangで「エラトステネスの篩」で「1.3秒で百万個」の素数を計算できる「無限シーケンス」を作ってみた](https://qiita.com/antimon2/items/cada59fb3b1f028f4fb3)
* [Golangで「エラトステネスの篩」で「2.1秒で百万個」の素数を計算できる「無限シーケンス」を作ってみた](https://qiita.com/cia_rana/items/2a878181da41033ec1d8)
*/

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
