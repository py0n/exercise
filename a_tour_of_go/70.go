package tour

import (
	"fmt"

	"github.com/golang/tour/tree"
)

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		walk(t2, ch2)
		close(ch2)
	}()
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 && !ok2 {
			return true
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}

// C70 ...
func C70(opts Options, args []string) int {
	t1 := tree.New(1)
	t2 := tree.New(1)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(same(t1, t2))
	return 0
}
