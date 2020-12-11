package main

import (
	"math"
	"strings"

	"advent-of-code-2018/utils"
)

func main() {
	utils.DownloadDayInput(2018, 5, false)
	input := utils.ReadInputFileRelativeSplitNewline()[0]

	var reactions []string
	for c := 'a'; c <= 'z'; c++ {
		s := string(c) + strings.ToUpper(string(c))
		reactions = append(reactions, s)
		s = strings.ToUpper(string(c)) + string(c)
		reactions = append(reactions, s)
	}

	l := 0
	for l != len(input) {
		l = len(input)
		for _, r := range reactions {
			input = strings.ReplaceAll(input, r, "")
		}
	}

	println("Part 1:", len(input))

	input = utils.ReadInputFileRelativeSplitNewline()[0]
	shortest := math.MaxInt32
	for c := 'a'; c <= 'z'; c++ {
		test := strings.ReplaceAll(input, string(c), "")
		test = strings.ReplaceAll(test, strings.ToUpper(string(c)), "")

		l := 0
		for l != len(test) {
			l = len(test)
			for _, r := range reactions {
				test = strings.ReplaceAll(test, r, "")
			}
		}
		if l < shortest {
			shortest = l
		}
	}

	println("Part 2:", shortest)
}
