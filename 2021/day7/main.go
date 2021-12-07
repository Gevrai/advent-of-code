package main

import (
	"strings"

	. "advent-of-code-2021/utils"
)

func main() {
	DownloadDayInput(2021, 7, false)
	input := ReadInputFileRelative()

	var k []int
	for _, l := range strings.Split(input, ",") {
		k = append(k, ParseInt(l))
	}

	pos := make([]int, Max(k...))
	for i := range pos {
		for j := range k {
			pos[i] += Abs(k[j] - i)
		}
	}
	println("Part 1:", Min(pos...))

	pos = make([]int, Max(k...))
	for i := range pos {
		for j := range k {
			pos[i] += cost(Abs(k[j] - i))
		}
	}
	println("Part 2:", Min(pos...))
}

func cost(i int) int {
	if i == 1 || i == 0 {
		return i
	}
	return i + cost(i-1)
}
