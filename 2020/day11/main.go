package main

import (
	. "advent-of-code-2020/utils"
)

func main() {
	DownloadDayInput(2020, 11, false)
	input := SplitNewLine(ReadInputFileRelative())

	plan := make([][]byte, len(input))
	for i := range input {
		plan[i] = []byte(input[i])
	}
	for {
		newPlan := simulate(plan)
		if equal(newPlan, plan) {
			break
		}
		plan = newPlan
	}

	println("Part 1:", count(plan))

	plan = make([][]byte, len(input))
	for i := range input {
		plan[i] = []byte(input[i])
	}
	for {
		newPlan := simulate2(plan)
		if equal(newPlan, plan) {
			break
		}
		plan = newPlan
	}
	println("Part 2:", count(plan))
}

func printPlan(a [][]byte) {
	for _, i := range a {
		println(string(i))
	}
	println()
}

func equal(a, b [][]byte) bool {
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func count(a [][]byte) int {
	count := 0
	for i := range a {
		for j := range a[i] {
			if a[i][j] == '#' {
				count++
			}
		}
	}
	return count
}

func simulate(plan [][]byte) [][]byte {
	newPlan := make([][]byte, len(plan))
	for i := range plan {
		newPlan[i] = make([]byte, len(plan[i]))
		copy(newPlan[i], plan[i])
		for j := range plan[i] {
			switch plan[i][j] {
			case '.':
				continue
			case '#':
				if countAdjacentOccupied(plan, i, j) >= 4 {
					newPlan[i][j] = 'L'
				}
			case 'L':
				if countAdjacentOccupied(plan, i, j) == 0 {
					newPlan[i][j] = '#'
				}
			}
		}
	}
	return newPlan
}

func simulate2(plan [][]byte) [][]byte {
	newPlan := make([][]byte, len(plan))
	for i := range plan {
		newPlan[i] = make([]byte, len(plan[i]))
		copy(newPlan[i], plan[i])
		for j := range plan[i] {
			switch plan[i][j] {
			case '.':
				continue
			case '#':
				if countVisibleOccupied(plan, i, j) >= 5 {
					newPlan[i][j] = 'L'
				}
			case 'L':
				if countVisibleOccupied(plan, i, j) == 0 {
					newPlan[i][j] = '#'
				}
			}
		}
	}
	return newPlan
}

var around = []struct{ i, j int }{
	{-1, -1}, {-1, 0}, {-1, +1},
	{0, -1}, {0, +1},
	{+1, -1}, {+1, 0}, {+1, +1},
}

func countAdjacentOccupied(plan [][]byte, x, y int) (count int) {
	for _, a := range around {
		if isOccupied(plan, x+a.i, y+a.j) {
			count++
		}
	}
	return count
}

func countVisibleOccupied(plan [][]byte, x, y int) (count int) {
	for _, a := range around {
		if isVisibleOccupied(plan, x, y, a.i, a.j) {
			count++
		}
	}
	return count
}

func isOccupied(plan [][]byte, i, j int) bool {
	return !outOfBounds(plan, i, j) && plan[i][j] == '#'
}

func isVisibleOccupied(plan [][]byte, x, y, i, j int) bool {
	for {
		x += i
		y += j
		if outOfBounds(plan, x, y) {
			return false
		}
		if plan[x][y] == '#' {
			return true
		}
		if plan[x][y] == 'L' {
			return false
		}
	}
}

func outOfBounds(plan [][]byte, i, j int) bool {
	return i < 0 || i >= len(plan) || j < 0 || j >= len(plan[i])
}
