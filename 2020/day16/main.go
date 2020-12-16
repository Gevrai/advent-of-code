package main

import (
	"strings"

	. "advent-of-code-2020/utils"
)

func main() {
	DownloadDayInput(2020, 16, false)
	input := SplitNewLine(ReadInputFileRelative())

	i := 0
	rules := map[string]Ranges{}
	for {
		line := strings.TrimSpace(input[i])
		i++
		if line == "" {
			break
		}
		parts := strings.Split(line, ":")
		rule := parts[0]
		ranges := Ranges{}
		for _, r := range strings.Split(parts[1], "or") {
			r = strings.TrimSpace(r)
			parts := strings.Split(r, "-")
			ranges = append(ranges, Range{
				min: ParseInt(parts[0], 10),
				max: ParseInt(parts[1], 10),
			})
		}
		rules[rule] = ranges
	}

	AssertEqual(input[i], "your ticket:")
	i++
	myTicket := asTicket(input[i])
	println(myTicket)
	i++
	i++

	AssertEqual(input[i], "nearby tickets:")
	i++
	scanningErrorRate := 0
	tickets := [][]int{}
	for ; i < len(input); i++ {
		t := asTicket(input[i])

		allValid := true
		for _, entry := range t {
			valid := false
			for _, rule := range rules {
				if rule.matches(entry) {
					valid = true
					break
				}
			}
			if !valid {
				scanningErrorRate += entry
				allValid = false
			}
		}
		if allValid {
			tickets = append(tickets, t)
		}
	}

	println("Part 1:", scanningErrorRate)

	// Get all possible matches
	rulesPotentialIndex := map[string][]int{}
	for name, rule := range rules {
		for i := range tickets[0] {
			matchesAll := true
			for j := range tickets {
				if !rule.matches(tickets[j][i]) {
					matchesAll = false
					break
				}
			}
			if matchesAll {
				rulesPotentialIndex[name] = append(rulesPotentialIndex[name], i)
			}
		}
	}

	// Go by elimination, checking rule which only has single possibility
	singleRuleIndex := map[string]int{}
	for len(rulesPotentialIndex) > 0 {
		rule := getSingle(rulesPotentialIndex)
		index := rulesPotentialIndex[rule][0]
		singleRuleIndex[rule] = index
		delete(rulesPotentialIndex, rule)

		for rule := range rulesPotentialIndex {
			rulesPotentialIndex[rule] = removeValue(index, rulesPotentialIndex[rule])
		}
	}

	departures := 1
	for name := range singleRuleIndex {
		if strings.HasPrefix(name, "departure") {
			departures *= myTicket[singleRuleIndex[name]]
		}
	}
	println("Part 2:", departures)
}

type Range struct {
	min, max int
}

func (r *Range) inRange(i int) bool {
	return r.min <= i && r.max >= i
}

type Ranges []Range

func (r *Ranges) matches(i int) bool {
	for _, rg := range *r {
		if rg.inRange(i) {
			return true
		}
	}
	return false
}

func removeValue(val int, list []int) []int {
	for i := range list {
		if val == list[i] {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func getSingle(rules map[string][]int) string {
	for name, list := range rules {
		if len(list) == 1 {
			return name
		}
	}
	panic("no single")
}

func asTicket(input string) (ticket []int) {
	for _, n := range strings.Split(input, ",") {
		ticket = append(ticket, ParseInt(n, 10))
	}
	return ticket
}
