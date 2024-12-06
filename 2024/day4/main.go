package main

import (
	. "advent-of-code-2024/utils"
	"strings"
)

func main() {
	DownloadDayInput(2024, 4, false)
	input := ReadInputFileRelative()

	split := strings.Split(input, "\n")

	lines := map[int]map[int]byte{}
	for i := range split {
		lines[i] = map[int]byte{}
		for j := range split[i] {
			lines[i][j] = split[i][j]
		}
	}

	var count int
	for i := range len(lines) {
		for j := range len(lines[i]) {
			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					if di == 0 && dj == 0 {
						continue
					}
					if dir(lines, "XMAS", i, j, di, dj) {
						count++
					}
				}
			}
		}
	}
	println("Part 1:", count)

	count = 0
	for i := range len(lines) {
		for j := range len(lines[i]) {
			if (dir(lines, "MAS", i+1, j+1, -1, -1) || dir(lines, "MAS", i-1, j-1, 1, 1)) &&
				(dir(lines, "MAS", i-1, j+1, 1, -1) || dir(lines, "MAS", i+1, j-1, -1, 1)) {
				count++
			}
		}
	}
	println("Part 2:", count)
}

func dir(lines map[int]map[int]byte, search string, i, j, di, dj int) bool {
	for k := range search {
		if lines[i+(di*k)][j+(dj*k)] != search[k] {
			return false
		}
	}
	return true
}
