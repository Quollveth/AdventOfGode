package day4

import (
	"fmt"

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

func makeWindow(orig *[][]rune, centerRow, centerCol int) [][]rune {
	lineLen := len((*orig)[0])
	lines := len((*orig))

	newMatrix := make([][]rune, 7)
	for i := range 7 {
		newMatrix[i] = make([]rune, 7)
		for j := range 7 {
			// out of bounds indexes are filled with 1 as square matrices are easier and thats not a real rune
			newMatrix[i][j] = 1
		}
	}

	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			origRow := centerRow + i - 3
			origCol := centerCol + j - 3

			if inRange(0, lines, origRow) && inRange(0, lineLen, origCol) {
				newMatrix[i][j] = (*orig)[origRow][origCol]
			}
		}
	}

	return newMatrix
}

func Run() {
	in := util.ReadFileLines("day4/input")

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

				window = makeWindow(&input, i, j)

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
