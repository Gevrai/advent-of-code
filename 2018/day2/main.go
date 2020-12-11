package main

import (
	"strings"

	"advent-of-code-2018/utils"
)

func main() {
	utils.DownloadDayInput(2018, 2, false)
	input := utils.ReadInputFileRelativeSplitNewline()

	repeat := func(i int) func(string) bool {
		return func(s string) bool {
			for c := 'a'; c <= 'z'; c++ {
				if strings.Count(s, string(c)) == i {
					return true
				}
			}
			return false
		}
	}
	println("Part 1:", utils.Count(input, repeat(2))*utils.Count(input, repeat(3)))

	for i := range input {
		for j := range input {
			if i == j {
				continue
			}
			differences := 0
			index := 0
			for k := range input[i] {
				if input[i][k] != input[j][k] {
					differences++
					index = k
				}
			}
			if differences == 1 {
				println("Part 2:", input[i][:index]+input[i][index+1:])
			}
		}

	}
}
