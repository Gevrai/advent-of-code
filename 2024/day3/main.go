package main

import (
	. "advent-of-code-2024/utils"
	"regexp"
	"strings"
)

func main() {
	DownloadDayInput(2024, 3, false)
	input := ReadInputFileRelative()

	mult := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	var res int
	for _, m := range mult.FindAllStringSubmatch(input, -1) {
		res += ParseInt(m[1]) * ParseInt(m[2])
	}
	println("Part 1:", res)

	res = 0

	var enabled bool
	for len(input) > 0 {
		var i int
		if enabled {
			i = strings.Index(input, "do()")
		} else {
			i = strings.Index(input, "don't()")
			for _, m := range mult.FindAllStringSubmatch(input[:i], -1) {
				res += ParseInt(m[1]) * ParseInt(m[2])
			}
		}
		if i < 0 {
			break
		}
		input = input[i+1:]
		enabled = !enabled
	}
	println("Part 2:", res)
}
