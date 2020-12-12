package main

import (
	"fmt"

	. "advent-of-code-2020/utils"
)

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("[%d %d]", p.x, p.y)
}

func main() {
	DownloadDayInput(2020, 12, false)
	input := SplitNewLine(ReadInputFileRelative())

	direction := map[rune]Point{
		'N': {0, 1},
		'S': {0, -1},
		'E': {1, 0},
		'W': {-1, 0},
	}

	pos := Point{0, 0}
	facing := direction['E']
	for _, l := range input {
		n := ParseInt(l[1:], 10)
		c := rune(l[0])

		if d, ok := direction[c]; ok {
			pos.y += n * d.y
			pos.x += n * d.x
		} else {
			switch c {
			case 'L':
				facing = rotate(facing, -n)
			case 'R':
				facing = rotate(facing, n)
			case 'F':
				pos.y += n * facing.y
				pos.x += n * facing.x
			default:
				panic(c)
			}
		}
	}
	println("Part 1:", Abs(pos.x)+Abs(pos.y))

	ship := Point{0, 0}
	waypoint := Point{10, 1}
	for _, l := range input {
		n := ParseInt(l[1:], 10)
		c := rune(l[0])

		if d, ok := direction[c]; ok {
			waypoint.y += n * d.y
			waypoint.x += n * d.x
		} else {
			switch c {
			case 'L':
				waypoint = rotate(waypoint, -n)
			case 'R':
				waypoint = rotate(waypoint, n)
			case 'F':
				ship.x += n * waypoint.x
				ship.y += n * waypoint.y
			default:
				panic(c)
			}
		}
	}
	println("Part 2:", Abs(ship.x)+Abs(ship.y))
}

func rotate(p Point, degrees int) Point {
	switch Mod(degrees, 360) {
	case 0:
		return p
	case 90:
		return Point{p.y, -p.x}
	case 180:
		return Point{-p.x, -p.y}
	case 270:
		return Point{-p.y, p.x}
	default:
		panic(degrees)
	}

}
