package day1

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/quollveth/AdventOfGode/util"
)

func absdiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func Run() {
	input := util.ReadFileWords("day1/input")

	var ll, rl []int

	for i, e := range input {
		v, _ := strconv.Atoi(e)
		if i%2 == 0 {
			ll = append(ll, v)
			continue
		}
		rl = append(rl, v)
	}

	sort.Slice(ll, func(i, j int) bool {
		return ll[i] < ll[j]
	})

	sort.Slice(rl, func(i, j int) bool {
		return rl[i] < rl[j]
	})

	var s int = 0
	for i := range ll {
		s += absdiff(ll[i], rl[i])
	}

	fmt.Println(s)
}
