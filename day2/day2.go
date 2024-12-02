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

func Run() {
	input := util.ReadFileLines("day2/input")

	nSafe := 0
	for _, line := range input {
		splits := strings.Split(line, " ")

		var prev int
		increasing := true
		safe := true

		for i, v := range splits {
			level, _ := strconv.Atoi(v)

			if i == 0 {
				prev = level
				continue
			}

			diff := absdiff(prev, level)

			if diff == 0 || diff > 3 {
				safe = false
				break
			}

			if i == 1 {
				increasing = level > prev

				prev = level
				continue
			}

			if (level > prev) != increasing {
				safe = false
				break
			}

			prev = level
		}

		if safe {
			nSafe++
		}
	}

	fmt.Println(nSafe)
}
