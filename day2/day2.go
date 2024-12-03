package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/quollveth/AdventOfGode/util"
)

func absdiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func removeAtIndex(slice []int, index int) []int {
	// no bound check just dont call a oob index
	new := make([]int, 0, len(slice)-1)
	new = append(new, slice[:index]...)
	new = append(new, slice[index+1:]...)

	return new
}

func checkReport(report []int) bool {
	prev := 0
	increasing := false
	for i, level := range report {
		if i == 0 {
			prev = level
			continue
		}

		diff := absdiff(prev, level)

		if diff == 0 || diff > 3 {
			return false
		}

		if i == 1 {
			increasing = level > prev

			prev = level
			continue
		}

		if (level > prev) != increasing {
			return false
		}

		prev = level
	}
	return true
}

func Run(part1 bool) {
	input := util.ReadFileLines("day2/input")

	nSafe := 0
	for _, line := range input {
		// convert to int array
		var report []int
		for _, level := range strings.Split(line, " ") {
			v, _ := strconv.Atoi(level)
			report = append(report, v)
		}

		valid := checkReport(report)

		if valid {
			nSafe++
			continue
		}

		if part1 {
			continue
		}

		// time complexity is O(shit)
		for i := range len(report) {
			if checkReport(removeAtIndex(report, i)) {
				nSafe++
				break
			}
		}
	}

	fmt.Println(nSafe)
}
