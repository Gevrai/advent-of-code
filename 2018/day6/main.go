package main

import (
	"strings"

	"advent-of-code-2018/utils"
)

func main() {
	utils.DownloadDayInput(2018, 6, false)
	input := utils.ReadInputFileRelativeSplitNewline()

	points := ParsePoints(input)

	println("Part 1:", points)

	println("Part 2:")
}

type Point struct {
	id   int
	x, y int
}

func ParsePoints(input []string) (points []Point) {
	id := 0
	for _, l := range input {
		parts := strings.Split(l, ",")
		points = append(points, Point{
			id: id,
			x:  utils.ParseInt(strings.TrimSpace(parts[0]), 10),
			y:  utils.ParseInt(strings.TrimSpace(parts[1]), 10),
		})
		id++
	}
	return points
}
