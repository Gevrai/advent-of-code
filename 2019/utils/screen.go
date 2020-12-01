package utils

import (
	"fmt"
	"strings"
	"sync"
)

type Tile int64

const (
	Empty Tile = iota
	Wall
	Block
	HorizontalPaddle
	Ball
)

func (c Tile) String() string {
	switch c {
	case Empty:
		return " "
	case Wall:
		return "█"
	case Block:
		return "■"
	case HorizontalPaddle:
		return "═"
	case Ball:
		return "o"
	default:
		panic(fmt.Sprintf("unknown color %d", c))
	}
}

type Screen interface {
	Peek(Point) Tile
	Paint(Point, Tile)
	Show() string
	FindTiles(tile Tile) []Point
}

type MapScreen struct {
	m *sync.Map
}

func NewMapScreen() MapScreen {
	return MapScreen{&sync.Map{}}
}

func (m MapScreen) Peek(point Point) Tile {
	tile, ok := m.m.Load(point)
	if !ok {
		return Empty
	}
	return tile.(Tile)
}

func (m MapScreen) Paint(point Point, tile Tile) {
	m.m.Store(point, tile)
}

func (m MapScreen) Show() string {
	var top, bottom, right, left int
	m.m.Range(func(key, value interface{}) bool {
		p := key.(Point)
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
		return true
	})

	sb := strings.Builder{}
	for j := top; j >= bottom; j-- {
		for i := left; i <= right; i++ {
			sb.WriteString(m.Peek(Point{i, j}).String())
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (m MapScreen) FindTiles(tile Tile) (tilesPositions []Point) {
	m.m.Range(func(key, value interface{}) bool {
		if value.(Tile) == tile {
			tilesPositions = append(tilesPositions, key.(Point))
		}
		return true
	})
	return
}
