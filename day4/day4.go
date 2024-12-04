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

	for i := 0; i < cols; i++ {
		matrix[i] = make([]rune, rows)
		for j := 0; j < rows; j++ {
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
			fmt.Printf(string(char))
		}
		fmt.Printf("\n")
	}
}

func Run() {
	in := util.ReadFileLines("day4/numbers")
	inputCols := len(in[0])
	inputRows := len(in)

	input := make([][]rune, inputRows)
	for i := range inputRows {
		input[i] = make([]rune, inputCols)
		for j := range inputCols {
			input[i][j] = rune(in[i][j])
		}
	}

	printMatrix(input)
	fmt.Println("------------")
	printMatrix(rotateMatrix(input))
}
