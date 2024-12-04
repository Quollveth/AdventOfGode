package day4

import (
	"fmt"

	"github.com/quollveth/AdventOfGode/util"
)

/*
	the input is rectangular (all lines are the same lenght)
	both xmas and samx are valid
	a vertical xmas is composed by the letters all having the same offset from start of line
	a diagonal is the same logic but the offset can only vary by one for each subsequent letter

	**X**
	**M**
	**A**
	**S**

	each letter is 2 characters away from line start
	this value can vary by 1 or -1 from the previous offset to form a diagonal

	X***
	*M**
	**A*
	***S
*/

func countOccurrences(arr []string, val string) int {
	count := 0
	for _, e := range arr {
		if e == val {
			count++
		}
	}
	return count
}

func Run() {
	input := util.ReadFileLines("day4/testin")
	// ^[]string

	count := 0
	inputVer := len(input)
	inputHor := len(input[0])

	var attempts []string
	for i, line := range input {
		for j, char := range line {
			attempts = []string{}

			c := rune(char)

			if c == 'X' {
				////// horizontal
				// left to right
				if j <= inputHor-4 {
					attempts = append(attempts, line[j:j+4])
				}
				// right to left
				if j >= 4 {
					attempts = append(attempts, line[j-3:j+1])
				}

				////// vertical
				// top to bottom
				if i <= inputVer-4 {
					curr := ""
					for k := range 4 {
						curr += string(input[i+k][j])
					}
					attempts = append(attempts, curr)
				}
				// bottom to top
				if i >= 4 {
					curr := ""
					for k := range 4 {
						curr += string(input[i-k][j])
					}
					attempts = append(attempts, curr)
				}

				////// diagonal
				//// top to bottom
				// left to right
				if i <= inputVer-4 && j <= inputHor-4 {
					curr := ""
					for k := range 4 {
						curr += string(input[i+k][j+k])
					}
					attempts = append(attempts, curr)
				}
				// right to left
				if i <= inputVer-4 && j >= 4 {
					curr := ""
					for k := range 4 {
						curr += string(input[i+k][j-k])
					}
					attempts = append(attempts, curr)
				}

				//// bottom to top
				// left to right
				if i >= 4 && j <= inputHor-4 {
					curr := ""
					for k := range 4 {
						curr += string(input[i-k][j+k])
					}
					attempts = append(attempts, curr)
				}
				// right to left
				if i >= 4 && j >= 4 {
					curr := ""
					for k := range 4 {
						curr += string(input[i-k][j-k])
					}
					attempts = append(attempts, curr)
				}

				////// Check attempts
				count += countOccurrences(attempts, "XMAS")
				count += countOccurrences(attempts, "SAMX")
			}
		}
	}

	fmt.Println(count)
}
