package day8

import (
	"fmt"
	"math"

	"github.com/quollveth/AdventOfGode/util"
)

type point struct {
	x, y int
}

// first global variable of the year letsgo
var antinodes map[point]bool = make(map[point]bool)

func Run() {
	antennaLocations, gridSize := getInput("day8/tinyin")

	for _, v := range antennaLocations {
		//fmt.Printf("'%v': %v\n", string(k), v)

		findAntinodes(&v)
	}

	allAntennas := make(map[point]rune)
	for c, locs := range antennaLocations {
		for _, l := range locs {
			allAntennas[l] = c
		}
	}

	for i := range gridSize.x {
		for j := range gridSize.y {
			toPrint := "."
			if _, has := antinodes[point{i, j}]; has {
				toPrint = "#"
			}
			if _, has := allAntennas[point{i, j}]; has {
				toPrint = string(allAntennas[point{i, j}])
			}
			fmt.Print(toPrint)
		}
		fmt.Print("\n")
	}
}

func getInput(which string) (map[rune][]point, point) {
	in := util.ReadFileLines(which)

	inputCols := len(in[0])
	inputRows := len(in)

	antenna := make(map[rune][]point)

	// read the location of every antenna into a map
	for i := range inputRows {
		for j := range inputCols {
			c := rune(in[i][j])

			if rune(c) == '.' {
				continue
			}
			// character not there yet
			if _, has := antenna[c]; !has {
				antenna[c] = make([]point, 0)
			}
			antenna[c] = append(antenna[c], point{x: i, y: j})
		}
	}

	gridSize := point{
		x: inputCols,
		y: inputRows,
	}

	return antenna, gridSize
}

// finds all antinodes are formed by these points
func findAntinodes(nodes *[]point) {
	// we need at least two antenna to have a line
	if len(*nodes) < 2 {
		return
	}

	for _, node := range *nodes {
		for _, next := range *nodes {
			// any two points form a line and have two antinodes
			antis := getAntis(node, next)
			antinodes[antis[0]] = true
			antinodes[antis[1]] = true
		}
	}
}

// // hell yeah more linear algebra
type vec2d struct {
	x, y float64
}

func (v vec2d) scale(scalar float64) vec2d {
	return vec2d{
		x: v.x * scalar,
		y: v.y * scalar,
	}
}

func (v vec2d) euclideanDistance(o vec2d) float64 {
	return math.Sqrt(math.Pow(v.x-o.x, 2) + math.Pow(v.y-o.y, 2))
}

func (v vec2d) add(other vec2d) vec2d {
	return vec2d{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

func (v vec2d) subtract(other vec2d) vec2d {
	return vec2d{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

func vecFromPoint(p point) vec2d {
	// most useless function prize goes to
	return vec2d{
		x: float64(p.x),
		y: float64(p.y),
	}
}

func pointFromVec(v vec2d) point {
	// actually this one is even more useless
	// like just take the decimals away
	return point{
		x: int(v.x),
		y: int(v.y),
	}
}

// any two points form a line and have two antinodes
func getAntis(a, b point) [2]point {
	var antis [2]point

	direction := vecFromPoint(b).subtract(vecFromPoint(a))

	aButVector := vecFromPoint(a)
	bButVector := vecFromPoint(b)

	distance := aButVector.euclideanDistance(bButVector)

	antis[0] = pointFromVec(bButVector.add(direction.scale(distance)))
	antis[1] = pointFromVec(aButVector.subtract(direction.scale(distance)))

	return antis
}
