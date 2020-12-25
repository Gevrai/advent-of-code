package main

import (
	"math"
	"strings"

	. "advent-of-code-2020/utils"
)

func main() {
	DownloadDayInput(2020, 13, false)
	input := SplitNewLine(ReadInputFileRelative())

	embarkTime := ParseInt(input[0], 10)

	closest := math.MaxInt32
	bus := 0
	for _, b := range strings.Split(input[1], ",") {
		if b == "x" {
			continue
		}
		m := ParseInt(b, 10)
		t := m - (embarkTime % m)
		if t < closest {
			closest = t
			bus = m
		}
	}
	println("Part 1:", closest*bus)

	AssertEqual(earliest("17,x,13,19"), 3417)
	println("Part 2:", earliest(input[1]))
}

func earliest(input string) int {
	buses := strings.Split(input, ",")
	mod := 1
	rest := 0
	for t, b := range buses {
		println("bus", b, t+1, "/", len(buses))
		if b == "x" {
			continue
		}
		busID := ParseInt(b, 10)
		// find 'phase' between all previous buses and this one. rest is the remaining time before leaving (busID - t)
		result := inPhase(mod, rest, busID, busID-t)
		// Prepare next iteration (see excalidraw board
		rest = result
		mod = mod * busID
	}
	return rest
}

// Bruteforcing, finds n where following is true
//	n mod x = xoffset
//	n mod y = yoffset
func inPhase(x, xoffset, y, yoffset int) int {
	i := xoffset
	j := yoffset

	for i != j {
		if i < j {
			d := (j - i) / x
			if d < 1 {
				d = 1
			}
			i += d * x
		} else {
			d := (i - j) / y
			if d < 1 {
				d = 1
			}
			j += d * y
		}
	}
	return i
}
