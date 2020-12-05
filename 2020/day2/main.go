package main

import (
	"strconv"
	"strings"

	"advent-of-code-2020/utils"
)

func main() {
	utils.DownloadDayInput(2020, 2, false)
	input := utils.ReadInputFileRelativeSplitNewline()
	println("Part 1:", utils.Count(input, isValidFirstPolicy))
	println("Part 2:", utils.Count(input, isValidSecondPolicy))
}

func isValidFirstPolicy(line string) bool {

	min, max, char, password := parseInput(line)

	n := strings.Count(password, char)
	return n >= min && n <= max
}

func isValidSecondPolicy(line string) bool {

	a, b, char, password := parseInput(line)

	aValid := a <= (len(password)) && string(password[a-1]) == char
	bValid := b <= (len(password)) && string(password[b-1]) == char
	return aValid != bValid
}

func parseInput(input string) (min, max int, char, password string) {

	splits := strings.Split(input, ":")
	password = strings.TrimSpace(splits[1])
	splits = strings.Split(splits[0], " ")
	char = strings.TrimSpace(splits[1])
	splits = strings.Split(splits[0], "-")

	var err error
	min, err = strconv.Atoi(strings.TrimSpace(splits[0]))
	utils.PanicIfError(err)
	max, err = strconv.Atoi(strings.TrimSpace(splits[1]))
	utils.PanicIfError(err)
	return
}
