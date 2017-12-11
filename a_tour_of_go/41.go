package tour

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Fields(s) {
		m[w]++
	}
	return m
}

// C41 ...
func C41(opts Options, args []string) int {
	if len(args) != 2 {
		return 1
	}
	fmt.Println(args[1])
	fmt.Println(wordCount(args[1]))
	return 0
}
