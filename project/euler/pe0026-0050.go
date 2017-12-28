package euler

import (
	"fmt"
	"math/big"
	"strconv"
)

// PE0026 d<nの内、1/dの循環節が最も長いものを計算する
//
// https://projecteuler.net/problem=26
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

// 循環少數は面白い
// http://www2.math.kindai.ac.jp/~chinen/junkan_f/junkan.html

// PE0027a 素数生成式
//
// |a| < 1000, |b| < 1000 の時, n^2+an+b で n=0 から初めて連続する整数で素数を生成したときに,
// 最も多く素数を生成できるa, bを計算する.
//
// a, b の条件を考える
//     1. n=0 の時, 0^2+a*0+b=b なので, b は素数
//
// https://projecteuler.net/problem=27
func PE0027a(aMax, bMax int) int {
	// 1000未満の素数を生成
	pg := NewPrimeGenerator()
	pm := map[int]bool{} // 素数マップ
	bs := []int{}        // bの取りうる値(bMax未満の素数)
	for p := pg.Next(); p < bMax; p = pg.Next() {
		pm[p] = true
		bs = append(bs, p)
	}

	aResult, bResult, nMax := 0, -1*aMax, -1*bMax

	// a, bを動かす
	for _, b := range bs {
		for a := -1 * (aMax - 1); a <= aMax-1; a++ {
			for n := 0; ; n++ {
				m := n*n + a*n + b
				for p := pg.Next(); p <= m; p = pg.Next() {
					pm[p] = true
				}
				// n^2+a*n+b が素数で在る間は継続
				if _, ok := pm[m]; ok {
					continue
				}
				if n > nMax {
					aResult, bResult, nMax = a, b, n
				}
				break
			}
		}
	}

	return aResult * bResult
}

// PE0028a 対角線上に在る數字の總和を計算
//
// 中央の1を0番目として、其乃外側の環を1番目、其乃外側を2番目とする。
// 3番目の環迄書き出すと。
//
// 43 44 45 46 47 48 49
// 42 21 22 23 24 25 26
// 41 20  7  8  9 10 27
// 40 18  6  1  2 11 28
// 39 18  5  4  3 12 29
// 38 17 16 15 14 13 30
// 37 36 35 34 33 32 31
//
// ### iの走る範圍
//
//    5x   5 の時は i:1->2
//    7x   7 の時は i:1->3
// 1001x1001 の時は i:1->500
//
// ### 部分対角線の和
//
// * 右斜め上の和 3^2 + 5^2 + 7^2 + ... = S(2*i+1)^2
// * 右斜め下の和 S((2*i+1)^2-6*i)
// * 左斜め下の和 S((2*i+1)^2-4*i)
// * 左斜め上の和 S((2*i+1)^2-2*i)
//
// ### 環の和
//
// i番目の環の和は
//
// (2*i+1)^2+(2*i+1)^2-6*i+(2*i+1)^2-4*i+(2*i+1)^-2*i
//     = 4*((2*i+1)^2)-6*i-4*i-2*i
//     = 4*((2*i+1)^2)-12*i
//
// 1番目の環 4*((2*1+1)^2)-12*1 =  24
// 2番目の環 4*((2*2+1)^2)-12*2 =  76
// 3番目の環 4*((2*3+1)^2)-12*3 = 160
//
// ### n番目の環迄の合計 s(n)
//
// s(n) = 1+S(i:1->500)(4*((2*i+1)^2)-12*i)
//
// https://projecteuler.net/problem=28
func PE0028a(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	sum := 1
	for i := 1; i <= (n-1)/2; i++ {
		sum += 4*(2*i+1)*(2*i+1) - 12*i
	}
	return sum
}

// PE0029a 2<=a, b<=n の時、a^b の個數(重複を除く)を計算
//
// https://projecteuler.net/problem=29
func PE0029a(n int) int {
	if n < 2 {
		return 0
	}

	m := map[string]bool{}

	for a := 2; a <= n; a++ {
		aBig := big.NewInt(int64(a))
		for b := 2; b <= n; b++ {
			p := big.NewInt(int64(1))
			for i := 1; i <= b; i++ {
				p.Mul(p, aBig)
			}
			m[p.String()] = true
		}
	}

	return len(m)
}

// PE0030a 各桁の五乘の和が元の數になるやうな數の總和を計算.
//
// 各桁の四乘の和が元の數になるのは以下の三つのみ.
//
// * 1634 = 1^4 + 6^4 + 3^4 + 4^4
// * 8208 = 8^4 + 2^4 + 0^4 + 8^4
// * 9474 = 8^4 + 4^4 + 7^4 + 4^4
//
// 各桁を三乗したものの和
//
//     9 = 1 * 9^3 =   729
//    99 = 2 * 9^3 = 1,458
//   999 = 3 * 9^3 = 2,187
// 9,999 = 4 * 9^3 = 2,916
//
// 各桁を四乗したものの和
//
//     99 = 2 * 9^4 = 13,122
//    999 = 3 * 9^4 = 19,683
//  9,999 = 4 * 9^4 = 26,244
// 99,999 = 5 * 9^4 = 32,805
//
// 各桁を五乗したものの和
//
//     9,999 = 4 * 9^5 = 236,196
//    99,999 = 5 * 9^5 = 295,245
//   999,999 = 6 * 9^5 = 354,294
// 9,999,999 = 7 * 9^5 = 413,343
//
// https://projecteuler.net/problem=30
func PE0030a(e int) int {
	supremum := 10 * Pow(9, e) // 探索上限

	results := []int{} // 結果記録用

	for n := 2; n <= supremum; n++ {
		result := 0
		for i := n; i > 0; i /= 10 {
			result += Pow(i%10, e)
		}
		if n == result {
			results = append(results, n)
		}
	}

	sum := 0
	for _, n := range results {
		sum += n
	}

	return sum
}

// PE0031a 2Lになる硬貨の組み合わせの個数を計算
//
// 英国で流通している硬貨は以下の八種類
//
// 1p 2p 5p 10p 20p 50p L1, L2
//
// https://projecteuler.net/problem=31
func PE0031a() int {
	// 合計が2Lになる組み合わせ
	m := map[string]bool{}

	// 残りの金額
	remaining := func(x200, x100, x50, x20, x10, x5, x2, x1 int) int {
		return 200 - (200*x200 + 100*x100 + 50*x50 + 20*x20 + 10*x10 + 5*x5 + 2*x2 + 1*x1)
	}

	for x200 := 0; x200 <= remaining(0, 0, 0, 0, 0, 0, 0, 0)/200; x200++ {
		for x100 := 0; x100 <= remaining(x200, 0, 0, 0, 0, 0, 0, 0)/100; x100++ {
			for x50 := 0; x50 <= remaining(x200, x100, 0, 0, 0, 0, 0, 0)/50; x50++ {
				for x20 := 0; x20 <= remaining(x200, x100, x50, 0, 0, 0, 0, 0)/20; x20++ {
					for x10 := 0; x10 <= remaining(x200, x100, x50, x20, 0, 0, 0, 0)/10; x10++ {
						for x5 := 0; x5 <= remaining(x200, x100, x50, x20, x10, 0, 0, 0)/5; x5++ {
							for x2 := 0; x2 <= remaining(x200, x100, x50, x20, x10, x5, 0, 0)/2; x2++ {
								for x1 := 0; x1 <= remaining(x200, x100, x50, x20, x10, x5, x2, 0)/1; x1++ {
									if remaining(x200, x100, x50, x20, x10, x5, x2, x1) == 0 {
										m[fmt.Sprintf("%d:%d:%d:%d:%d:%d:%d:%d", x200, x100, x50, x20, x10, x5, x2, x1)] = true
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return len(m)
}

// PE0032a パンデジタル数
//
// m桁の正の整数Amとn (>= m)桁の正の整数Bnの積Cの桁数ははm+n-1 or m+nのいずれか。
//
// 1) m + n + (m + n - 1) = 2 * (m + n) - 1 = 9 -> m + n = 10/2 = 5
//    (m, n) = (1, 4), (2, 3) が存在する。
//    (m, n) = (1, 4) の時、積Cの桁数は4
//    (m, n) = (2, 3) の時、積Cの桁数は4
//
// 2) m + n + (m + n) = 2 * (m + n) = 9 -> m + n = 9/2
//    上記を満たすm, nの組は存在しない。
//    こちらは考えなくて良い。
//
// https://projecteuler.net/problem=32
func PE0032a() int {
	used := map[int]bool{} // 既に加算した数
	sum := 0               // 総和
	rss := Permutation([]rune("123456789"), 5)
	for _, rs := range rss {
		// (m, n) = (1, 4)
		a0, _ := strconv.Atoi(string(rs[:1]))
		b0, _ := strconv.Atoi(string(rs[1:5]))
		c0 := a0 * b0
		if _, ok := used[c0]; !ok && IsPandigital10(100000*c0+10*b0+a0) {
			used[c0] = true
			sum += c0
		}
		// (m, n) = (2, 3)
		a1, _ := strconv.Atoi(string(rs[:2]))
		b1, _ := strconv.Atoi(string(rs[2:5]))
		c1 := a1 * b1
		if _, ok := used[c1]; !ok && IsPandigital10(100000*c1+100*b1+a1) {
			used[c1] = true
			sum += c1
		}
	}
	return sum
}

// PE0033 特殊な分數の積の分母を計算
//
// https://projecteuler.net/problem=33
func PE0033() {
}
