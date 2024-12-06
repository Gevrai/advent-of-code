package main

import (
	. "advent-of-code-2024/utils"
	"slices"
	"strings"
)

func main() {
	DownloadDayInput(2024, 5, false)
	input := ReadInputFileRelative()
	split := SplitEmptySlice(SplitNewLine(input))

	countOrdered := 0
	countUnordered := 0

	before := map[int][]int{}
	after := map[int][]int{}
	for _, s := range split[0] {
		a, b, _ := strings.Cut(s, "|")
		before[ParseInt(b)] = append(before[ParseInt(b)], ParseInt(a))
		after[ParseInt(a)] = append(after[ParseInt(a)], ParseInt(b))
	}

	for _, line := range split[1] {
		var pages []int
		for _, n := range strings.Split(line, ",") {
			pages = append(pages, ParseInt(n))
		}
		ordered := func() bool {
			for i, n := range pages {
				if ms, ok := before[n]; ok {
					for _, m := range ms {
						if slices.Contains(pages[i+1:], m) {
							return false
						}
					}
				}
			}
			return true
		}()
		if ordered {
			countOrdered += pages[len(pages)/2]
			continue
		}
		slices.SortStableFunc(pages, func(a, b int) int {
			if ms, ok := before[b]; ok {
				if slices.Contains(ms, a) {
					return -1
				}
			}
			if ms, ok := after[b]; ok {
				if slices.Contains(ms, a) {
					return 1
				}
			}
			return 0
		})
		countUnordered += pages[len(pages)/2]
	}
	println("Part 1:", countOrdered)
	println("Part 2:", countUnordered)
}
