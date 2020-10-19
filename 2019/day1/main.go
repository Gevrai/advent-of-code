package main

import (
	"advent-of-code-2019/utils"
	"strconv"
)

func main() {
	input := utils.ReadInputFileRelative()

	totalFuelPart1 := 0
	totalFuelPart2 := 0
	for _, line := range input {
		mass, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		totalFuelPart1 += utils.FuelNeeded(mass)
		totalFuelPart2 += utils.FuelNeededIncludingFuelMass(mass)
	}

	println("Part 1:", totalFuelPart1)
	println("Part 2:", totalFuelPart2)
}
