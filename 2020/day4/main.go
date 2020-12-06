package main

import (
	"regexp"
	"strconv"
	"strings"

	. "advent-of-code-2020/utils"
)

type Field struct {
	Mandatory bool
	Validator func(string) bool
}

var validators = map[string]Field{
	"byr": {true, func(s string) bool { return isNumber(s, 1920, 2002) }},
	"iyr": {true, func(s string) bool { return isNumber(s, 2010, 2020) }},
	"eyr": {true, func(s string) bool { return isNumber(s, 2020, 2030) }},
	"hgt": {true, func(s string) bool {
		n, unit := s[:len(s)-2], s[len(s)-2:]
		return (unit == "cm" && isNumber(n, 150, 193)) || (unit == "in" && isNumber(n, 59, 76))
	}},
	"hcl": {true, func(s string) bool { return regexp.MustCompile("^#[0-9a-f]{6}$").MatchString(s) }},
	"ecl": {true, func(s string) bool { return regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$").MatchString(s) }},
	"pid": {true, func(s string) bool { return regexp.MustCompile("^[0-9]{9}$").MatchString(s) }},
	"cid": {false, func(s string) bool { return true }},
}

func main() {
	DownloadDayInput(2020, 4, false)
	input := SplitNewLine(ReadInputFileRelative())
	passports := strings.Split(strings.Join(input, " "), "  ")
	println("Part 1:", Count(passports, containsAllMandatoryFields))
	println("Part 2:", Count(passports, func(s string) bool {
		return containsAllMandatoryFields(s) && allFieldsAreValid(s)
	}))
}

func isNumber(s string, min, max int) bool {
	i, err := strconv.Atoi(s)
	return err == nil && min <= i && i <= max
}

func containsAllMandatoryFields(passport string) bool {
	for k, v := range validators {
		if v.Mandatory && !strings.Contains(passport, k) {
			return false
		}
	}
	return true
}

func allFieldsAreValid(passport string) bool {
	for _, l := range strings.Split(passport, " ") {
		splits := strings.Split(l, ":")
		field, ok := validators[splits[0]]
		if !ok || !field.Validator(strings.TrimSpace(splits[1])) {
			return false
		}
	}
	return true
}
