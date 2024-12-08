package day7

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/quollveth/AdventOfGode/util"
)

/*
go from right to left working through the array
divide and subtract the first item in the array from the value
until we have only one number left
if it is the current value it's valid

eg:
	292: 11 6 16 20
	: 292 20 16 6 11
	292 - 20 = 272 | 16 6 11
	272 / 16 = 17  | 6 11
	17 - 6 = 11    | 11
	result is equal to last number remaining, so this path is valid
	as long as one path down the tree is valid the whole tree is valid
*/

func Run() {
	input := util.ReadFileLines("day7/input")

	var count int64

	var wg sync.WaitGroup
	// dispatch one thread per line
	for _, line := range input {
		wg.Add(1)

		go func(cur string) {
			defer wg.Done()

			n := processRow(cur)
			atomic.AddInt64(&count, n)

		}(line)
	}
	wg.Wait()

	fmt.Println(count)
}

// Check if this row has valid operands
func processRow(row string) int64 {
	// Parse data
	// We need to parse and reverse, we can do both at once
	splits := strings.Fields(row)
	target, _ := strconv.Atoi(strings.TrimSuffix(splits[0], ":"))
	nums := make([]int, len(splits)-1)
	for i, j := len(splits)-1, 0; i >= 1; i, j = i-1, j+1 {
		nums[j], _ = strconv.Atoi(splits[i])
	}

	if validateLine(target, nums) {
		//fmt.Printf("\033[32m[%v] - Valid\n\033[0m", row)
		return int64(target)
	}
	//fmt.Printf("\033[31m[%v] - Invalid\n\033[0m", row)
	return 0
}

func validateLine(current int, remain []int) bool {
	// shouldnt happen
	if len(remain) == 0 {
		return false
	}

	next := remain[0]

	// base case
	if len(remain) == 1 {
		// if last element is root this path is valid
		return current == next
	}

	// recursive case
	// check division first as it can be easily discarded if it's a non integer result
	if current%next == 0 {
		if validateLine(current/next, remain[1:]) {
			return true
		}
	}
	return validateLine(current-next, remain[1:])
}