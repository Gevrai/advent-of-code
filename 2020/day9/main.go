package main

import (
	. "advent-of-code-2020/utils"
)

func main() {
	DownloadDayInput(2020, 9, false)
	input := SplitNewLine(ReadInputFileRelative())
	preamble := 25

	var n []int
	for _, l := range input {
		n = append(n, ParseInt(l, 10))
	}

	invalid := -1
	for i := preamble; i < len(n); i++ {
		if !checkIsSum(n[i], n[i-preamble:i]) {
			invalid = n[i]
			break
		}
	}
	println("Part 1:", invalid)

	i := 0
	j := 0
	sum := 0
	for sum != invalid {
		if sum < invalid {
			sum += n[j]
			j++
		} else {
			sum -= n[i]
			i++
		}
	}

	println("Part 2:", Min(n[i:j]...)+Max(n[i:j]...))
}

func checkIsSum(i int, l []int) bool {

	m := map[int]int{}
	for _, n := range l {
		m[n]++
	}

	for k := range m {
		n := i - k
		if n == k && m[n] > 1 {
			return true
		}
		if m[n] > 0 {
			return true
		}
	}
	return false
}
