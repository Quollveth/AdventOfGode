package main

import (
	"os"

	"github.com/quollveth/AdventOfGode/day1"
)

func main() {
	n := os.Args[1]

	switch n {
	case "1":
		day1.Run()
	}
}
