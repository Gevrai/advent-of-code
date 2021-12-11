package main

import (
	. "advent-of-code-2021/utils"
)

const example = `
5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
`

type point struct{ x, y int }

func main() {
	DownloadDayInput(2021, 11, false)
	input := SplitNewLine(ReadInputFileRelative())
	//input = SplitNewLine(example[1:])

	count := 0
	m := map[point]int{}
	for j, l := range input {
		for i, c := range l {
			m[point{i, j}] = ParseInt(string(c))
		}
	}

	for i := 0; i < 100; i++ {
		f := map[point]bool{}
		for p := range m {
			inc(m, f, p)
		}
		for p := range m {
			if m[p] > 9 {
				m[p] = 0
				count++
			}
		}
	}
	println("Part 1:", count)

	m = map[point]int{}
	for j, l := range input {
		for i, c := range l {
			m[point{i, j}] = ParseInt(string(c))
		}
	}
	i := 0
	for {
		i++
		f := map[point]bool{}
		for p := range m {
			inc(m, f, p)
		}
		for p := range m {
			if m[p] > 9 {
				m[p] = 0
				count++
			}
		}
		if len(m) == len(f) {
			println("Part 2:", i)
			return
		}
	}
}

func inc(m map[point]int, f map[point]bool, p point) {
	if _, ok := m[p]; !ok {
		return
	}
	m[p]++
	if m[p] > 9 && !f[p] {
		f[p] = true
		for _, d := range directions(p) {
			inc(m, f, d)
		}
	}
}

func directions(p point) []point {
	return []point{
		{p.x + 1, p.y},
		{p.x - 1, p.y},
		{p.x, p.y + 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y + 1},
		{p.x - 1, p.y - 1},
		{p.x + 1, p.y - 1},
		{p.x - 1, p.y + 1},
	}
}
