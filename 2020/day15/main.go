package main

import (
	"math"
	"strings"

	. "advent-of-code-2020/utils"
)

func main() {
	DownloadDayInput(2020, 15, false)
	input := ReadInputFileRelative()

	AssertEqual(series("0,3,6", 2020), 436)

	println("Part 1:", series(input, 2020))
	println("Part 2:", series(input, 30000000))
}

func series(input string, length int) int {
	parts := strings.Split(input, ",")
	nums := make(map[int]int, int(math.Log(float64(length))))
	last := -1
	for i, s := range parts {
		if last != -1 {
			nums[last] = i
		}
		last = ParseInt(s, 10)
	}
	for i := len(parts); i < length; i++ {
		p, ok := nums[last]
		nums[last] = i
		if ok {
			last = i - p
		} else {
			last = 0
		}
	}
	return last
}
