package main

import (
	"strconv"
	"strings"

	"advent-of-code-2020/utils"
)

func main() {
	utils.DownloadDayInput(2020, 4, false)
	input := utils.ReadInputFileRelative()

	passport := []string{""}
	current := 0
	for _, l := range input {
		if strings.TrimSpace(l) == "" {
			current++
			passport = append(passport, "")
		} else {
			passport[current] += " " + l
		}
	}

	mandatory := []string{
		"byr:",
		"iyr:",
		"eyr:",
		"hgt:",
		"hcl:",
		"ecl:",
		"pid:",
	}

	count := 0
	for _, p := range passport {
		invalid := false
		for _, f := range mandatory {
			if !strings.Contains(p, f) {
				invalid = true
			}
		}
		if !invalid {
			count++
		}
	}

	println("Part 1:", count)

	isNumber := func(s string, min, max int) bool {
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return min <= i && i <= max
	}

	validators := map[string]func(string) bool{
		"byr": func(s string) bool { return len(s) == 4 && isNumber(s, 1920, 2002) },
		"iyr": func(s string) bool { return len(s) == 4 && isNumber(s, 2010, 2020) },
		"eyr": func(s string) bool { return len(s) == 4 && isNumber(s, 2020, 2030) },
		"hgt": func(s string) bool {
			return (strings.Contains(s, "cm") && isNumber(s[:len(s)-2], 150, 193)) ||
				(strings.Contains(s, "in") && isNumber(s[:len(s)-2], 59, 76))
		},
		"hcl": func(s string) bool {
			if s[0] != '#' || len(s) != 7 {
				return false
			}
			for _, c := range "0123456789abcdef" {
				s = strings.ReplaceAll(s, string(c), "")
			}
			return len(s) == 1 // only '#' is left
		},
		"ecl": func(s string) bool {
			return s == "amb" || s == "blu" || s == "brn" || s == "gry" || s == "grn" || s == "hzl" || s == "oth"
		},
		"pid": func(s string) bool { return len(s) == 9 && isNumber(s, 0, 10e10) },
		"cid": func(s string) bool { return true },
	}

	count = 0
	for _, p := range passport {
		invalid := false
		for _, l := range strings.Split(p, " ") {
			if l == "" {
				continue
			}
			splits := strings.Split(l, ":")
			v, ok := validators[splits[0]]
			if !ok || !v(strings.TrimSpace(splits[1])) {
				invalid = true
			}
		}
		for _, f := range mandatory {
			if !strings.Contains(p, f) {
				invalid = true
			}
		}
		if !invalid {
			count++
		}
	}

	println("Part 2:", count)
}
