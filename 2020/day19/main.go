package main

import (
	"fmt"
	"strings"

	. "advent-of-code-2020/utils"
)

var rules = make(map[int]Rule)

func main() {
	DownloadDayInput(2020, 19, false)
	input := SplitNewLine(ReadInputFileRelative())

	for _, l := range input {
		if l == "" {
			break
		}
		parts := strings.Split(l, ":")
		id := ParseInt(parts[0])
		rule := NewRule(parts[1])
		rules[id] = rule
	}

	println("Part 1:", MatchAll(input[len(rules)+1:]))

	rules[8] = NewRule("42 | 42 8")
	rules[11] = NewRule("42 31 | 42 11 31")
	println("Part 2:", MatchAll(input[len(rules)+1:]))
}

func MatchAll(messages []string) (count int) {
	for _, message := range messages {
		ok, rest := rules[0].Matches(message)
		if ok {
			for _, r := range rest {
				if r == "" {
					count++
					break
				}
			}
		}
	}
	return count
}

func NewRule(input string) Rule {
	input = strings.TrimSpace(input)
	if strings.Contains(input, `"`) {
		return leaf(Trims(input, `"`))
	}

	parts := strings.Split(input, "|")
	rule := orRule{}
	for _, p := range parts {
		cr := concatRule{}

		for _, p := range strings.Split(strings.TrimSpace(p), " ") {
			i := ParseInt(p)
			cr = append(cr, node(i))
		}
		rule = append(rule, cr)
	}

	return rule
}

type Rule interface {
	Matches(string) (bool, []string)
}

type leaf string

func (l leaf) Matches(input string) (bool, []string) {
	if len(input) < len(l) {
		return false, nil
	}
	if input[:len(l)] == string(l) {
		return true, []string{input[len(l):]}
	}
	return false, nil
}

type node int

func (n node) Matches(input string) (bool, []string) {
	r, ok := rules[int(n)]
	if !ok {
		panic(fmt.Sprintf("rule %d does not exist", int(n)))
	}
	return r.Matches(input)
}

type concatRule []Rule

func (c concatRule) Matches(input string) (matches bool, rests []string) {
	if len(c) == 0 {
		return true, []string{input}
	}
	if len(input) == 0 {
		return false, nil
	}

	ok, rest := c[0].Matches(input)
	if !ok {
		return false, nil
	}

	for _, r := range rest {
		ok, newrest := c[1:].Matches(r)
		if ok {
			matches = true
			rests = append(rests, newrest...)
		}
	}
	return
}

type orRule []Rule

func (r orRule) Matches(input string) (matches bool, rests []string) {
	for _, r := range r {
		ok, rest := r.Matches(input)
		if ok {
			matches = true
			rests = append(rests, rest...)
		}
	}
	return
}
