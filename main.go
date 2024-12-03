package main

import (
	"os"

	"github.com/quollveth/AdventOfGode/day1"
	"github.com/quollveth/AdventOfGode/day2"
	"github.com/quollveth/AdventOfGode/day3"
)

func main() {
	n := os.Args[1]

	switch n {
	case "1":
		day1.Run()
	case "2":
		day2.Run(true)
		day2.Run(false)
	case "3":
		day3.Run(true)
		day3.Run(false)
	}
}
