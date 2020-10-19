package main

import (
	. "advent-of-code-2019/utils"
	"fmt"
	"math"
	"sort"
)

func main() {
	input := ReadInputFileRelative()

	sm := NewSpaceMap(input)

	loc, max := sm.BestLocation()
	println("Part one:", max)

	twoHundreth := sm.VaporizeAsteroids(loc)[199]
	println("Part two:", twoHundreth.X*100+twoHundreth.Y)

}

type SpaceMap [][]SpaceObject

type SpaceObject rune

const (
	None     SpaceObject = '.'
	Asteroid SpaceObject = '#'
)

func NewSpaceMap(input []string) SpaceMap {
	if len(input) == 0 {
		return SpaceMap{}
	}

	sm := make(SpaceMap, len(input[0]))
	for i := range input[0] {
		sm[i] = make([]SpaceObject, len(input))
		for j := range input {
			sm[i][j] = SpaceObject(input[j][i])
		}
	}
	return sm
}

func (s SpaceMap) BestLocation() (location Point, asteroidsInSight int) {
	for j := range s[0] {
		for i := range s {
			p := Point{i, j}
			if s.Get(p) == Asteroid {
				nb := len(s.AsteroidsInLineOfSight(p))
				if nb > asteroidsInSight {
					location = Point{i, j}
					asteroidsInSight = nb
				}
			}
		}
	}
	return
}

func (s SpaceMap) VaporizeAsteroids(src Point) (vaporizedOrder []Point) {
	// Each loop performs a full rotation
	for {
		asteroids := s.AsteroidsInLineOfSight(src)
		if len(asteroids) == 0 {
			// We are done
			return vaporizedOrder
		}
		up := Vector{0, -1}
		sort.SliceStable(asteroids, func(i, j int) bool {
			// Sort by angle with UP vector, clockwise
			iAngle := up.AngleWith(src.DirectionTo(asteroids[i]))
			jAngle := up.AngleWith(src.DirectionTo(asteroids[j]))
			// if left quadrants, take bigger angle
			if asteroids[i].X < src.X {
				iAngle = 2*math.Pi - iAngle
			}
			if asteroids[j].X < src.X {
				jAngle = 2*math.Pi - jAngle
			}
			return iAngle < jAngle
		})

		for _, pos := range asteroids {
			s.Set(pos, None)
			vaporizedOrder = append(vaporizedOrder, pos)
		}
	}
}

func (s SpaceMap) AsteroidsInLineOfSight(src Point) []Point {
	collisions := map[Point]struct{}{}
	for i := range s {
		for j := range s[i] {
			col := s.RayCollides(src, src.DirectionTo(Point{i, j}))
			if col != nil {
				collisions[*col] = struct{}{}
			}
		}
	}

	l := make([]Point, 0, len(collisions))
	for p := range collisions {
		l = append(l, p)
	}
	return l
}

// Assumes ray is normalized
func (s SpaceMap) RayCollides(src Point, ray Vector) *Point {
	// Zero vector doesn't go anywhere
	if ray.X == 0 && ray.Y == 0 {
		return nil
	}
	dst := src.Add(ray)
	if !s.IsInBounds(dst) {
		return nil
	}
	switch s.Get(dst) {
	case Asteroid:
		return &dst
	case None:
		// Check next point in line of sight
		return s.RayCollides(dst, ray)
	default:
		panic(fmt.Errorf("unknown Space Object %q", string(s.Get(dst))))
	}
}

func (s SpaceMap) Get(p Point) SpaceObject {
	return s[p.X][p.Y]
}

func (s SpaceMap) Set(p Point, object SpaceObject) {
	s[p.X][p.Y] = object
}

func (s SpaceMap) IsInBounds(p Point) bool {
	return p.X >= 0 && p.X < len(s) &&
		p.Y >= 0 && p.Y < len(s[p.X])
}
