package main

import (
	"fmt"
	"math"

	. "advent-of-code-2021/utils"
)

const example = `
1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
`

type point struct{ x, y int }

func main() {
	DownloadDayInput(2021, 15, false)
	input := SplitNewLine(ReadInputFileRelative())
	//input = SplitNewLine(example[1:])

	m := map[point]int{}
	for j, l := range input {
		for i, c := range l {
			m[point{i, j}] = ParseInt(string(c))
		}
	}
	start, end := bounds(m)
	//println("Part 1:", shortestPathCostBruteForce(m, start, end))
	println("Part 1:", shortestPathCost(m, start, end))

	fullmap := expandMap(m, 5)
	start, end = bounds(fullmap)
	println("Part 2:", shortestPathCost(fullmap, start, end))
}

func expandMap(m map[point]int, size int) map[point]int {
	fullmap := make(map[point]int, len(m)*size*size)
	_, max := bounds(m)

	for p, c := range m {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				P := point{
					x: i*(max.x+1) + p.x,
					y: j*(max.y+1) + p.y,
				}
				fullmap[P] = ((c + i + j - 1) % 9) + 1
			}
		}
	}
	return fullmap
}

// pretty much dijkstra
func shortestPathCost(graph map[point]int, start, end point) int {
	shortestPath := map[point]int{}
	unvisited := map[point]bool{}
	candidates := map[point]bool{}
	for p := range graph {
		shortestPath[p] = math.MaxInt64
		unvisited[p] = true
	}
	shortestPath[start] = 0
	delete(unvisited, start)

	curr := start
	for unvisited[end] {
		forDirections(curr, func(d point) {
			if c, ok := graph[d]; ok && shortestPath[d] > shortestPath[curr]+c {
				shortestPath[d] = shortestPath[curr] + c
				if unvisited[d] {
					candidates[d] = true // new candidate since it now has a value
				}
			}
		})
		delete(unvisited, curr)

		// Find next point to visit in candidates (lowest cost)
		min := math.MaxInt64
		for p := range candidates {
			if shortestPath[p] < min {
				min = shortestPath[p]
				curr = p
			}
		}
		delete(candidates, curr)
	}
	return shortestPath[end]
}

func printMap(m map[point]int) {
	min, max := bounds(m)

	for j := min.y; j <= max.y; j++ {
		for i := min.x; i <= max.x; i++ {
			fmt.Printf("%d", m[point{i, j}])
		}
		println()
	}
	println()
}

// Very very inefficient method kept for posterity...
func shortestPathCostBruteForce(m map[point]int, start point, end point) int {
	costs := map[point]int{}
	for p := range m {
		costs[p] = math.MaxUint32
	}
	costs[end] = m[end]

	var fillUpCosts func(current point)
	fillUpCosts = func(current point) {
		forDirections(current, func(d point) {
			if ceil, ok := m[d]; ok {
				if costs[d] > costs[current]+ceil {
					costs[d] = costs[current] + ceil
					fillUpCosts(d)
				}
			}
		})
	}
	fillUpCosts(end)
	return costs[start] - m[start]
}

func bounds(m map[point]int) (min, max point) {
	min.x = math.MaxInt64
	min.y = math.MaxInt64
	max.x = math.MinInt64
	max.y = math.MinInt64
	for p := range m {
		min.x = Min(min.x, p.x)
		min.y = Min(min.y, p.y)
		max.x = Max(max.x, p.x)
		max.y = Max(max.y, p.y)
	}
	return min, max
}

func forDirections(p point, fn func(p point)) {
	fn(point{p.x + 1, p.y})
	fn(point{p.x - 1, p.y})
	fn(point{p.x, p.y + 1})
	fn(point{p.x, p.y - 1})
}
