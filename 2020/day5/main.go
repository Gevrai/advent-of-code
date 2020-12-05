package main

import (
	"advent-of-code-2020/utils"
)

func main() {
	utils.DownloadDayInput(2020, 5, false)
	input := utils.ReadInputFileRelativeSplitNewline()

	max := 0
	for _, l := range input {
		p := seatPos(l)
		if p > max {
			max = p
		}
	}
	println("Part 1:", max)

	ids := make([]bool, 128*8)
	for _, l := range input {
		ids[seatPos(l)] = true
	}
	for i, id := range ids {
		if i == 0 {
			continue
		}
		if !id && ids[i-1] && ids[i+1] {
			println("Part 2:", i)
		}
	}
}

func seatPos(s string) int {
	var row, col int

	t := 127
	for _, c := range s[:7] {
		p := (row + t) / 2
		if c == 'F' {
			t = p
		} else if c == 'B' {
			row = p + 1
		} else {
			panic(string(c))
		}
	}

	t = 7
	for _, c := range s[7:] {
		p := (col + t) / 2
		if c == 'L' {
			t = p
		} else if c == 'R' {
			col = p + 1
		} else {
			panic(string(c))
		}
	}
	return 8*row + col
}
