package main

import (
	"advent-of-code-2020/utils"
)

func main() {
	utils.DownloadDayInput(2020, 3, false)
	input := utils.ReadInputFileRelative()

	println("Part 1:", skiDown(input, 3, 1))

	println("Part 2:", skiDown(input, 1, 1)*
		skiDown(input, 3, 1)*
		skiDown(input, 5, 1)*
		skiDown(input, 7, 1)*
		skiDown(input, 1, 2),
	)
}

func skiDown(input []string, x, y int) (trees int) {
	var i, j int
	for {
		i = (i + x) % len(input[0])
		j += y

		if j >= len(input) {
			return trees
		}
		if input[j][i] == '#' {
			trees++
		}
	}
}
