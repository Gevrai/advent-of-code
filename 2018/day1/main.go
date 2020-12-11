package main

import (
	"strconv"

	"advent-of-code-2018/utils"
)

func main() {
	utils.DownloadDayInput(2018, 1, false)
	input := utils.ReadInputFileRelativeSplitNewline()

	freq := 0
	for _, f := range input {
		i, _ := strconv.Atoi(f[1:])
		if f[0] == '+' {
			freq += i
		} else {
			freq -= i
		}
	}

	println("Part 1:", freq)

	frequencies := map[int]struct{}{}
	freq = 0
	for {
		for _, f := range input {
			i, _ := strconv.Atoi(f[1:])
			if f[0] == '+' {
				freq += i
			} else {
				freq -= i
			}
			if _, ok := frequencies[freq]; ok {
				println("Part 2:", freq)
				return
			} else {
				frequencies[freq] = struct{}{}
			}
		}
	}
}
