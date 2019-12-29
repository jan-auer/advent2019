package main

import (
	"fmt"
	"strconv"
	"strings"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type Point struct {
	x, y int
}

// Returns the manhattan distance of the point to the origin {0, 0}.
func (p Point) ManhattanDistance() int {
	return abs(p.x) + abs(p.y)
}

// Maps points along the wire to the minimum wire length it takes to reach that
// point from the origin {0, 0}. If this wire crosses multiple times through a
// point, only the minimum length is stored.
type Wire map[Point]int

func readWord() string {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		panic(err)
	}
	return str
}

func parseWire(str string) Wire {
	point := Point{0, 0}
	points := make(Wire, 0)
	totalLength := 0

	for _, instr := range strings.Split(str, ",") {
		dx, dy := 0, 0
		switch instr[0] {
		case 'U':
			dy = -1
		case 'D':
			dy = 1
		case 'L':
			dx = -1
		case 'R':
			dx = 1
		}

		length, err := strconv.Atoi(instr[1:])
		if err != nil {
			panic(err)
		}

		for ; length > 0; length-- {
			point.x += dx
			point.y += dy
			totalLength++

			// Keep the minimum totalLength
			if _, set := points[point]; !set {
				points[point] = totalLength
			}
		}
	}

	return points
}

// Intersects the two wires and returns the minimum value determined by the
// weight function.
func intersectMinBy(w1, w2 Wire, f func(Point) int) int {
	min := -1

	for point := range w1 {
		_, contains := w2[point]
		if !contains {
			continue
		}

		weight := f(point)
		if min < 0 || weight < min {
			min = weight
		}
	}

	return min
}

// Returns the closest intersection to the Origin {0, 0} by manhattan distance.
func intersectClosest(w1, w2 Wire) int {
	return intersectMinBy(w1, w2, Point.ManhattanDistance)
}

// Returns the intersection with the shortest combined wire length.
func intersectShortest(w1, w2 Wire) int {
	combinedLength := func(p Point) int {
		return w1[p] + w2[p]
	}

	return intersectMinBy(w1, w2, combinedLength)
}

func main() {
	wire1 := parseWire(readWord())
	wire2 := parseWire(readWord())

	distance := intersectClosest(wire1, wire2)
	fmt.Println("Part 1:", distance)

	length := intersectShortest(wire1, wire2)
	fmt.Println("Part 2:", length)
}
