package main

import (
	. "advent-of-code-2024/utils"
)

func main() {
	DownloadDayInput(2024, 8, false)
	input := SplitNewLine(ReadInputFileRelative())

	antennas := make([][]byte, len(input))
	antinodes := make([][]int, len(input))
	antennasLoc := make(map[string][]vec2)

	for j := range input {
		antennas[j] = []byte(input[j])
		antinodes[j] = make([]int, len(input[j]))
		for i := range input[j] {
			if isAlphanum(input[j][i]) {
				antennasLoc[string(input[j][i])] = append(antennasLoc[string(input[j][i])], vec2{i, j})
			}
		}
	}

	for _, locs := range antennasLoc {
		for n, nloc := range locs {
			for _, mloc := range locs[n+1:] {
				dir := sub(mloc, nloc)
				an1 := add(mloc, dir)
				an2 := sub(nloc, dir)

				if inBounds(antinodes, an1) {
					antinodes[an1.y][an1.x]++
				}
				if inBounds(antinodes, an2) {
					antinodes[an2.y][an2.x]++
				}
			}
		}
	}

	var part1 int
	for j := range antinodes {
		for i := range antinodes[j] {
			if antinodes[j][i] > 0 {
				part1++
			}
		}
	}
	println("Part 1:", part1)

	// no need to cleanup antinodes, they'll be the same but with more
	for _, locs := range antennasLoc {
		for n, nloc := range locs {
			for _, mloc := range locs[n+1:] {
				dir := sub(mloc, nloc)
				an1 := mloc
				for inBounds(antinodes, an1) {
					antinodes[an1.y][an1.x]++
					an1 = add(an1, dir)
				}

				an2 := nloc
				for inBounds(antinodes, an2) {
					antinodes[an2.y][an2.x]++
					an2 = sub(an2, dir)
				}
			}
		}
	}

	var part2 int
	for j := range antinodes {
		for i := range antinodes[j] {
			if antinodes[j][i] > 0 {
				part2++
			}
		}
	}
	println("Part 2:", part2)
}

func isAlphanum(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		(b >= '0' && b <= '9')
}

func inBounds[T any](list [][]T, point vec2) bool {
	out := point.y < 0 || point.x < 0 || point.y > len(list)-1 || point.x > len(list[point.y])-1
	return !out
}

type vec2 struct{ x, y int }

func sub(a, b vec2) vec2 {
	return vec2{a.x - b.x, a.y - b.y}
}

func add(a, b vec2) vec2 {
	return vec2{a.x + b.x, a.y + b.y}
}
