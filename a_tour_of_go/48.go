package tour

import (
	"fmt"
	"math/cmplx"
	"strconv"
)

func cbrtByNewtonMethod(x complex128) complex128 {
	z0 := x / 2
	for {
		z1 := z0 - (cmplx.Pow(z0, 3)-x)/(3*cmplx.Pow(z0, 2))
		if cmplx.Abs(z1-z0) < 1e-15 {
			return z1
		}
		z0 = z1
	}
}

// C48 ...
func C48(opts Options, args []string) int {
	if len(args) != 2 {
		return 1
	}
	x, _ := strconv.ParseFloat(args[1], 64)
	fmt.Println(args[1])
	fmt.Println(cbrtByNewtonMethod(complex(x, 0)))
	return 0
}
