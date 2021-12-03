package main

import (
	. "advent-of-code-2021/utils"
)

func main() {
	DownloadDayInput(2021, 3, false)
	input := SplitNewLine(ReadInputFileRelative())

	var gamma, epsilon string
	for _, c := range getRates(len(input[0]), input) {
		if c < (len(input) / 2) {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	println("Part 1:", ParseInt(gamma, 2)*ParseInt(epsilon, 2))

	l := len(input[0])

	input = SplitNewLine(ReadInputFileRelative())
	for i := 0; i < l && len(input) > 1; i++ {
		rates := getRates(l, input)
		c := uint8('1')
		if float32(rates[i]) < (float32(len(input)) / float32(2)) {
			c = '0'
		}
		input = Filter(input, func(s string) bool { return s[i] == c })
	}
	oxy := ParseInt(input[0], 2)

	input = SplitNewLine(ReadInputFileRelative())
	for i := 0; i < l && len(input) > 1; i++ {
		rates := getRates(l, input)
		c := uint8('0')
		if float32(rates[i]) < (float32(len(input)) / float32(2)) {
			c = '1'
		}
		input = Filter(input, func(s string) bool { return s[i] == c })
	}
	co2 := ParseInt(input[0], 2)
	println("Part 2:", oxy*co2)
}

func getRates(l int, input []string) []int {
	rates := make([]int, l)
	for _, s := range input {
		for i, c := range s {
			if c == '1' {
				rates[i]++
			}
		}
	}
	return rates
}
