package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"
	tour "github.com/py0n/exercise/a_tour_of_go"
)

var opts tour.Options

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "tour"
	parser.Usage = "CMDNO ..."
	args, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}
	if len(args) < 1 {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	tour.Run(opts, args)
	os.Exit(0)
}
