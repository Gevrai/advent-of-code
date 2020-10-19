package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	. "advent-of-code-2019/utils"
)

func TestCreateWire(t *testing.T) {
	tests := []struct {
		name   string
		inputs string
		want   []Point
	}{
		{"up", "U2", []Point{{0, 1}, {0, 2}}},
		{"down", "D2", []Point{{0, -1}, {0, -2}}},
		{"right", "R2", []Point{{1, 0}, {2, 0}}},
		{"left", "L2", []Point{{-1, 0}, {-2, 0}}},

		{"left up", "L2,U2", []Point{
			{-1, 0}, {-2, 0},
			{-2, 1}, {-2, 2},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wire := CreateWire(tt.inputs)
			require.Equal(t, tt.want, wire.path)
		})
	}
}

func TestWire_IntersectsWith(t *testing.T) {
	tests := []struct {
		name              string
		wire1             string
		wire2             string
		wantIntersections []Point
	}{
		{
			name:              "empty wires",
			wire1:             "",
			wire2:             "",
			wantIntersections: nil,
		},
		{
			name:              "no intersects",
			wire1:             "R8,U5,L5,D3",
			wire2:             "D1,R9,U1",
			wantIntersections: nil,
		},
		{
			name:              "one intersect",
			wire1:             "R8,U5,L5,D3",
			wire2:             "U7,R6,D4",
			wantIntersections: []Point{{6, 5}},
		},
		{
			name:              "two intersects",
			wire1:             "R8,U5,L5,D3",
			wire2:             "U7,R6,D4,L4",
			wantIntersections: []Point{{6, 5}, {3, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wire1 := CreateWire(tt.wire1)
			wire2 := CreateWire(tt.wire2)
			intersections := wire1.IntersectsWith(wire2)
			require.Equal(t, tt.wantIntersections, intersections)
		})
	}
}

func TestWire_ClosestIntersection(t *testing.T) {
	tests := []struct {
		name                 string
		wire1                string
		wire2                string
		wantClosestManhattan int
	}{
		{
			name:                 "example 1",
			wire1:                "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			wire2:                "U62,R66,U55,R34,D71,R55,D58,R83",
			wantClosestManhattan: 159,
		},
		{
			name:                 "example 2",
			wire1:                "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			wire2:                "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			wantClosestManhattan: 135,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wire1 := CreateWire(tt.wire1)
			wire2 := CreateWire(tt.wire2)
			closest := wire1.ClosestIntersection(wire2)
			require.Equal(t, tt.wantClosestManhattan, closest.ManhattanFromOrigin())
		})
	}
}

func TestPoint_ManhattanFromOrigin(t *testing.T) {
	tests := []struct {
		x, y int
		want int
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
		{-1, 0, 1},
		{0, -1, 1},
		{1, -1, 2},
		{-1, -1, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d, %d", tt.x, tt.y), func(t *testing.T) {
			p := &Point{tt.x, tt.y}
			require.Equal(t, tt.want, p.ManhattanFromOrigin())
		})
	}
}

func TestWire_MinimizedSignalIntersection(t *testing.T) {
	tests := []struct {
		name       string
		wire1      string
		wire2      string
		wantSignal int
	}{
		{
			name:       "part 1 example",
			wire1:      "R8,U5,L5,D3",
			wire2:      "U7,R6,D4,L4",
			wantSignal: 30,
		},
		{
			name:       "example 1",
			wire1:      "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			wire2:      "U62,R66,U55,R34,D71,R55,D58,R83",
			wantSignal: 610,
		},
		{
			name:       "example 2",
			wire1:      "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			wire2:      "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			wantSignal: 410,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wire1 := CreateWire(tt.wire1)
			wire2 := CreateWire(tt.wire2)
			signal := wire1.MinimizedSignalIntersection(wire2)
			require.Equal(t, signal, wire2.MinimizedSignalIntersection(wire1))
			require.Equal(t, tt.wantSignal, signal)
		})
	}
}
