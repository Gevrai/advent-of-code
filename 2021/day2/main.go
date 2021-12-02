package main

import (
	"strings"

	. "advent-of-code-2021/utils"
)

func main() {
	DownloadDayInput(2021, 2, false)
	input := SplitNewLine(ReadInputFileRelative())

	{
		var x, y int
		for _, s := range input {
			parts := strings.Split(s, " ")
			l := ParseInt(parts[1])
			switch parts[0] {
			case "forward":
				x += l
			case "down":
				y += l
			case "up":
				y -= l
			}
		}
		println("Part 1:", x*y)
	}

	{
		var x, y int
		aim := 0
		for _, s := range input {
			parts := strings.Split(s, " ")
			l := ParseInt(parts[1])
			switch parts[0] {
			case "forward":
				x += l
				y += aim * l
			case "down":
				aim += l
			case "up":
				aim -= l
			}
		}
		println("Part 2:", x*y)
	}
}
