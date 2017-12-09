package tour

import (
	"fmt"
	"math"
	"strconv"
)

type errNegativeSqrt float64

func (e errNegativeSqrt) Error() string {
	return "cannot Sqrt negative numer: " + strconv.FormatFloat(
		float64(e), 'E', -1, 64,
	)
}

func sqrtByNewtonMethod2(x float64) (float64, error) {
	if x < 0 {
		return math.NaN(), errNegativeSqrt(x)
	} else if x == 0 {
		return 0, nil
	} else {
		return sqrtByNewtonMethod(x), nil
	}
}

// C56 ...
func C56(opts Options, args []string) int {
	if len(args) != 2 {
		return 1
	}
	x, _ := strconv.ParseFloat(args[1], 64)
	fmt.Println(x)
	r, err := sqrtByNewtonMethod2(x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
	return 0
}
