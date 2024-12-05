package day4

import (
	"fmt"
	"log"

	"github.com/quollveth/AdventOfGode/util"
)

func countOccurrences(arr []string, val string) int {
	count := 0
	for _, e := range arr {
		if e == val {
			count++
		}
	}
	return count
}

func transposeMatrix(orig [][]rune) [][]rune {
	rows, cols := len(orig), len(orig[0])
	matrix := make([][]rune, cols)

	for i := range cols {
		matrix[i] = make([]rune, rows)
		for j := range rows {
			matrix[i][j] = orig[j][i]
		}
	}

	return matrix
}

func reverseArray(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// https://stackoverflow.com/a/8664879
func rotateMatrix(orig [][]rune) [][]rune {
	rot := transposeMatrix(orig)

	for i := range rot {
		rot[i] = reverseArray(rot[i])
	}

	return rot
}

func printMatrix(input [][]rune) {
	for _, line := range input {
		for _, char := range line {
			if char == 1 {
				fmt.Printf(".")
				continue
			}
			fmt.Printf(string(char))
		}
		fmt.Printf("\n")
	}
}

func inRange(min, max, val int) bool {
	return val >= min && val < max
}

func makeWindow(orig *[][]rune, centerRow, centerCol, size int) [][]rune {
	if size%2 == 0 {
		log.Panic("Window size must be an odd number")
	}

	lineLen := len((*orig)[0])
	lines := len((*orig))

	// Create a new square matrix of the specified size
	newMatrix := make([][]rune, size)
	for i := range newMatrix {
		newMatrix[i] = make([]rune, size)
		for j := range newMatrix[i] {
			// Fill out-of-bounds indexes with 1 as placeholder
			newMatrix[i][j] = 1
		}
	}

	// Calculate the offset for centering
	offset := size / 2

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			origRow := centerRow + i - offset
			origCol := centerCol + j - offset

			if inRange(0, lines, origRow) && inRange(0, lineLen, origCol) {
				newMatrix[i][j] = (*orig)[origRow][origCol]
			}
		}
	}

	return newMatrix
}

const (
	real = "day4/input"
	test = "day4/testin"
	tiny = "day4/tinyin"
)

func getInput(which string) [][]rune {
	in := util.ReadFileLines(which)

	// turn input into byte matrix to make working with it easier
	inputCols := len(in[0])
	inputRows := len(in)

	input := make([][]rune, inputRows)
	for i := range inputRows {
		input[i] = make([]rune, inputCols)
		for j := range inputCols {
			input[i][j] = rune(in[i][j])
		}
	}
	return input
}

func Part1() {
	input := getInput(real)
	/*
		make a 7x7 window into the input matrix
		X in the middle and MAS in each direction
		if we are at the edges and can't fit the window shrink it
		move window through input until X is at the middle
		check for left to right horizontal and left to right top to bottom diagonal
		rotate window 4 times and repeat checks


		in practice we just move through the matrix until a X is found and create the window there
	*/

	var window [][]rune
	count := 0
	//	var windowLeft, windowTop int

	for i, line := range input {
		for j, char := range line {
			if char == 'X' {

				window = makeWindow(&input, i, j, 7)

				// check horizontal and diagonal
				// rotate matrix 90°
				// check again
				for range 4 {
					// because makeWindow makes a 7x7 square matrix these checks are easy
					// horizontal first
					attempt := window[3][3:]
					if string(attempt) == "XMAS" || string(attempt) == "SAMX" {
						count++
					}
					// diagonal
					attempt = []rune{}

					for k := range 4 {
						attempt = append(attempt, window[3+k][3+k])
					}

					if string(attempt) == "XMAS" || string(attempt) == "SAMX" {
						count++
					}

					// rotate
					window = rotateMatrix(window)
				}
			}
		}
	}

	fmt.Println(count)
}

func matchMas(window *[][]rune) bool {
	// part 1 solution was general enough i earn a little stupid, as a treat
	// as the window gets rotated this matches all possible positions (theres only 4)
	var MAS = [][]rune{
		{'M', '.', 'S'},
		{'.', 'A', '.'},
		{'M', '.', 'S'},
	}

	for i := range 3 {
		for j := range 3 {
			if MAS[i][j] == '.' { // wildcard
				continue
			}
			if MAS[i][j] != (*window)[i][j] {
				return false
			}
		}
	}
	return true
}

func Part2() {
	input := getInput(real)
	fmt.Println()

	var window [][]rune
	count := 0
	for i, line := range input {
		for j, char := range line {
			if char == 'A' {
				window = makeWindow(&input, i, j, 3)

				for range 4 {
					if matchMas(&window) {
						count++
						break // stop on match to avoid overcounting
					}
					window = rotateMatrix(window)
				}
			}
		}
	}
	fmt.Println(count)

}
