package main

import (
	"math"
	"regexp"
	"strings"

	. "advent-of-code-2021/utils"
)

const example = `
NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
`

var reg = regexp.MustCompile(`([A-Z]{2}) -> ([A-Z])`)

func main() {
	DownloadDayInput(2021, 14, false)
	input := SplitEmptySlice(SplitNewLine(ReadInputFileRelative()))
	//input = SplitEmptySlice(SplitNewLine(example[1:]))

	rules := map[string]string{}
	for _, l := range input[1] {
		m := reg.FindStringSubmatch(l)
		if len(m) != 3 {
			panic(l)
		}
		rules[m[1]] = m[2]
	}

	min, max := growPolymer(input[0][0], rules, 10)
	println("Part 1:", max-min)

	min, max = growPolymerFast(input[0][0], rules, 40)
	println("Part 2:", max-min)
}

func growPolymerFast(input string, rules map[string]string, steps int) (min, max int) {

	couples := map[string]int{}
	for i := 0; i < len(input)-1; i++ {
		couples[input[i:i+2]]++
	}

	for i := 0; i < steps; i++ {
		newCouples := map[string]int{}
		for c, n := range couples {
			if letter, ok := rules[c]; ok {
				newCouples[string(c[0])+letter] += n
				newCouples[letter+string(c[1])] += n
			} else {
				newCouples[c] = n
			}
		}
		couples = newCouples
	}

	// Letters are doubled, except once for the extremities
	m := map[rune]int{}
	m[rune(input[0])]++
	m[rune(input[len(input)-1])]++

	for c, n := range couples {
		m[rune(c[0])] += n
		m[rune(c[1])] += n
	}
	min = math.MaxInt64
	max = 0
	for _, n := range m {
		min = Min(min, n)
		max = Max(max, n)
	}
	return min / 2, max / 2
}

func growPolymer(input string, rules map[string]string, steps int) (min, max int) {
	sb := strings.Builder{}
	for i := 0; i < steps; i++ {
		for j := 0; j < len(input)-1; j++ {
			letter := rules[input[j:j+2]]
			sb.WriteRune(rune(input[j]))
			if letter != "" {
				sb.WriteString(letter)
			}
		}
		sb.WriteRune(rune(input[len(input)-1]))
		input = sb.String()
		sb.Reset()
	}

	m := map[rune]int{}
	for _, c := range input {
		m[c]++
	}
	min = math.MaxInt64
	max = 0
	for _, n := range m {
		min = Min(min, n)
		max = Max(max, n)
	}
	return min, max
}
