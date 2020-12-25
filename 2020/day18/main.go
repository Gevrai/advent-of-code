package main

import (
	"regexp"
	"strconv"
	"strings"

	. "advent-of-code-2020/utils"
)

func main() {
	DownloadDayInput(2020, 18, false)
	input := SplitNewLine(ReadInputFileRelative())

	total := 0
	for _, eq := range input {
		total += eval(eq)
	}
	println("Part 1:", total)

	total = 0
	for _, eq := range input {
		total += advanceEval(eq)
	}

	println("Part 2:", total)
}

var (
	parens = regexp.MustCompile(`\s*\(([^()]*)\)\s*`)
	oper   = regexp.MustCompile(`\s*(\d+)\s*([+|*])\s*(\d+)\s*`)
	add    = regexp.MustCompile(`\s*(\d+)\s*([+])\s*(\d+)\s*`)
	mult   = regexp.MustCompile(`\s*(\d+)\s*([*])\s*(\d+)\s*`)
)

func eval(eq string) int {
	for {
		// Evaluate biggest parenthesis
		m := parens.FindStringSubmatch(eq)
		if len(m) > 0 {
			res := eval(m[1])
			eq = strings.Replace(eq, m[0], strconv.Itoa(res), 1)
			continue
		}

		// Evaluate first possible operation
		m = oper.FindStringSubmatch(eq)
		if len(m) > 0 {
			var res int
			switch m[2] {
			case "*":
				res = ParseInt(m[1], 10) * ParseInt(m[3], 10)
			case "+":
				res = ParseInt(m[1], 10) + ParseInt(m[3], 10)
			}
			eq = strings.Replace(eq, m[0], strconv.Itoa(res), 1)
			continue
		}

		// Didn't do anything, finished!
		break
	}
	return ParseInt(eq, 10)
}

func advanceEval(eq string) int {
	for {
		// Evaluate biggest parenthesis
		m := parens.FindStringSubmatch(eq)
		if len(m) > 0 {
			res := advanceEval(m[1])
			eq = strings.Replace(eq, m[0], strconv.Itoa(res), 1)
			continue
		}

		// Evaluate first possible add operation
		m = add.FindStringSubmatch(eq)
		if len(m) > 0 {
			res := ParseInt(m[1], 10) + ParseInt(m[3], 10)
			eq = strings.Replace(eq, m[0], strconv.Itoa(res), 1)
			continue
		}

		// Evaluate first possible mult operation
		m = mult.FindStringSubmatch(eq)
		if len(m) > 0 {
			res := ParseInt(m[1], 10) * ParseInt(m[3], 10)
			eq = strings.Replace(eq, m[0], strconv.Itoa(res), 1)
			continue
		}

		// Didn't do anything, finished!
		break
	}
	return ParseInt(eq, 10)

}
