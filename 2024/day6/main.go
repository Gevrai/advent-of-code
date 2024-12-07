package main

import (
	. "advent-of-code-2024/utils"
	"strings"
)

func main() {
	DownloadDayInput(2024, 6, false)
	input := SplitNewLine(ReadInputFileRelative())

	type vec2 struct{ x, y int }

	initialpos := vec2{}

	for j := range input {
		i := strings.IndexByte(input[j], '^')
		if i > 0 {
			initialpos = vec2{i, j}
			break
		}
	}

	max := vec2{len(input[0]), len(input)}
	pos := initialpos
	dir := vec2{0, -1}
	distinct := map[vec2]struct{}{}
	for {
		distinct[pos] = struct{}{}
		n := vec2{pos.x + dir.x, pos.y + dir.y}
		if n.x < 0 || n.x >= max.x || n.y < 0 || n.y >= max.y {
			break
		}
		if input[n.y][n.x] == '#' {
			dir.x, dir.y = -dir.y, dir.x
			continue
		}
		pos = n
	}
	println("Part 1:", len(distinct))

	count := 0
	for j := range input {
		for i := range input[j] {
			if input[j][i] != '.' {
				continue
			}

			looped := func(pos vec2) bool {
				obstacle := vec2{i, j}
				dir := vec2{0, -1}
				type move struct{ pos, dir vec2 }
				moves := map[move]struct{}{}
				for {
					moves[move{pos, dir}] = struct{}{}
					n := vec2{pos.x + dir.x, pos.y + dir.y}
					if n.x < 0 || n.x >= max.x || n.y < 0 || n.y >= max.y {
						return false
					}
					if n == obstacle || input[n.y][n.x] == '#' {
						dir.x, dir.y = -dir.y, dir.x
						continue
					}
					pos = n
					if _, ok := moves[move{pos, dir}]; ok {
						return true
					}
				}
			}(initialpos)
			if looped {
				count++
			}
		}
	}
	println("Part 2:", count)
}
