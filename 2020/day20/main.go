package main

import (
	. "advent-of-code-2020/utils"
)

func main() {
	DownloadDayInput(2020, 20, false)
	input := SplitNewLine(ReadInputFileRelative())

	for _, l := range input {
		println(l)
	}

	println("Part 1:")
	println("Part 2:")
}
