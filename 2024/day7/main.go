package main

import (
	. "advent-of-code-2024/utils"
	"strconv"
	"strings"
)

type op byte

const (
	add  op = '+'
	mul  op = '*'
	conc op = '|'
)

func main() {
	DownloadDayInput(2024, 7, false)
	input := SplitNewLine(ReadInputFileRelative())

	var part1 int
	var part2 int
	for _, line := range input {
		teststr, valuesstr, ok := strings.Cut(line, ":")
		AssertEqual(true, ok)
		test := ParseInt(teststr)
		values := Map(strings.Split(strings.TrimSpace(valuesstr), " "), ParseInt)

		if isCalibrated(test, values, []op{add, mul}) {
			part1 += test
		}
		if isCalibrated(test, values, []op{add, mul, conc}) {
			part2 += test
		}
	}
	println("Part 1:", part1)
	println("Part 2:", part2)
}

func isCalibrated(test int, values []int, ops []op) bool {
	if len(values) == 0 {
		return test == 0
	}

	remaining, last := values[:len(values)-1], values[len(values)-1]
	for _, o := range ops {
		var newtest int
		switch o {
		case add:
			newtest = test - last
			if newtest < 0 {
				continue
			}
		case mul:
			if test%last != 0 {
				continue
			}
			newtest = test / last
		case conc:
			slast := strconv.Itoa(last)
			stest := strconv.Itoa(test)
			snewtest, ok := strings.CutSuffix(stest, slast)
			if !ok {
				continue
			}
			if snewtest != "" {
				newtest = ParseInt(snewtest)
			}
		}
		if isCalibrated(newtest, remaining, ops) {
			return true
		}
	}
	return false
}
