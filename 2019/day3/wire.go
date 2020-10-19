package main

import (
	. "advent-of-code-2019/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Wire struct {
	path []Point
}

func CreateWire(input string) *Wire {

	wire := &Wire{}

	var path []Point
	currentPos := Point{0, 0}

	commands := strings.Split(strings.TrimSpace(input), ",")
	for _, command := range commands {
		if command == "" {
			continue
		}

		dist, err := strconv.Atoi(command[1:])
		if err != nil {
			panic(err)
		}

		switch command[0] {
		case 'U':
			path = currentPos.LineUp(dist)
		case 'D':
			path = currentPos.LineDown(dist)
		case 'R':
			path = currentPos.LineRight(dist)
		case 'L':
			path = currentPos.LineLeft(dist)
		default:
			panic(fmt.Sprintf("unknown commmand %s", string(command[0])))
		}

		currentPos = path[len(path)-1]
		wire.path = append(wire.path, path[1:]...)
	}

	return wire
}

func (w *Wire) IntersectsWith(other *Wire) (intersections []Point) {

	pathSet := make(map[Point]struct{}, len(w.path))
	for _, p := range w.path {
		pathSet[p] = struct{}{}
	}

	for _, p := range other.path {
		if _, ok := pathSet[p]; ok {
			intersections = append(intersections, p)
		}
	}

	return intersections
}

func (w *Wire) ClosestIntersection(other *Wire) (closest *Point) {
	intersections := w.IntersectsWith(other)
	for _, p := range intersections {
		if closest == nil {
			closest = &Point{p.X, p.Y}
			continue
		}
		if p.ManhattanFromOrigin() < closest.ManhattanFromOrigin() {
			closest = &Point{p.X, p.Y}
		}
	}
	return closest
}

func (w *Wire) StepsToPoint(p Point) int {
	for i, pw := range w.path {
		if p == pw {
			return i + 1
		}
	}
	return -1
}

func (w *Wire) MinimizedSignalIntersection(other *Wire) int {

	minSteps := math.MaxInt64
	for _, intersection := range w.IntersectsWith(other) {
		steps := w.StepsToPoint(intersection) + other.StepsToPoint(intersection)
		if steps < minSteps {
			minSteps = steps
		}
	}
	return minSteps
}
