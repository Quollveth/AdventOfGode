package day8

import (
	"fmt"

	"github.com/quollveth/AdventOfGode/util"
)

type point struct {
	x, y int
}

/*
 for each antenna we want to
 - get all its positions
 - get all unique pairs of positions
 - calculate the antis for those pairs
 - ignore if its out of bounds
*/

func Run() {
	in := util.ReadFileLines("day8/input")

	gridSize := point{
		x: len(in[0]),
		y: len(in),
	}

	// the locations of each antenna frequency
	antenna := make(map[rune][]point)

	for j := range gridSize.y {
		for i := range gridSize.x {
			c := rune(in[j][i])
			if c == '.' {
				continue
			}
			l := point{j, i}
			antenna[c] = append(antenna[c], l)
		}
	}

	// assigns the location of an antinode with the fact it exists
	antinodes := make(map[point]bool)

	for _, positions := range antenna {
		pairs := util.Combinations(positions)
		for _, pair := range pairs {
			antis := getAntis(pair[0], pair[1])

			if validAnti(antis[0], gridSize) {
				antinodes[antis[0]] = true
			}

			if validAnti(antis[1], gridSize) {
				antinodes[antis[1]] = true
			}
		}
	}

	fmt.Println(len(antinodes))
}

func validAnti(anti point, gridSize point) bool {
	if !util.InRange(anti.x, 0, gridSize.x) {
		return false
	}
	if !util.InRange(anti.y, 0, gridSize.y) {
		return false
	}

	return true
}

// any two points form a line and have two antinodes
func getAntis(a, b point) [2]point {
	var antis [2]point

	/*
		d = b - a
		return [a-d,b+d]
	*/

	delta := pointSub(b, a)

	antis[0] = pointSub(a, delta)
	antis[1] = pointAdd(b, delta)

	return antis
}

func pointAdd(a, b point) point {
	return point{
		x: a.x + b.x,
		y: a.y + b.y,
	}
}

func pointSub(a, b point) point {
	return point{
		x: a.x - b.x,
		y: a.y - b.y,
	}
}
