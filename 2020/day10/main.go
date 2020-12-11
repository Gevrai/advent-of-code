package main

import (
	"sort"

	. "advent-of-code-2020/utils"
)

func main() {
	DownloadDayInput(2020, 10, false)
	input := SplitNewLine(ReadInputFileRelative())

	n := []int{0}
	for _, l := range input {
		n = append(n, ParseInt(l, 10))
	}
	sort.Ints(n)
	n = append(n, n[len(n)-1]+3)

	one := 0
	three := 0
	for i := range n[:len(n)-1] {
		k := n[i+1] - n[i]
		if k == 1 {
			one++
		}
		if k == 3 {
			three++
		}
	}
	println("Part 1:", one*three)

	const maxJump = 3
	m := make([]int, len(n))
	m[0] = 1
	for i := 1; i < len(n); i++ {
		count := 0
		for j := 1; j <= maxJump; j++ {
			if i-j >= 0 && n[i]-n[i-j] <= 3 {
				count += m[i-j]
			}
		}
		m[i] = count
	}
	println("Part 2:", m[len(m)-1])
}

func exists(i int, n []int) bool {
	return i > 0 && i < len(n)
}
