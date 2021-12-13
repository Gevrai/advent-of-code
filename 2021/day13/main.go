package main

import (
	"fmt"
	"regexp"
	"strings"

	. "advent-of-code-2021/utils"
)

const example = `
6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5
`

type coord struct {
	x, y int
}

func main() {
	DownloadDayInput(2021, 13, false)
	input := SplitEmptySlice(SplitNewLine(ReadInputFileRelative()))
	//input = SplitEmptySlice(SplitNewLine(example[1:]))

	coords := map[coord]bool{}
	for _, l := range input[0] {
		parts := strings.Split(l, ",")
		coords[coord{
			x: ParseInt(parts[0]),
			y: ParseInt(parts[1]),
		}] = true
	}
	for i, f := range input[1] {
		m := regexp.MustCompile(`fold along ([x|y])=(\d+)`).FindStringSubmatch(f)
		if len(m) == 0 {
			panic(f)
		}
		switch m[1] {
		case "x":
			foldX(coords, ParseInt(m[2]))
		case "y":
			foldY(coords, ParseInt(m[2]))
		default:
			panic(f)
		}
		if i == 0 {
			println("Part 1:", len(coords))
		}

	}
	println("Part 2:")
	print(coords)
}

func print(coords map[coord]bool) {
	var min, max coord
	for c := range coords {
		min = c
		max = c
	}
	for c := range coords {
		min.x = Min(min.x, c.x)
		min.y = Min(min.y, c.y)
		max.x = Max(max.x, c.x)
		max.y = Max(max.y, c.y)
	}
	for j := min.y; j <= max.y; j++ {
		for i := min.x; i <= max.x; i++ {
			_, ok := coords[coord{i, j}]
			if ok {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func foldX(coords map[coord]bool, x int) {
	for c := range coords {
		if c.x > x {
			delete(coords, c)
			coords[coord{
				x: x - (c.x - x),
				y: c.y,
			}] = true
		}
	}
}

func foldY(coords map[coord]bool, y int) {
	for c := range coords {
		if c.y > y {
			delete(coords, c)
			coords[coord{
				x: c.x,
				y: y - (c.y - y),
			}] = true
		}
	}
}
