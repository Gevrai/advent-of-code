package main

import (
	. "advent-of-code-2024/utils"
	"strconv"
	"strings"
)

func main() {
	DownloadDayInput(2024, 2, false)
	input := ReadInputFileRelative()
	levels := SplitNewLine(input)

	var res int
	for _, l := range levels {
		ns := []int{}
		for _, n := range strings.Split(strings.TrimSpace(l), " ") {
			ns = append(ns, Must(strconv.Atoi(n)))
		}
		if safe(ns) {
			res++
		}
	}
	println("Part 1:", res)

	res = 0
	for _, l := range levels {
		ns := []int{}
		for _, n := range strings.Split(strings.TrimSpace(l), " ") {
			ns = append(ns, Must(strconv.Atoi(n)))
		}
		for i := range ns {
			nss := make([]int, 0, len(ns)-1)
			nss = append(nss, ns[:i]...)
			nss = append(nss, ns[i+1:]...)
			if safe(nss) {
				res++
				break
			}
		}
	}
	println("Part 2:", res)
}

func safe(ns []int) bool {
	ok := true
	for i := range ns[1:] {
		diff := ns[i+1] - ns[i]
		if diff < 1 || diff > 3 {
			ok = false
			break
		}
	}
	if ok {
		return true
	}
	ok = true
	for i := range ns[1:] {
		diff := ns[i] - ns[i+1]
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}
