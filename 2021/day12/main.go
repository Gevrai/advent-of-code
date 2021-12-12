package main

import (
	"strings"

	. "advent-of-code-2021/utils"
)

const example = `
start-A
start-b
A-c
A-b
b-d
A-end
b-end
`

type cave struct {
	name        string
	isBig       bool
	connections []*cave
}

func main() {
	DownloadDayInput(2021, 12, false)
	input := SplitNewLine(ReadInputFileRelative())
	//input = SplitNewLine(example[1:])

	m := map[string]*cave{}
	for _, l := range input {
		parts := strings.Split(l, "-")
		n1, n2 := parts[0], parts[1]

		if m[n1] == nil {
			m[n1] = &cave{name: n1, isBig: strings.ToUpper(n1) == n1}
		}
		if m[n2] == nil {
			m[n2] = &cave{name: n2, isBig: strings.ToUpper(n2) == n2}
		}
		m[n1].connections = append(m[n1].connections, m[n2])
		m[n2].connections = append(m[n2].connections, m[n1])
	}
	println("Part 1:", visit(m["start"], m["end"], m, map[string]bool{}, false))
	println("Part 2:", visit(m["start"], m["end"], m, map[string]bool{}, true))
}

func visit(curr, end *cave, m map[string]*cave, visited map[string]bool, remainingSmallVisit bool) (nbPaths int) {
	if curr == end {
		return 1
	}

	if !curr.isBig {
		if visited[curr.name] {
			if !remainingSmallVisit || curr.name == "start" {
				return 0
			}
			remainingSmallVisit = false
		} else {
			visited[curr.name] = true
			defer delete(visited, curr.name)
		}
	}

	for _, conn := range curr.connections {
		c := visit(conn, end, m, visited, remainingSmallVisit)
		nbPaths += c
	}
	return nbPaths
}
