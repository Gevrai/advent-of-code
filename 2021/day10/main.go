package main

import (
	"sort"

	. "advent-of-code-2021/utils"
)

const example = `
[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]
`

type point struct{ x, y int }

var mapping = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var ill = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var scores = map[string]int{
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

func main() {
	DownloadDayInput(2021, 10, false)
	input := SplitNewLine(ReadInputFileRelative())
	//input = SplitNewLine(example[1:])

	count := 0
	for _, l := range input {
		count += illegal(l)
	}
	println("Part 1:", count)

	var scores []int
	for _, l := range input {
		if illegal(l) != 0 {
			continue
		}
		s := complete(l)
		if s != -1 {
			scores = append(scores, s)
		}
	}
	sort.Ints(scores)
	println("Part 2:", scores[len(scores)/2])
}

func illegal(l string) int {
	var stack []rune
	for _, c := range l {
		switch c {
		case '(', '[', '{', '<':
			stack = append(stack, c)
		case ')', ']', '}', '>':
			if stack[len(stack)-1] != mapping[c] {
				return ill[string(c)]
			}
			stack = stack[:len(stack)-1]
		}
	}
	return 0
}

func complete(l string) (score int) {
	var stack []rune
	for _, c := range l {
		switch c {
		case '(', '[', '{', '<':
			stack = append(stack, c)
		case ')', ']', '}', '>':
			if stack[len(stack)-1] != mapping[c] {
				return -1
			}
			stack = stack[:len(stack)-1]
		}
	}
	for _, c := range Reverse(string(stack)) {
		score = score*5 + scores[string(c)]
	}
	return score
}
