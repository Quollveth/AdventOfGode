package day6

import (
	"fmt"
	"sort"

	"github.com/quollveth/AdventOfGode/util"
)

func absdiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

/*
	find starting position of guard
	find all obstacles
	find obstacle in the same direction as guard and count distance
	repeat until no more obstacles are in that direction
*/

// stores a pair of coordinates
type point struct {
	x, y int
}

type guard struct {
	position point
	// guard directions are relative to the grid, so positive Y is down and positive X is right
	direction point
}

func (p point) String() string {
	return fmt.Sprintf("[%v,%v]", p.x, p.y)
}

// directions are set in relation to the matrix
// positive y is down | positive x is right

// modified from day 4 solution
func getInput(which string) (
	[]point, // obstacles
	guard,
	point, // size of input
) {
	in := util.ReadFileLines(which)

	inputCols := len(in[0])
	inputRows := len(in)

	size := point{
		x: inputCols,
		y: inputRows,
	}

	obstacles := []point{}
	guard := guard{}

	for i := range inputRows {
		for j := range inputCols {
			c := in[i][j]
			switch c {
			case '#':
				obstacles = append(obstacles, point{x: j, y: i})
			// guard directions are relative to the grid, so positive Y is down and positive X is right
			case '^':
				guard.position.x = j
				guard.position.y = i
				guard.direction.y = -1
				guard.direction.x = 0
			case 'V':
				guard.position.x = j
				guard.position.y = i
				guard.direction.y = 1
				guard.direction.x = 0

			case '>':
				guard.position.x = j
				guard.position.y = i
				guard.direction.y = 0
				guard.direction.x = 1
			case '<':
				guard.position.x = j
				guard.position.y = i
				guard.direction.y = 0
				guard.direction.x = -1
			}
		}
	}

	return obstacles, guard, size
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func getHit(arr []point, smaller bool) point {
	if arr[0].x == arr[1].x {
		// sort by Y
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].y < arr[j].y
		})
	} else {
		// sort by x
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].x < arr[j].x
		})
	}

	if smaller {
		return arr[len(arr)-1]
	}
	return arr[0]

}

func Run() {
	// this reads into a rune matrix
	obstacles, guard, gridSize := getInput("day6/input")
	visited := make(map[point]bool)
	visited[guard.position] = true

	for { // loop until guard is out of the area
		var (
			matchX, // obstacle should have the same X coordinate
			matchY, // obstacle should have the same Y coordinate
			smaller bool // obstacle should have a coordinate smaller than the guard's
		)

		potential := []point{} // potential obstacles hit
		var hit point          //one we actually hit

		// vertically
		matchX = guard.direction.x == 0
		// horizontally
		matchY = guard.direction.y == 0

		// up or left
		smaller = (guard.direction.y == -1 || guard.direction.x == -1)

		// Find all obstacles
		for _, ob := range obstacles {
			// coordinates don't match
			if matchX && ob.x != guard.position.x {
				continue
			}
			if matchY && ob.y != guard.position.y {
				continue
			}

			// obstacle is behind the guard
			// X is the same, Y is bigger and should be smaller, or is smaller and should be bigger
			if matchX && ((ob.y > guard.position.y) == smaller) {
				continue
			}
			// same but Y is the same
			if matchY && ((ob.x > guard.position.x) == smaller) {
				continue
			}

			// obstacle is in front of the guard, so it's a potential hit
			potential = append(potential, ob)
		}

		// potential hits acquired
		// if theres 0 we're done and can leave the area
		// if theres 1 thats the hit
		// if theres more sort the array and get the first or last
		// otherwise we need to sort the array and pick the proper one depending on direction

		var dist int
		done := false

		// get distance we need to walk
		switch len(potential) {
		case 0:
			done = true
			// get guard direction and calculate distance to edge of area
			if matchX && smaller {
				// up
				dist = guard.position.y
			}
			if matchX && !smaller {
				//down
				dist = gridSize.y - guard.position.y
			}
			if matchY && smaller {
				//left
				dist = guard.position.x
			}
			if matchY && !smaller {
				// right
				dist = gridSize.x - guard.position.x
			}
		case 1:
			hit = potential[0]
		default:
			hit = getHit(potential, smaller)
		}

		// -1 since we end right next to the obstacle
		if matchX && !done {
			dist += absdiff(hit.y, guard.position.y) - 1
		}
		if matchY && !done {
			dist += absdiff(hit.x, guard.position.x) - 1
		}

		// walk the guard marking all visited positions
		for range dist {
			guard.position.x += guard.direction.x
			guard.position.y += guard.direction.y

			visited[guard.position] = true
		}

		// rotate the guard 90° clockwise
		guard.direction = point{x: -guard.direction.y, y: guard.direction.x}

		// we left the area
		if done {
			break
		}
	}

	fmt.Println(len(visited))
}
