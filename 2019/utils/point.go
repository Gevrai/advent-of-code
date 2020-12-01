package utils

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) Add(d Vector) Point {
	return Point{p.X + d.X, p.Y + d.Y}
}

func (p Point) DirectionTo(other Point) Vector {
	d := Vector{
		X: other.X - p.X,
		Y: other.Y - p.Y,
	}
	if err := d.Normalize(); err != nil {
		return Vector{0, 0}
	}
	return d
}

func (p *Point) ManhattanFromOrigin() int {
	x := p.X
	if x < 0 {
		x = -x
	}
	y := p.Y
	if y < 0 {
		y = -y
	}
	return x + y
}

func (p Point) LineRight(dist int) (path []Point) {
	for i := p.X; i <= p.X+dist; i++ {
		path = append(path, Point{i, p.Y})
	}
	return path
}

func (p Point) LineLeft(dist int) (path []Point) {
	for i := p.X; i >= p.X-dist; i-- {
		path = append(path, Point{i, p.Y})
	}
	return path
}

func (p Point) LineUp(dist int) (path []Point) {
	for j := p.Y; j <= p.Y+dist; j++ {
		path = append(path, Point{p.X, j})
	}
	return path
}

func (p Point) LineDown(dist int) (path []Point) {
	for j := p.Y; j >= p.Y-dist; j-- {
		path = append(path, Point{p.X, j})
	}
	return path
}

func (p Point) String() string {
	return fmt.Sprintf("<x=%3d, y=%3d>", p.X, p.Y)
}
