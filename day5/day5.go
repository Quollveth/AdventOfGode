package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/quollveth/AdventOfGode/util"
)

/*
	read rules into a int -> []int map
	keeps track of which pages have to come before a given page
	13 -> 97,61,29,47,75,53
	means these pages have to come before 13
	read updates into [][]int
	go through each and check each page if any pages ahead should be before
*/

func readInput() (map[int][]int, [][]int) {
	in := util.ReadFileLines("day5/input")
	rules := make(map[int][]int)
	updates := [][]int{}

	atRules := true
	for _, line := range in {
		if line == "" {
			atRules = false
			continue
		}
		if atRules {
			// read rules
			// 47|53 will have 47 added to 53's slice
			r := strings.Split(line, "|")
			r1, _ := strconv.Atoi(r[0])
			r2, _ := strconv.Atoi(r[1])

			rules[r2] = append(rules[r2], r1)
			continue
		}
		// read updates
		// each lin is just a comma separated list of numbers
		arr := []int{}
		n := strings.Split(line, ",")
		for _, num := range n {
			x, _ := strconv.Atoi(num)
			arr = append(arr, x)
		}
		updates = append(updates, arr)
	}

	return rules, updates
}

func contains(s []int, v int) bool {
	for _, n := range s {
		if n == v {
			return true
		}
	}
	return false
}

func Part1() {
	rules, updates := readInput()
	var valid bool
	sum := 0
	for _, update := range updates {
		valid = true
		for j, page := range update {
			// check all pages in front if they should not be there
			for _, pageAfter := range update[j:] {
				if slice, exists := rules[page]; exists {
					if contains(slice, pageAfter) {
						valid = false
						// entire update is invalid and we can stop checking it
						goto breakBoth
					}
				}
			}
		}
	breakBoth:
		if !valid {
			continue
		}
		// update is valid
		median := int(len(update) / 2)
		sum += update[median]
	}

	fmt.Println(sum)
}
