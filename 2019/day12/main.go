package main

import (
	. "advent-of-code-2019/utils"
	"fmt"
)

func main() {
	input := ReadInputFileRelative()

	//input = []string{
	//	"<x=-1, y=0, z=2>",
	//	"<x=2, y=-10, z=-7>",
	//	"<x=4, y=-8, z=8>",
	//	"<x=3, y=5, z=-1>",
	//}

	system, err := NewSystem(input)
	if err != nil {
		panic(err)
	}

	println("Step 0")
	system.Print()
	for i := 0; i < 1000; i++ {
		system.Update()
		println("Step ", i+1)
		system.Print()
		println(system.Energy())
	}

	println("Part one:", system.Energy())

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

func (b *Body) Update() {
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
	bodies []*Body
}

func NewSystem(input []string) (*System, error) {
	system := &System{}
	for _, s := range input {
		p, err := NewPointFromInput(s)
		if err != nil {
			return nil, err
		}
		system.bodies = append(system.bodies, &Body{pos: p})
	}
	return system, nil
}

func (s *System) Update() {
	for _, b1 := range s.bodies {
		for _, b2 := range s.bodies {
			b1.AttractFrom(*b2)
		}
	}
	for _, b := range s.bodies {
		b.Update()
	}
}

func (s *System) Energy() (totalEnergy int64) {
	for _, b := range s.bodies {
		totalEnergy += b.Energy()
	}
	return totalEnergy
}

func (s *System) Print() {
	for _, b := range s.bodies {
		println(b.String())
	}
	println()
}
