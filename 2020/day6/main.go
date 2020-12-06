package main

import (
	"strings"

	. "advent-of-code-2020/utils"
)

func main() {
	DownloadDayInput(2020, 6, false)
	input := ReadInputFileRelative()

	groups := strings.Split(input, "\n\n")
	total := 0
	for _, l := range groups {
		for i := 'a'; i <= 'z'; i++ {
			if strings.Count(l, string(i)) > 0 {
				total++
			}
		}
	}
	println("Part 1:", total)

	total = 0
	newGroup := true
	candidates := ""
	for _, l := range append(SplitNewLine(input), "") {
		if newGroup {
			newGroup = false
			candidates = l
			continue
		}
		if l == "" {
			total += len(candidates)
			newGroup = true
			continue
		}
		for _, c := range candidates {
			if strings.Count(l, string(c)) == 0 {
				candidates = strings.ReplaceAll(candidates, string(c), "")
			}
		}
	}
	println("Part 2:", total)

}
