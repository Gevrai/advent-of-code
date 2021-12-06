package main

import (
	"strings"

	. "advent-of-code-2021/utils"
)

type Point struct{ x, y int }

type Line struct{ start, end Point }

type Field struct {
	covered map[Point]int
}

func main() {
	DownloadDayInput(2021, 5, false)
	input := SplitNewLine(ReadInputFileRelative())

	var lines []Line
	for _, l := range input {
		parts := strings.Split(l, " -> ")
		lines = append(lines, Line{
			start: NewPoint(parts[0]),
			end:   NewPoint(parts[1]),
		})
	}

	field := Field{map[Point]int{}}

	for _, l := range lines {
		if l.start.x == l.end.x || l.start.y == l.end.y {
			field.cover(l)
		}
	}
	count := 0
	for _, n := range field.covered {
		if n > 1 {
			count++
		}
	}
	println("Part 1:", count)

	field.covered = map[Point]int{}
	for _, l := range lines {
		field.cover(l)
	}
	count = 0
	for _, n := range field.covered {
		if n > 1 {
			count++
		}
	}
	println("Part 2:", count)
}

func NewPoint(s string) Point {
	p := strings.Split(s, ",")
	AssertEqual(len(p), 2)
	return Point{
		x: ParseInt(p[0]),
		y: ParseInt(p[1]),
	}
}

func (f *Field) cover(l Line) {
	// Horizontal
	if l.start.x == l.end.x {
		start, end := l.start.y, l.end.y
		if start > end {
			start, end = end, start
		}
		for i := start; i <= end; i++ {
			f.covered[Point{l.start.x, i}]++
		}
		return
	}
	// Vertical
	if l.start.y == l.end.y {
		start, end := l.start.x, l.end.x
		if start > end {
			start, end = end, start
		}
		for i := start; i <= end; i++ {
			f.covered[Point{i, l.start.y}]++
		}
		return
	}
	s, e := l.start, l.end
	if s.x > e.x {
		s, e = e, s
	}
	// Diagonal ascending
	if s.y > e.y {
		for p := s; p.x <= e.x; {
			f.covered[p]++
			p.x++
			p.y--
		}
		return
	}
	// Diagonal descending
	for p := s; p.x <= e.x; {
		f.covered[p]++
		p.x++
		p.y++
	}
}
