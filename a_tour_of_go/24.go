package tour

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func sqrtByNewtonMethod(x float64) float64 {
	if x < 0 {
		return math.NaN()
	}
	if x == 0 {
		return 0
	}
	z0, z1 := x, 0.0
	for {
		z1 = z0 - (z0*z0-x)/(2*z0)
		if math.Abs(z1-z0) < 1e-15 {
			return z1
		}
		z0 = z1
	}
}

// C24 ...
func C24(opts Options, args []string) int {
	if len(args) != 2 {
		return 1
	}
	x, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 0
	}
	fmt.Println(x)
	fmt.Println(sqrtByNewtonMethod(x))
	return 0
}
