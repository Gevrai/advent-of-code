package main

import (
	. "advent-of-code-2021/utils"
)

func main() {
	DownloadDayInput(2021, 1, false)
	input := SplitNewLine(ReadInputFileRelative())

	measurements := make([]int, len(input))
	for i, s := range input {
		measurements[i] = ParseInt(s)
	}

	count := 0
	for i := 1; i < len(measurements); i++ {
		if measurements[i-1] < measurements[i] {
			count++
		}
	}
	println("Part 1:", count)

	count = 0
	for i := 1; i < len(measurements); i++ {
		prev := calcWindow(measurements, i-1, i+2)
		curr := calcWindow(measurements, i, i+3)
		println(i, curr, prev < curr)
		if prev < curr {
			count++
		}
	}
	println("Part 2:", count)
}

func calcWindow(list []int, start, end int) (count int) {
	start = Max(start, 0)
	end = Min(end, len(list))
	for i := start; i < end; i++ {
		count += list[i]
	}
	return count
}
