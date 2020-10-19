package main

import "advent-of-code-2019/utils"

func main() {
	input := utils.ReadInputFileRelative()

	wire1 := CreateWire(input[0])
	wire2 := CreateWire(input[1])

	println("Part 1:", wire1.ClosestIntersection(wire2).ManhattanFromOrigin())
	println("Part 2:", wire1.MinimizedSignalIntersection(wire2))

}
