package tour

import "fmt"

func fibonacci() func() int {
	f0, f1, f2 := 0, 1, 0
	return func() int {
		f2 = f0 + f1
		f0 = f1
		f1 = f2
		return f2
	}
}

// C44 ...
func C44(opts Options, args []string) int {
	if len(args) != 1 {
		return 1
	}
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	return 0
}
