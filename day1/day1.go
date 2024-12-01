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

	sll := make([]int, len(ll))
	srl := make([]int, len(rl))

	copy(sll, ll)
	copy(srl, rl)

	sort.Slice(sll, func(i, j int) bool {
		return sll[i] < sll[j]
	})

	sort.Slice(srl, func(i, j int) bool {
		return srl[i] < srl[j]
	})

	s := 0
	for i := range sll {
		s += absdiff(sll[i], srl[i])
	}

	fmt.Println("Part 1 solution:", s)

	//////// PART 2

	n := 0
	sc := 0
	for _, i := range ll {
		n = 0
		for _, j := range rl {
			if i == j {
				n++
			}
		}
		sc += i * n
	}

	fmt.Println("Part 2 solution:", sc)
}
