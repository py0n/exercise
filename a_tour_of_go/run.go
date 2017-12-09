package tour

import (
	"fmt"
	"os"
	"strconv"
)

// Run ...
func Run(opts Options, args []string) {
	chapter, _ := strconv.ParseInt(args[0], 10, 64)
	switch chapter {
	case 24:
		os.Exit(C24(opts, args))
	case 41:
		os.Exit(C41(opts, args))
	case 44:
		os.Exit(C44(opts, args))
	case 48:
		os.Exit(C48(opts, args))
	case 56:
		os.Exit(C56(opts, args))
	case 57:
		os.Exit(C57(opts, args))
	case 58:
		os.Exit(C58(opts, args))
	case 61:
		os.Exit(C61(opts, args))
	case 70:
		os.Exit(C70(opts, args))
	case 71:
		os.Exit(C71(opts, args))
	default:
		fmt.Println(opts)
		fmt.Println(args)
		os.Exit(1)
	}
}
