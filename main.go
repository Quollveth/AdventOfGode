package main

import (
	"os"

	"github.com/quollveth/AdventOfGode/day1"
	"github.com/quollveth/AdventOfGode/day2"
	"github.com/quollveth/AdventOfGode/day3"
	"github.com/quollveth/AdventOfGode/day4"
	"github.com/quollveth/AdventOfGode/day5"
	"github.com/quollveth/AdventOfGode/day6"
	"github.com/quollveth/AdventOfGode/day7"
)

func main() {
	n := os.Args[1]

	switch n {
	case "1":
		day1.Run()
	case "2":
		day2.Run(true) // part 1
		day2.Run(false)
	case "3":
		day3.Run(true) // part 1
		day3.Run(false)
	case "4":
		day4.Part1() // it took 4 days to think of a non stupid way to do this
		day4.Part2()
	case "5":
		day5.Run()
	case "6":
		day6.Run()
	case "7":
		day7.Run()
	}
}
