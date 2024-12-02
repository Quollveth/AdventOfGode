package main

import (
	"os"

	"github.com/quollveth/AdventOfGode/day1"
	"github.com/quollveth/AdventOfGode/day2"
)

func main() {
	n := os.Args[1]

	switch n {
	case "1":
		day1.Run()
	case "2":
		day2.Run()
	}
}
