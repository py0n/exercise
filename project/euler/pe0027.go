package euler

// PE0027 素数生成式
//
// |a| < 1000, |b| < 1000 の時, n^2+an+b で n=0 から初めて連続する整数で素数を生成したときに,
// 最も多く素数を生成できるa, bを計算する.
//
// a, b の条件を考える
//     1. n=0 の時, 0^2+a*0+b=b なので, b は素数
func PE0027() {

}

func pgFuncGenerator(a, b int) func(int) int {
	return func(n int) int {
		return n*n + a*n + b
	}
}
