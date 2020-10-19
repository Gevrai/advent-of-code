package main

import (
	"advent-of-code-2019/computer"
	. "advent-of-code-2019/utils"
	"fmt"
)

func main() {
	input := ReadInputFileRelative()

	cp, err := computer.InitComputer(input[0])
	if err != nil {
		panic(err.Error())
	}

	go cp.Input(BigInt(1))
	cp.Run()

	var outputs []int64
	for val := range cp.Output() {
		outputs = append(outputs, val.Int64())
	}
	failures := 0
	for val := range outputs[:len(outputs)-1] {
		if val != 0 {
			failures++
		}
	}

	fmt.Printf("Part 1: %d (%d failure(s))\n", outputs[len(outputs)-1], failures)

	cp, err = computer.InitComputer(input[0])
	if err != nil {
		panic(err.Error())
	}
	go cp.Input(BigInt(5))
	cp.Run()

	outputs = []int64{}
	for val := range cp.Output() {
		outputs = append(outputs, val.Int64())
	}

	fmt.Printf("Part 2: %d (%d output(s))\n", outputs[0], len(outputs))

}
