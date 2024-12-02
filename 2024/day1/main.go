package main

import (
	. "advent-of-code-2024/utils"
	"slices"
	"strconv"
	"strings"
)

func main() {
	DownloadDayInput(2024, 1, false)
	input := ReadInputFileRelative()

	left := make([]int, 0, len(input))
	right := make([]int, 0, len(input))
	for _, line := range SplitNewLine(input) {
		l, r, ok := strings.Cut(line, " ")
		AssertEqual(true, ok)
		left = append(left, Must(strconv.Atoi(strings.TrimSpace(l))))
		right = append(right, Must(strconv.Atoi(strings.TrimSpace(r))))
	}
	slices.Sort(left)
	slices.Sort(right)

	var sum int
	for i := range left {
		sum += Abs(left[i] - right[i])
	}
	println("Part 1:", sum)

	rightCount := map[int]int{}
	for _, n := range right {
		rightCount[n] = rightCount[n] + 1
	}

	sum = 0
	for _, n := range left {
		sum += n * rightCount[n]
	}
	println("Part 2:", sum)
}
