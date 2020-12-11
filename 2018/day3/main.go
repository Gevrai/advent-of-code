package main

import (
	"strings"

	"advent-of-code-2018/utils"
)

func main() {
	utils.DownloadDayInput(2018, 3, false)
	input := utils.ReadInputFileRelativeSplitNewline()

	fabric := make(map[Point]int)

	count := 0
	for _, l := range input {
		sq := ParseSquare(l)
		for i := sq.x; i < sq.x+sq.w; i++ {
			for j := sq.y; j < sq.y+sq.h; j++ {
				fabric[Point{i, j}] = fabric[Point{i, j}] + 1
				if fabric[Point{i, j}] == 2 {
					count++
				}
			}
		}
	}

	println("Part 1:", count)

	for _, l := range input {
		overlaps := false
		sq := ParseSquare(l)
		for i := sq.x; i < sq.x+sq.w; i++ {
			for j := sq.y; j < sq.y+sq.h; j++ {
				if fabric[Point{i, j}] > 1 {
					overlaps = true
				}
			}
		}
		if !overlaps {
			println("Part 2:", sq.id)
		}
	}

}

type Point struct {
	x, y int
}

func grow(fabric [][]int, w, h int) [][]int {
	if len(fabric) < h {
		fabric = append(fabric, make([][]int, h-len(fabric))...)
	}

	for i := range fabric {
		if len(fabric[i]) < h {
			fabric[i] = append(fabric[i], make([]int, h-len(fabric[i]))...)
		}
	}
	return fabric
}

type Square struct {
	id, x, y, w, h int
}

func ParseSquare(input string) Square {
	parts := strings.Split(input, " ")
	pos := strings.Split(strings.Trim(parts[2], ":"), ",")
	size := strings.Split(parts[3], "x")
	return Square{
		id: utils.ParseInt(strings.Trim(parts[0], "#"), 10),
		x:  utils.ParseInt(pos[0], 10),
		y:  utils.ParseInt(pos[1], 10),
		w:  utils.ParseInt(size[0], 10),
		h:  utils.ParseInt(size[1], 10),
	}
}
