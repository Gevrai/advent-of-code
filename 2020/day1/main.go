package main

import (
	"strconv"
	"strings"

	. "advent-of-code-2020/utils"
)

func main() {
	DownloadDayInput(2020, 1, false)
	input := SplitNewLine(ReadInputFileRelative())
	entries := createEntriesSet(input)

	println("Part 1:", multArray(sumTo(entries, 2, 2020)))
	println("Part 2:", multArray(sumTo(entries, 3, 2020)))
}

func createEntriesSet(input []string) map[int]struct{} {
	entries := make(map[int]struct{})
	for _, s := range input {
		if s == "" {
			continue
		}
		i, err := strconv.Atoi(strings.TrimSpace(s))
		PanicIfError(err)
		entries[i] = struct{}{}
	}
	return entries
}

func sumTo(entries map[int]struct{}, nb, sum int) []int {

	for k := range entries {
		if nb == 2 {
			diff := sum - k
			if _, ok := entries[diff]; ok {
				return []int{k, diff}
			}
		} else {
			e := sumTo(entries, nb-1, sum-k)
			if e != nil {
				return append(e, k)
			}
		}
	}
	return nil
}

func multArray(e []int) int {
	mult := 1
	for _, i := range e {
		mult *= i
	}
	return mult
}
