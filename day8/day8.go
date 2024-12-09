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

func printGrid(antennas map[point]rune, size point, color int) {
	for i := range size.x {
		for j := range size.y {
			tp := "."
			pl := point{i, j}

			if _, has := antinodes[pl]; has {
				tp = "#"
			}
			if _, has := antennas[pl]; has {
				tp = string(antennas[pl])
			}

			fmt.Printf("\033[%vm%v\033[39m", color, tp)
		}
		fmt.Println()
	}
	fmt.Println()
}

func Run() {
	antennaLocations, antennas, gridSize := getInput("day8/tinyin")
	_, _, _ = gridSize.x, gridSize.y, antennas

	for k, v := range antennaLocations {
		_ = k

		findAntinodes(&v)
	}

	// remove all antis out of bounds or in an antenna
	for anti := range antinodes {
		if !inRange(anti.x, 0, gridSize.x) {
			delete(antinodes, anti)
		}
		if !inRange(anti.y, 0, gridSize.y) {
			delete(antinodes, anti)
		}
		if _, has := antennas[anti]; has {
			delete(antinodes, anti)
		}
	}

	printGrid(antennas, gridSize, 32)

	fmt.Println(len(antinodes))

}

func inRange(a, min, max int) bool {
	return min < a && a < max
}

func getInput(which string) (map[rune][]point, map[point]rune, point) {
	in := util.ReadFileLines(which)

	inputCols := len(in[0])
	inputRows := len(in)

	antenna := make(map[rune][]point)
	allAntennas := make(map[point]rune)

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
			p := point{x: i, y: j}
			antenna[c] = append(antenna[c], p)
			allAntennas[p] = c
		}
	}

	gridSize := point{
		x: inputCols,
		y: inputRows,
	}

	return antenna, allAntennas, gridSize
}

// finds all antinodes are formed by these points
func findAntinodes(nodes *[]point) {
	// we need at least two antenna to have a line
	if len(*nodes) < 2 {
		return
	}

	for _, comb := range combinations(*nodes) {
		antis := getAntis(comb[0], comb[1])

		antinodes[antis[0]] = true
		antinodes[antis[1]] = true
	}
}

// every unique unordered pair
func combinations(points []point) [][2]point {
	comb := [][2]point{}
	plen := len(points)

	for i := range plen - 1 {
		for j := i + 1; j < plen; j++ {
			comb = append(comb, [2]point{points[i], points[j]})
		}
	}

	return comb
}

// // hell yeah more linear algebra
type vec2d struct {
	x, y float64
}

func (v vec2d) magnitude() float64 {
	return math.Sqrt((v.x * v.x) + (v.y * v.y))
}

func (v vec2d) normalize() vec2d {
	mag := v.magnitude()
	return vec2d{
		x: v.x / mag,
		y: v.y / mag,
	}
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

}
