package main

import (
	"strings"

	. "advent-of-code-2020/utils"
)

func main() {
	DownloadDayInput(2020, 7, false)
	//input := ReadInputFileRelative()
	input := SplitNewLine(ReadInputFileRelative())

	bags := map[string][]Bag{}
	for _, l := range input {
		parts := strings.Split(l, "contain")

		color := getColor(parts[0])

		if strings.Contains(parts[1], "no other") {
			bags[color] = nil
			continue
		}

		internal := []Bag{}
		for _, b := range strings.Split(parts[1], ",") {
			b = strings.TrimSpace(b)
			internal = append(internal, Bag{
				ammount: ParseInt(string(b[0]), 10),
				color:   getColor(b[1:]),
			})
		}
		bags[color] = internal
	}

	s := map[string]bool{
		"shiny gold": true,
	}
	l := 0
	for l != len(s) {
		l = len(s)
		for color, colors := range bags {
			for _, c := range colors {
				for k := range s {
					if strings.Contains(c.color, k) {
						s[color] = true
					}
				}
			}
		}
	}
	println("Part 1:", len(s)-1)

	println("Part 2:", contains("shiny gold", bags))
}

func getColor(color string) string {
	color = strings.ReplaceAll(color, "bags", "")
	color = strings.ReplaceAll(color, "bag", "")
	color = strings.Trim(color, ".")
	color = strings.TrimSpace(color)
	return color
}

func contains(color string, bags map[string][]Bag) int {
	thisBag := bags[color]
	ammount := 0
	for _, b := range thisBag {
		ammount += b.ammount + b.ammount*contains(b.color, bags)
	}
	return ammount
}

type Bag struct {
	ammount int
	color   string
}
