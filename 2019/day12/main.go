package main

import (
	. "advent-of-code-2019/utils"
	"fmt"
	"strings"
)

func main() {
	input := ReadInputFileRelative()

	system := NewSystem(input)

	for i := 0; i < 10; i++ {
		system.Update()
		println(system.SPrint())
	}

	println("Part one:", system.Energy())

	initialSystem := NewSystem(input)
	system = NewSystem(input)

	steps := 0
	for {
		steps++
		system.Update()
		if system.Equal(initialSystem) {
			break
		}
	}
	println("Part two:", steps)
}

type Body struct {
	pos Point3D
	vel Vector3D
}

//AttractFrom only affects objects on which method is called
func (b *Body) AttractFrom(other Body) {
	b.vel.X += compare(b.pos.X, other.pos.X)
	b.vel.Y += compare(b.pos.Y, other.pos.Y)
	b.vel.Z += compare(b.pos.Z, other.pos.Z)
}

func (b *Body) UpdatePosition() {
	b.pos.X += b.vel.X
	b.pos.Y += b.vel.Y
	b.pos.Z += b.vel.Z
}

func (b *Body) Energy() int64 {
	return b.PotentialEnergy() * b.KineticEnergy()
}

func (b *Body) PotentialEnergy() int64 {
	return AbsInt64(b.pos.X) + AbsInt64(b.pos.Y) + AbsInt64(b.pos.Z)
}

func (b *Body) KineticEnergy() int64 {
	return AbsInt64(b.vel.X) + AbsInt64(b.vel.Y) + AbsInt64(b.vel.Z)
}

func (b *Body) String() string {
	pos := fmt.Sprintf("pos=<x=%3d, y=%3d, z=%3d>", b.pos.X, b.pos.Y, b.pos.Z)
	vel := fmt.Sprintf("vel=<x=%3d, y=%3d, z=%3d>", b.vel.X, b.vel.Y, b.vel.Z)
	return pos + ", " + vel
}

type System struct {
	bodies []Body
}

func NewSystem(input []string) *System {
	system := &System{}
	system.bodies = make([]Body, len(input))
	for i, s := range input {
		p := NewPointFromInput(s)
		system.bodies[i] = Body{pos: p}
	}
	return system
}

func (s *System) Update() {
	for i := range s.bodies {
		for j := range s.bodies {
			s.bodies[i].AttractFrom(s.bodies[j])
		}
	}
	for i := range s.bodies {
		s.bodies[i].UpdatePosition()
	}
}

func (s *System) UpdateFast() {
	size := len(s.bodies)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			s.AttractBodies(i, j)
		}
		s.bodies[i].UpdatePosition()
	}
}

func (s *System) AttractBodies(i, j int) {

	if s.bodies[i].pos.X < s.bodies[j].pos.X {
		s.bodies[i].vel.X++
		s.bodies[j].vel.X--
	} else if s.bodies[i].pos.X > s.bodies[j].pos.X {
		s.bodies[i].vel.X--
		s.bodies[j].vel.X++
	}

	if s.bodies[i].pos.Y < s.bodies[j].pos.Y {
		s.bodies[i].vel.Y++
		s.bodies[j].vel.Y--
	} else if s.bodies[i].pos.Y > s.bodies[j].pos.Y {
		s.bodies[i].vel.Y--
		s.bodies[j].vel.Y++
	}

	if s.bodies[i].pos.Z < s.bodies[j].pos.Z {
		s.bodies[i].vel.Z++
		s.bodies[j].vel.Z--
	} else if s.bodies[i].pos.Z > s.bodies[j].pos.Z {
		s.bodies[i].vel.Z--
		s.bodies[j].vel.Z++
	}
}

func compare(i, j int64) int64 {
	switch {
	case i < j:
		return 1
	case i > j:
		return -1
	default:
		return 0
	}
}

func (s *System) Energy() (totalEnergy int64) {
	for _, b := range s.bodies {
		totalEnergy += b.Energy()
	}
	return totalEnergy
}

func (s *System) SPrint() string {
	sb := strings.Builder{}
	for _, b := range s.bodies {
		sb.WriteString(b.String())
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (s *System) Equal(system *System) bool {
	// We assume same size and same ordering of objects...
	for i := range s.bodies {
		if s.bodies[i] != system.bodies[i] {
			return false
		}
	}
	return true
}
