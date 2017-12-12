package euler

import (
	"encoding/csv"
	"io"
	"math"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// PE0001 1000以下の3と5の倍數の總和を計算
//
// https://projecteuler.net/problem=1
func PE0001(n int) int {
	result := 0
	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			result = result + i
		}
	}
	return result
}

// PE0002 4,000,000を越えないFibonacci數の内、偶數である物の總和を計算
//
// https://projecteuler.net/problem=2
func PE0002(n int) (int, error) {
	if n < 3 {
		return 0, errors.New("n must be >= 3")
	}
	i, j, result := 1, 1, 0
	for ; j < n; i, j = j, i+j {
		if j%2 == 0 {
			result = result + j
		}
	}
	return result, nil
}

// PE0003SortSort 600851475143の最大の素因數を計算 (sort.Sort版)
//
// https://projecteuler.net/problem=3
func PE0003SortSort(n int) (int, error) {
	primeNumbers := []int{}

	uLimit := int(math.Sqrt(float64(n)))

	primeGenerator := NewPrimeGenerator()
	for {
		m := primeGenerator.Next()
		if m > uLimit {
			break
		}
		primeNumbers = append(primeNumbers, m)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(primeNumbers)))

	for _, p := range primeNumbers {
		if n%p == 0 {
			return p, nil
		}
	}

	return 0, errors.New("no solution")
}

// PE0003SortSlice 600851475143の最大の素因數を計算 (sort.SliceStable版)
func PE0003SortSlice(n int) (int, error) {
	primeNumbers := []int{}

	uLimit := int(math.Sqrt(float64(n)))

	primeGenerator := NewPrimeGenerator()
	for {
		m := primeGenerator.Next()
		if m > uLimit {
			break
		}
		primeNumbers = append(primeNumbers, m)
	}

	sort.SliceStable(primeNumbers, func(i, j int) bool { return j < i })

	for _, p := range primeNumbers {
		if n%p == 0 {
			return p, nil
		}
	}

	return 0, errors.New("no solution")
}

// PE0004 二つの三桁の数値の積で囘文と成る物の内、最大の物を計算する
//
// https://projecteuler.net/problem=4
func PE0004(n int) int {

	imax := Pow10(n) - Pow10(n-1) - 1

	max := Pow10(n) - 1

	mMax, x0, y0 := 0, 0, 0
	for i := 0; i <= imax; i++ {
		for j := 0; j <= i; j++ {
			x := max - i
			y := max - j
			if x < x0 && x < y0 && y < x0 && y < y0 {
				continue
			}
			m := x * y
			if m <= mMax {
				continue
			}
			if IsPalindrome(strconv.Itoa(m)) {
				mMax, x0, y0 = m, x, y
			}
		}
	}
	return mMax
}

// PE0005 n迄の數の最小公倍數を計算する
//
// https://projecteuler.net/problem=5
func PE0005(n int) int {
	if n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result = Lcm(result, i)
	}
	return result
}

// PE0006 100迄の和の二乘と、二乘の和の差を計算する。
//
// https://projecteuler.net/problem=6
func PE0006(n int) int {
	sum0, sum1 := 0, 0
	for i := 1; i <= n; i++ {
		sum0 = sum0 + i*i
		sum1 = sum1 + i
	}
	return sum1*sum1 - sum0
}

// PE0007 10001番目の素數を計算する。
//
// https://projecteuler.net/problem=7
func PE0007(n int) int {
	primeGenerator := NewPrimeGenerator()
	for i := 0; i < n-1; i++ {
		primeGenerator.Next()
	}
	return primeGenerator.Next()
}

// PE0008 13の隣接する數字の積の内、最大の物を計算する。
//
// https://projecteuler.net/problem=8
func PE0008(n int, grid string) (int, error) {
	maxNum := 0
	for i := 0; i < len(grid)-n-1; i++ {
		s := grid[i : i+n]
		m, err := strconv.Atoi(s)
		if err != nil {
			return 0, errors.Wrap(err, s+" is invalid digits")
		}
		product := multiplyDigits(m)
		if product > maxNum {
			maxNum = product
		}
	}
	return maxNum, nil
}

// PE0009 a+b+c=1000を滿たすピタゴラス數の積を計算する。
//
// a < b < c and a + b + c = 1000 の時の abc の値
// a: a < 1000/3
// c: c = 1000 - a - b
// b: b < c = 1000 - a - b → 2b < 1000 - a → b < (1000 - a)/2
//    a < b < (1000 - a) / 2
//
// https://projecteuler.net/problem=9
func PE0009() int {
	maxNum := 0
	for a := 1; a <= 1000/3; a++ {
		for b := a + 1; b <= (1000-a)/2; b++ {
			c := 1000 - a - b
			if c <= b {
				continue
			}
			if a*a+b*b != c*c {
				continue
			}
			m := a * b * c
			if m > maxNum {
				maxNum = m
			}
		}
	}
	return maxNum
}

// PE0010 2,000,000以下の素數の和を計算する。
//
// https://projecteuler.net/problem=10
func PE0010(n int) int {
	sum := 0

	primeGenerator := NewPrimeGenerator()
	for p := 0; p < n; p = primeGenerator.Next() {
		sum += p
	}

	return sum
}

const pe0011GridSize = 20
const pe0011Length = 4

var pe0011DigitGrid = []string{
	"08", "02", "22", "97", "38", "15", "00", "40", "00", "75", "04", "05", "07", "78", "52", "12", "50", "77", "91", "08",
	"49", "49", "99", "40", "17", "81", "18", "57", "60", "87", "17", "40", "98", "43", "69", "48", "04", "56", "62", "00",
	"81", "49", "31", "73", "55", "79", "14", "29", "93", "71", "40", "67", "53", "88", "30", "03", "49", "13", "36", "65",
	"52", "70", "95", "23", "04", "60", "11", "42", "69", "24", "68", "56", "01", "32", "56", "71", "37", "02", "36", "91",
	"22", "31", "16", "71", "51", "67", "63", "89", "41", "92", "36", "54", "22", "40", "40", "28", "66", "33", "13", "80",
	"24", "47", "32", "60", "99", "03", "45", "02", "44", "75", "33", "53", "78", "36", "84", "20", "35", "17", "12", "50",
	"32", "98", "81", "28", "64", "23", "67", "10", "26", "38", "40", "67", "59", "54", "70", "66", "18", "38", "64", "70",
	"67", "26", "20", "68", "02", "62", "12", "20", "95", "63", "94", "39", "63", "08", "40", "91", "66", "49", "94", "21",
	"24", "55", "58", "05", "66", "73", "99", "26", "97", "17", "78", "78", "96", "83", "14", "88", "34", "89", "63", "72",
	"21", "36", "23", "09", "75", "00", "76", "44", "20", "45", "35", "14", "00", "61", "33", "97", "34", "31", "33", "95",
	"78", "17", "53", "28", "22", "75", "31", "67", "15", "94", "03", "80", "04", "62", "16", "14", "09", "53", "56", "92",
	"16", "39", "05", "42", "96", "35", "31", "47", "55", "58", "88", "24", "00", "17", "54", "24", "36", "29", "85", "57",
	"86", "56", "00", "48", "35", "71", "89", "07", "05", "44", "44", "37", "44", "60", "21", "58", "51", "54", "17", "58",
	"19", "80", "81", "68", "05", "94", "47", "69", "28", "73", "92", "13", "86", "52", "17", "77", "04", "89", "55", "40",
	"04", "52", "08", "83", "97", "35", "99", "16", "07", "97", "57", "32", "16", "26", "26", "79", "33", "27", "98", "66",
	"88", "36", "68", "87", "57", "62", "20", "72", "03", "46", "33", "67", "46", "55", "12", "32", "63", "93", "53", "69",
	"04", "42", "16", "73", "38", "25", "39", "11", "24", "94", "72", "18", "08", "46", "29", "32", "40", "62", "76", "36",
	"20", "69", "36", "41", "72", "30", "23", "88", "34", "62", "99", "69", "82", "67", "59", "85", "74", "04", "36", "16",
	"20", "73", "35", "29", "78", "31", "90", "01", "74", "31", "49", "71", "48", "86", "81", "16", "23", "57", "05", "54",
	"01", "70", "54", "71", "83", "51", "54", "69", "16", "92", "33", "48", "61", "43", "52", "01", "89", "19", "67", "48",
}

// PE0011 四つの隣接する數字の積の内、最大の物を計算する。
//
// +---+-----+---+
// |   |     |   |
// | A |  B  | C |
// |   |     |   |
// +---+-----+---+
// | D |  E  | F |
// +---+-----+---+
//
// A: R, D, RD
// B: R, D, LD, RD
// C: D, LD
// D: R
// E: R
// F: なし
//
// R: A B D E
// D: A B C
// LD: B C
// RD: A
//
// https://projecteuler.net/problem=11
func PE0011() int {

	digitGrid := make([]int, len(pe0011DigitGrid))
	for i := 0; i < len(pe0011DigitGrid); i++ {
		m, _ := strconv.Atoi(pe0011DigitGrid[i])
		digitGrid[i] = m
	}

	maxNumber := 0
	for i := 0; i < len(pe0011DigitGrid); i++ {
		if m := pe0011MultiplyTypeR(i, digitGrid); m > maxNumber {
			maxNumber = m
		}
		if m := pe0011MultiplyTypeD(i, digitGrid); m > maxNumber {
			maxNumber = m
		}
		if m := pe0011MultiplyTypeLD(i, digitGrid); m > maxNumber {
			maxNumber = m
		}
		if m := pe0011MultiplyTypeRD(i, digitGrid); m > maxNumber {
			maxNumber = m
		}
	}
	return maxNumber
}

func pe0011MultiplyTypeR(n int, digitGrid []int) int {
	x := n % pe0011GridSize

	if pe0011GridSize-pe0011Length < x {
		return 0
	}

	product := 1
	for i := 0; i < pe0011Length; i++ {
		product *= digitGrid[n+i]
	}
	return product
}

func pe0011MultiplyTypeD(n int, digitGrid []int) int {
	y := n / pe0011GridSize

	if pe0011GridSize-pe0011Length < y {
		return 0
	}

	product := 1
	for i := 0; i < pe0011Length; i++ {
		product *= digitGrid[n+i*pe0011GridSize]
	}
	return product
}

func pe0011MultiplyTypeLD(n int, digitGrid []int) int {
	x := n % pe0011GridSize
	y := n / pe0011GridSize

	if x < pe0011Length || pe0011GridSize-pe0011Length < y {
		return 0
	}

	product := 1
	for i := 0; i < pe0011Length; i++ {
		product *= digitGrid[n+i*(pe0011GridSize-1)]
	}
	return product
}

func pe0011MultiplyTypeRD(n int, digitGrid []int) int {
	x := n % pe0011GridSize
	y := n / pe0011GridSize

	if pe0011Length <= x || pe0011GridSize-pe0011Length < y {
		return 0
	}

	product := 1
	for i := 0; i < pe0011Length; i++ {
		product *= digitGrid[n+i*(pe0011GridSize+1)]
	}
	return product
}

// PE0012a 約數の個數が初めてnを越える三角數を計算
//
// n が a^x * b^y * c^z と素因数分解できる時、
// 約数の個数は(x+1)(y+1)(z+1)個である。
//
// http://hanatsubaki.shiseidogroup.jp/comic2/2368/
func PE0012a(n int) int {
	for i := 1; ; i++ {
		triangleNumber := i * (i + 1) / 2 // 三角数
		primeMap := factorizePrime(triangleNumber)
		solution := 1
		for _, v := range primeMap {
			solution *= (v + 1)
		}
		if solution > n {
			return triangleNumber
		}
	}
}

func factorizePrime(n int) map[int]int {
	pg := NewPrimeGenerator()
	pf := map[int]int{}
	for m, p := n, pg.Next(); m >= p; p = pg.Next() {
		for ; m%p == 0; m = m / p {
			pf[p]++
		}
	}
	return pf
}

// PE0012b 約數の個數が初めてnを越える三角數を計算(2)
func PE0012b(n int) int {
	pg := NewPrimeGenerator() // 素數ジェネレータ
	pm := []int{pg.Next()}    // 素數リスト
	for i := 1; ; i++ {
		t := i * (i + 1) / 2 // 三角数
		pf := map[int]int{}  // 素因数
		for m, j := t, 0; m >= pm[j]; j++ {
			for ; m%pm[j] == 0; m = m / pm[j] {
				pf[pm[j]]++
			}
			if j == len(pm)-1 {
				pm = append(pm, pg.Next())
			}
		}

		d := 1 // 約数の個数
		for _, v := range pf {
			d *= (v + 1)
		}
		if d > n {
			return t
		}
	}
}

/*
=== RUN   TestPE0012a
--- PASS: TestPE0012a (3.13s)
=== RUN   TestPE0012b
--- PASS: TestPE0012b (0.06s)
BenchmarkPE0012a-2            1        3138566396 ns/op        308812312 B/op    784230 allocs/op
BenchmarkPE0012b-2           20          60763821 ns/op         3117858 B/op      49740 allocs/op
PASS
ok      github.com/py0n/project/euler   7.615s
*/

// PE0013 与へられたn個の整數の和の最初の10桁を計算する。
func PE0013(digits []string) int {
	sum := big.NewInt(0)
	for _, s := range digits {
		n := new(big.Int)
		n.SetString(s, 10)
		sum.Add(sum, n)
	}
	s := sum.String()
	m, _ := strconv.Atoi(s[0:10])
	return m
}

/*
計算済の数字をマップで記録したら寧ろ遅くなった。
メモリ確保の為。
=== RUN   TestPE0014a
--- PASS: TestPE0014a (0.36s)
=== RUN   TestPE0014b
--- PASS: TestPE0014b (0.77s)
BenchmarkPE0014a-2        10000           2284030 ns/op               0 B/op          0 allocs/op
BenchmarkPE0014b-2        10000           4302648 ns/op          684764 B/op        627 allocs/op
PASS
ok      github.com/py0n/project/euler   67.015s
*/

// PE0014a 1,000,000以下の数字でCollatz Sequenceが最も長いものを計算
//
// https://projecteuler.net/problem=14
func PE0014a(n int) int {
	maxLength := 0      // Collatz Sequenceの最大の長さ
	startingNumber := 0 // その時の開始数値
	for i := 1; i < n; i++ {
		if l := CollatzLength(i); l > maxLength {
			maxLength = l
			startingNumber = i
		}
	}
	return startingNumber
}

// PE0014b 1,000,000以下の数字でCollatz Sequenceが最も長いものを計算(2)
func PE0014b(n int) int {

	cm := map[int]int{} // その数のCollatz Length

	collatzLength := func(n int) int {
		if n < 2 {
			return 0
		}
		cm[n] = 0
		for m := n; m > 1; cm[n]++ {
			if m%2 == 0 {
				m = m / 2
			} else {
				m = 3*m + 1
			}
			if _, ok := cm[m]; ok {
				cm[n] += cm[m]
				break
			}
		}
		return cm[n]
	}

	maxLength := 0
	startingNumber := 0

	for i := 1; i < n; i++ {
		if l := collatzLength(i); l+1 > maxLength {
			maxLength = l + 1
			startingNumber = i
		}
	}
	return startingNumber
}

/*
BenchmarkPE0015-4                        300000              9816 ns/op            3440 B/op         86 allocs/op
BenchmarkPE0015Dp-4                      100000             12500 ns/op            4096 B/op          1 allocs/op
BenchmarkPE0015Memoization-4             300000              5701 ns/op            4096 B/op          1 allocs/op
PASS
ok      github.com/py0n/project/euler   6.166s
*/

// PE0015 組み合はせを使用して計算 (bigも使用)
//
// +-+-+-+
// | | | |
// +-+-+-+
// | | | |
// +-+-+-+
// | | | |
// +-+-+-+
//
// |  | 1| 2| 3| 4| 5|
// | 1| 2| 3| 4| 5| 6|
// | 2| 3| 6|10|15|
// | 3| 4|10|20|  |
// | 4| 5|15|  |  |
// | 5| 6|
//
// n * m -> C(n+m,n)
//
// https://projecteuler.net/problem=15
func PE0015(n int) int64 {
	numerator := big.NewInt(1) // 分子
	for i := 0; i < n; i++ {
		numerator.Mul(numerator, big.NewInt(int64(2*n-i)))
	}
	denominator := big.NewInt(1) //分母
	for i := 0; i < n; i++ {
		denominator.Mul(denominator, big.NewInt(int64(n-i)))
	}
	r := big.NewInt(0)
	r.Quo(numerator, denominator)
	return r.Int64()
}

// PE0015Dp DPを使用して計算
//
// + - 1 - 1 - 1
// |   |   |   |
// 1 - 2 - 3 - 4
// |   |   |   |
// 1 - 3 - 6 -10
// |   |   |   |
// 1 - 4 -10 -20
func PE0015Dp(n int) int64 {
	m := n + 1 // 一辺の點の數
	grid := make([]int64, m*m)
	grid[0] = 1 // 始點
	for i := 1; i < m*m; i++ {
		// 左の點が存在するなら、其處迄の經路數を加算
		if j := i % m; j > 0 {
			grid[i] += grid[i-1]
		}
		// 上の點が存在するなら、其處までの經路數を加算
		if j := i - m; j >= 0 {
			grid[i] += grid[i-m]
		}
	}
	return grid[m*m-1]
}

// PE0015Memoization メモ化を利用して計算 (再歸を試したが計算が終了しなかつた)
func PE0015Memoization(n int) int64 {
	memo := make([]int64, (n+1)*(n+1))

	var routeNumber func(int, int) int64
	routeNumber = func(i, j int) int64 {
		if i < 0 || j < 0 {
			return 0
		} else if i == 0 && j == 0 {
			return 1
		}

		p0 := j*(n+1) + i
		if memo[p0] > 0 {
			return memo[p0]
		}
		p1 := i*(n+1) + j // 對稱性から
		if memo[p1] > 0 {
			return memo[p1]
		}

		memo[p0] = routeNumber(i-1, j) + routeNumber(i, j-1)

		return memo[p0]
	}

	return routeNumber(n, n)
}

// PE0016 2^1000 の各桁の和を計算
//
// 2^1000 は大体300桁(log10 2 = 0.30....)
//
//https://projecteuler.net/problem=16
func PE0016(n int) int {
	// 2^n
	m := big.NewInt(2)
	for i := 0; i < n-1; i++ {
		m.Mul(m, big.NewInt(2))
	}

	s := m.String()

	sum := 0
	for i := 0; i < len(s); i++ {
		j, _ := strconv.Atoi(s[i : i+1])
		sum += j
	}
	return sum
}

var pe0017WordLength = map[int]int{
	1000: len(strings.Replace("one thousand", " ", "", -1)),
	900:  len(strings.Replace("nine hundred ", " ", "", -1)),
	800:  len(strings.Replace("eight hundred", " ", "", -1)),
	700:  len(strings.Replace("seven hundred", " ", "", -1)),
	600:  len(strings.Replace("six hundred", " ", "", -1)),
	500:  len(strings.Replace("five hundred", " ", "", -1)),
	400:  len(strings.Replace("four hundred", " ", "", -1)),
	300:  len(strings.Replace("three hundred", " ", "", -1)),
	200:  len(strings.Replace("two hundred", " ", "", -1)),
	100:  len(strings.Replace("one hundred", " ", "", -1)),
	90:   len("ninety"),
	80:   len("eighty"),
	70:   len("seventy"),
	60:   len("sixty"),
	50:   len("fifty"),
	40:   len("forty"),
	30:   len("thirty"),
	20:   len("twenty"),
	19:   len("nineteen"),
	18:   len("eighteen"),
	17:   len("seventeen"),
	16:   len("sixteen"),
	15:   len("fifteen"),
	14:   len("fourteen"),
	13:   len("thirteen"),
	12:   len("twelve"),
	11:   len("eleven"),
	10:   len("ten"),
	9:    len("nine"),
	8:    len("eight"),
	7:    len("seven"),
	6:    len("six"),
	5:    len("five"),
	4:    len("four"),
	3:    len("three"),
	2:    len("two"),
	1:    len("one"),
}

// PE0017SortSort 1-1000までの数字のスペルに含まれる文字数の総和を計算
//
// https://projecteuler.net/problem=16
func PE0017SortSort(n int) (int, error) {
	if n < 1 || 1000 < n {
		return 0, errors.New("n must be between 1 and 1000")
	}

	// keys (降順)
	ks := []int{}
	for k := range pe0017WordLength {
		ks = append(ks, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ks)))

	sumWc := 0 // 総文字数
	for m := 1; m <= n; m++ {
		l := m
		for _, k := range ks {
			if l/k > 0 {
				sumWc += pe0017WordLength[k]
				if l = l % k; l > 0 && 100 <= k && k <= 900 {
					sumWc += 3 // "and" 分
				}
			}
		}
	}
	return sumWc, nil
}

// PE0017SortSlice 1-1000までの数字のスペルに含まれる文字数の総和を計算(2)
func PE0017SortSlice(n int) (int, error) {
	if n < 1 || 1000 < n {
		return 0, errors.New("n must be between 1 and 1000")
	}

	// keys (降順)
	ks := []int{}
	for k := range pe0017WordLength {
		ks = append(ks, k)
	}
	sort.Slice(ks, func(i, j int) bool { return j < i })

	sumWc := 0 // 総文字数
	for m := 1; m <= n; m++ {
		l := m
		for _, k := range ks {
			if l/k > 0 {
				sumWc += pe0017WordLength[k]
				if l = l % k; l > 0 && 100 <= k && k <= 900 {
					sumWc += 3 // "and" 分
				}
			}
		}
	}
	return sumWc, nil
}

// PE0018Memoization Problem 18をメモ化を使用して計算
//
// https://projecteuler.net/problem=18
func PE0018Memoization(size int, data []string) int {
	memo := make([]int, size*(size+1)/2) // メモ

	var routeSum func(int, int) int
	routeSum = func(i, j int) int {
		if i < 0 || j < 0 || i > j {
			return 0
		}

		p := j*(j+1)/2 + i
		if memo[p] > 0 {
			return memo[p]
		}

		m, _ := strconv.Atoi(data[p])

		lv := routeSum(i-1, j-1)
		rv := routeSum(i, j-1)

		if lv > rv {
			memo[p] = m + lv
		} else {
			memo[p] = m + rv
		}
		return memo[p]
	}

	nMax := 0
	for i := 0; i < size; i++ {
		if routeSum(i, size-1) > nMax {
			nMax = routeSum(i, size-1)
		}
	}
	return nMax
}

// PE0019 20世紀の月初が日曜日になる回数を計算
//
// https://projecteuler.net/problem=19
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

// PE0020 階乗の各桁の和を計算
//
// https://projecteuler.net/problem=20
func PE0020(n int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		return 1
	}

	u := big.NewInt(int64(1))
	m := big.NewInt(int64(n))
	f := big.NewInt(int64(1))
	for m.Cmp(u) > 0 {
		f.Mul(f, m)
		m.Sub(m, u)
	}

	sum := 0
	s := f.String()
	for i := 0; i < len(s)-1; i++ {
		m, _ := strconv.Atoi(s[i : i+1])
		sum += m
	}
	return sum
}

// PE0021a n以下の友愛数の總和
//
// 約数の総和
//
// 36の約数(自分自身を除く) 1 + 2 + 3 + 4 + 6 + 9 + 12 + 18 = 55
//
// 36を素因数分解:
// 36 = 2^2 * 3^2
// (2^0 + 2^1 + 2^2) * (3^0 + 3^1 + 3^2) - 36 = 7 * 13 - 36 = 91 - 36 = 55
//
// https://projecteuler.net/problem=21
func PE0021a(n int) int {
	sumMap := map[int]int{} // m, mの約數の總和-m
	for i := 2; i <= n; i++ {
		m := SumFactors(i) - i
		if m == 1 {
			continue
		}
		sumMap[i] = m
	}
	sum := 0
	for k, v := range sumMap {
		if k == v {
			continue
		} else if _, ok := sumMap[v]; ok && sumMap[v] == k {
			sum += k
		}
	}
	return sum
}

// PE0021Memoization n以下の友愛数の總和(2)
func PE0021Memoization(n int) int {
	pm := []int{} // n以下の素數
	pg := NewPrimeGenerator()
	for p := pg.Next(); p <= n; p = pg.Next() {
		pm = append(pm, p)
	}

	sumMap := map[int]int{} // m, mの約數の總和-m
	for i := 2; i <= n; i++ {
		m := i
		pf := map[int]int{} // mの素因數
		for j := 0; pm[j] <= i/2; j++ {
			p := pm[j]
			if m%p == 0 {
				for ; m%p == 0; m = m / p {
					pf[p]++
				}
			}
		}
		// 素數は友愛數にならないので省く
		if len(pf) > 0 {
			sumMap[i] = 1
			for k, v := range pf {
				sumMap[i] *= (Pow(k, (v+1)) - 1) / (k - 1)
			}
			sumMap[i] -= i
		}
	}
	sum := 0
	for k, v := range sumMap {
		if k == v {
			continue
		} else if _, ok := sumMap[v]; ok && sumMap[v] == k {
			sum += k
		}
	}
	return sum
}

// PE0022 リストに在る単語の点数を計算し総和を求める
//
// https://projecteuler.net/problem=22
func PE0022(fn string) (int, error) {
	fh, err := os.Open(fn)
	if err != nil {
		return 0, err
	}
	defer fh.Close()

	reader := csv.NewReader(fh)
	rows, err := reader.ReadAll()
	if err != io.EOF && err != nil {
		return 0, err
	}

	names := []string{}
	for _, row := range rows {
		names = append(names, row...)
	}
	sort.SliceStable(names, func(i, j int) bool {
		return strings.Compare(names[i], names[j]) < 1
	})

	sum := 0
	for i, name := range names {
		sum += (i + 1) * wordWorth(name)
	}

	return sum, nil
}

func wordWorth(word string) int {
	worth := 0
	for _, r := range word {
		worth += int(r-'A') + 1
	}
	return worth
}

// Pe0023Limit 此より大きい整數は二つの過剰數の和として表せる
var Pe0023Limit = 28123

// PE0023 28123以下の數で二つの過剰數の和で書き表せない正の整數の總和を計算する
//
// https://projecteuler.net/problem=23
func PE0023() int {
	abundants := []int{} // 28123以下の過剰數

	pg := NewPrimeGenerator() // 素數ジェネレータ
	pm := []int{pg.Next()}    // 素數配列

	for i := 12; i <= Pe0023Limit; i++ {
		m, pf := i, map[int]int{}
		for j := 0; m >= pm[j]; j++ {
			if j == len(pm)-1 {
				pm = append(pm, pg.Next())
			}
			for p := pm[j]; m%p == 0; m = m / p {
				pf[p]++
			}
		}
		// 約數の和
		n := 1
		for k, v := range pf {
			n *= (Pow(k, (v+1)) - 1) / (k - 1)
		}
		// 過剰數 (眞の約數の和が元の數より大きい)
		if n > 2*i {
			abundants = append(abundants, i)
		}
	}

	sumList := map[int]int{}
	for i := 0; i < len(abundants); i++ {
		for j := 0; j < len(abundants) && abundants[i]+abundants[j] <= Pe0023Limit; j++ {
			sumList[abundants[i]+abundants[j]]++
		}
	}

	sum := 0
	for i := 1; i < Pe0023Limit; i++ {
		if _, ok := sumList[i]; !ok {
			sum += i
		}
	}

	return sum
}

// PE0024a 10個の数字の順列を辞書式に並べたときにn番目のものを計算する
//
// 先頭が `0` で固定されたときの順列の数は P(9,9) = 9! = 362,880 個。
// 先頭が `2` の数列は
//     1 * 362,880 * 2     =   725,761 番目から
//     1 * 362,880 * 3 - 1 = 1,088,639 番目まで。
//
// 先頭二文字が固定されたときの順列の数は P(8,8) = 8! = 40,320 個。
// `27` で始まる数列は
//     725,761 + 40,320 * 6     =   967,681 番目から
//     725,761 + 40,320 * 7 - 1 = 1,008,000 番目まで。
//
// 先頭三文字が固定されたときの順列の数は P(7,7) = 7! = 5,040 個。
// `278` で始まる数列は
//     967,681 + 5,040 * 6     =   997,921 番目から
//     967,681 + 5,040 * 7 - 1 = 1,002,960 番目まで。
//
// 先頭四文字が固定されたときの順列の数は P(6,6) = 6! = 720 個。
// `2783` で始まる数列は
//     997,921 + 720 * 2     =   999,361 番目から
//     997,921 + 720 * 3 - 1 = 1,000,080 番目まで。
//
// 先頭五文字が固定されたときの順列の数は P(5,5) = 5! = 120 個。
// `27839` で始まる数列は
//     999,361 + 120 * 5     =   999,961 番目から
//     999,361 + 120 * 6 - 1 = 1,000,080 番目まで。
//
// 先頭六文字が固定されたときの順列の数は P(4,4) = 4! = 24 個。
// `278391` で始まる数列は
//     999,961 + 24 * 1     =   999,985 番目から
//     999,961 + 24 * 2 - 1 = 1,000,008 番目まで。
//
// 先頭七文字が固定されたときの順列の数は P(3,3) = 3! = 6 個。
// `2783915` で始まる数列は
//     999,985 + 6 * 2     =   999,997 番目から
//     999,985 + 6 * 3 - 1 = 1,000,002 番目まで。
//
// 先頭八文字が固定されたときの順列の数は P(2,2) = 2! = 2 個。
// `27839154` で始まる数列は
//     999,997 + 2 * 1     =   999,999 番目から
//     999,997 + 2 * 2 - 1 = 1,000,000 番目まで。
//
// 結局 `2783915460` が 1,000,000 番目の数列。
//
// https://projecteuler.net/problem=23
func PE0024a(n int) string {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	answer := []rune{}

	i := 1

	// n番目の順列を絞り込む
	for l := len(digits); l > 0; l = len(digits) {
		for k := 0; k < l; k++ {
			// 順列の数 P(n, n) = n!
			if p := factorial(l - 1); i+k*p <= n && n < i+(k+1)*p {
				answer = append(answer, digits[k])
				digits = append(digits[:k], digits[k+1:]...)
				i += k * p
				break
			}
		}
	}
	return string(answer)
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	product := n
	for i := n - 1; i > 0; i-- {
		product *= i
	}
	return product
}

// PE0024b digitsを入れ替えていくバージョン
func PE0024b(n int) string {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	i, L := 1, len(digits)

	// n番目の順列を絞り込む
	for l := 0; l < L; l++ {
		for k := 0; k < L-l; k++ {
			// 順列の数 P(n, n) = n!
			if p := factorial(L - l - 1); i+k*p <= n && n < i+(k+1)*p {
				for m := k + l; m > l; m-- {
					digits[m-1], digits[m] = digits[m], digits[m-1]
				}
				i += k * p
				break
			}
		}
	}

	return string(digits)
}

// PE0025a 初めて1000桁になるFibonacci数の順番を計算する
//
// https://projecteuler.net/problem=25
func PE0025a(n int) int {

	fibonacciMap := map[int]*big.Int{} // n番目のFibonacci数

	var fibonacci func(int) *big.Int

	fibonacci = func(nth int) *big.Int {
		if nth == 0 || nth == 1 {
			return big.NewInt(int64(1))
		} else if _, ok := fibonacciMap[nth]; ok {
			return fibonacciMap[nth]
		}
		fibonacciMap[nth] = big.NewInt(int64(0)).Add(
			fibonacci(nth-1),
			fibonacci(nth-2),
		)
		return fibonacciMap[nth]
	}

	m := 1
	for ; len(fibonacci(m).String()) < n; m++ {
	}
	return m + 1
}

// PE0025b クロージャを使用せず、初めて1000桁になるFibonacci數を計算
func PE0025b(n int) int {
	fibonacciMap := map[int]*big.Int{} // n番目のFibonacci数

	m := 1
	for len(Fibonacci(m, fibonacciMap).String()) < n {
		m++
	}
	return m + 1
}

// PE0025c 更に共有マップを使用しない
func PE0025c(n int) int {
	m := 1
	for len(Fibonacci(m, nil).String()) < n {
		m++
	}
	return m + 1
}
