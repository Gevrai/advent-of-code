package main

import (
	"advent-of-code-2019/computer"
	. "advent-of-code-2019/utils"
)

func main() {
	input := ReadInputFileRelative()

	cp, err := computer.InitComputer(input[0])
	if err != nil {
		panic(err)
	}
	cp.Input(BigInt(1))
	cp.Run()

	var last *int64
	for out := range cp.Output() {
		if last != nil {
			println("Malfunctioning opcode:", *last)
		}
		i := out.Int64()
		last = &i
	}

	println("Part one:", *last)

	cp, err = computer.InitComputer(input[0])
	if err != nil {
		panic(err)
	}
	cp.Input(BigInt(2))
	cp.Run()

	for out := range cp.Output() {
		println("Part two:", out.Int64())
	}

}
