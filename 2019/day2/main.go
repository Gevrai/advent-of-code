package main

import (
	"advent-of-code-2019/computer"
	. "advent-of-code-2019/utils"
)

func main() {
	input := ReadInputFileRelative()

	cp, err := computer.InitComputer(input[0])
	if err != nil {
		panic(err.Error())
	}

	cp.Put(1, BigInt(12))
	cp.Put(2, BigInt(2))
	out := <-cp.Run()
	println("Part 1:", out.Int64())

	// Check for output 19690720
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {

			cp, err := computer.InitComputer(input[0])
			if err != nil {
				panic(err.Error())
			}

			cp.Put(1, BigInt(noun))
			cp.Put(2, BigInt(verb))
			out := <-cp.Run()
			if out.Int64() == 19690720 {
				println("Part 2: 100 *", noun, "+", verb, "=", 100*noun+verb)
				return
			}
		}
	}
}
