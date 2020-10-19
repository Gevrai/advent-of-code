package robot

import (
	"advent-of-code-2019/computer"
	. "advent-of-code-2019/utils"
	"fmt"
)

type Peeker func(Point) Color
type Painter func(Point, Color)

type Robot interface {
	Run()
	SetPeeker(Peeker)
	SetPainter(Painter)
}

type robot struct {
	cp     computer.Computer
	pos    Point
	facing Vector
	peek   Peeker
	paint  Painter
}

func NewRobot(c computer.Computer) Robot {
	return &robot{
		cp:     c,
		facing: Vector{0, 1}, // UP
	}
}

func (r *robot) SetPeeker(p Peeker) {
	r.peek = p
}

func (r *robot) SetPainter(p Painter) {
	r.paint = p
}

func (r *robot) Run() {

	r.cp.Run()

	for {
		color := r.peek(r.pos)
		if err := r.cp.Input(BigInt(int64(color))); err != nil {
			return
		}

		out := <-r.cp.Output()
		r.paint(r.pos, Color(out.Int64()))

		out = <-r.cp.Output()
		switch out.Int64() {
		case 0:
			r.rotateLeft()
		case 1:
			r.rotateRight()
		default:
			panic(fmt.Sprintf("invalid rotation input %d", out.Int64()))
		}
		r.move()
	}

}

func (r *robot) move() {
	r.pos = r.pos.Add(r.facing)
}

func (r *robot) rotateLeft() {
	r.facing.X, r.facing.Y = -r.facing.Y, r.facing.X
}

func (r *robot) rotateRight() {
	r.facing.X, r.facing.Y = r.facing.Y, -r.facing.X
}
