package main

import (
	"strings"

	. "advent-of-code-2021/utils"
)

func main() {
	DownloadDayInput(2021, 6, false)
	input := ReadInputFileRelative()

	sim := make([]int, 9)
	for _, l := range strings.Split(input, ",") {
		sim[ParseInt(l)]++
	}

	for i := 0; i < 80; i++ {
		zero := sim[0]
		copy(sim, sim[1:])
		sim[6] += zero
		sim[8] = zero
	}

	println("Part 1:", Sum(sim))

	sim = make([]int, 9)
	for _, l := range strings.Split(input, ",") {
		sim[ParseInt(l)]++
	}
	for i := 0; i < 256; i++ {
		zero := sim[0]
		copy(sim, sim[1:])
		sim[6] += zero
		sim[8] = zero
	}
	println("Part 2:", Sum(sim))
}
