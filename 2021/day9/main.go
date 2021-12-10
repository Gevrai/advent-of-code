package main

import (
	"sort"

	. "advent-of-code-2021/utils"
)

const example = `2199943210
3987894921
9856789892
8767896789
9899965678
`

type point struct{ x, y int }

func main() {
	DownloadDayInput(2021, 9, false)
	input := SplitNewLine(ReadInputFileRelative())
	//input = SplitNewLine(example)

	count := 0
	m := map[point]int{}
	low := map[point]int{}

	for j, l := range input {
		for i, c := range l {
			m[point{i, j}] = ParseInt(string(c))
		}
	}

outer:
	for p, h := range m {
		for _, d := range directions(p) {
			if o, ok := m[d]; ok && o <= h {
				continue outer
			}
		}
		count += 1 + h
		low[p] = h
	}

	println("Part 1:", count)

	var sizes []int
	for l := range low {
		bassin := map[point]bool{}
		flood(m, bassin, l)
		sizes = append(sizes, len(bassin))
	}

	sort.Ints(sizes)
	println("Part 2:", Mult(sizes[len(sizes)-3:]))
}

func flood(m map[point]int, bassin map[point]bool, p point) int {
	bassin[p] = true
	count := 1
	for _, pot := range directions(p) {
		if o, ok := m[pot]; ok && o != 9 && !bassin[pot] {
			count += flood(m, bassin, pot)
		}
	}
	return count
}

func directions(p point) []point {
	return []point{
		{p.x + 1, p.y},
		{p.x - 1, p.y},
		{p.x, p.y + 1},
		{p.x, p.y - 1},
	}
}
