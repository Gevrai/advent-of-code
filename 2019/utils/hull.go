package utils

import (
	"fmt"
	"strings"
)

type Color int64

const (
	BLACK Color = iota
	WHITE
)

func (c Color) String() string {
	switch c {
	case BLACK:
		return " "
	case WHITE:
		return "█"
	default:
		panic(fmt.Sprintf("unknown color %d", c))
	}
}

type Hull interface {
	Peek(Point) Color
	Paint(Point, Color)
	Show() string
}

type MapHull map[Point]Color

func NewMapHull() Hull {
	return MapHull(make(map[Point]Color))
}

func (m MapHull) Peek(point Point) Color {
	color, ok := m[point]
	if !ok {
		return BLACK
	}
	return color
}

func (m MapHull) Paint(point Point, color Color) {
	m[point] = color
}

func (m MapHull) Show() string {
	var top, bottom, right, left int
	for p := range m {
		if p.X < left {
			left = p.X
		}
		if p.X > right {
			right = p.X
		}
		if p.Y < bottom {
			bottom = p.Y
		}
		if p.Y > top {
			top = p.Y
		}
	}

	sb := strings.Builder{}
	for j := top; j >= bottom; j-- {
		for i := left; i <= right; i++ {
			sb.WriteString(m.Peek(Point{i, j}).String())
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
