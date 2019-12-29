package main

import (
	"testing"
)

func TestCross(t *testing.T) {
	var tests = []struct {
		w1, w2   string
		distance int // Part 1
		length   int // Part 2
	}{
		{
			"R8,U5,L5,D3",
			"U7,R6,D4,L4",
			6,
			30,
		},

		{
			"R75,D30,R83,U83,L12,D49,R71,U7,L72",
			"U62,R66,U55,R34,D71,R55,D58,R83",
			159,
			610,
		},

		{
			"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			135,
			410,
		},
	}

	for _, tt := range tests {
		wire1, wire2 := parseWire(tt.w1), parseWire(tt.w2)
		distance := intersectClosest(wire1, wire2)
		if tt.distance != distance {
			t.Errorf("intersectClosest(%q %q) => Got %v, want %v", tt.w1, tt.w2, distance, tt.distance)
		}

		length := intersectShortest(wire1, wire2)
		if tt.length != length {
			t.Errorf("intersectShortest(%q %q) => Got %v, want %v", tt.w1, tt.w2, length, tt.length)
		}
	}
}
