package main

import (
	"advent-of-code-2019/computer"
	. "advent-of-code-2019/utils"
	"math"
)

func main() {
	input := ReadInputFileRelative()

	maxOutput := int64(math.MinInt64)
	for _, perm := range allPermutations([]int64{0, 1, 2, 3, 4}) {
		out := RunAmplifiersWithPhaseSettings(input[0], perm)
		if out > maxOutput {
			maxOutput = out
		}
	}

	println("Part 1:", maxOutput)

	maxOutput = math.MinInt64
	for _, perm := range allPermutations([]int64{5, 6, 7, 8, 9}) {
		out := RunAmplifiersWithPhaseSettingsFeedback(input[0], perm)
		if out > maxOutput {
			maxOutput = out
		}
	}

	println("Part 2:", maxOutput)
}

func allPermutations(base []int64) (permutations [][]int64) {
	if len(base) == 0 {
		return nil
	}
	if len(base) == 1 {
		return [][]int64{{base[0]}}
	}

	for i, n := range base {
		base := append([]int64(nil), base...) //slice copy
		for _, perm := range allPermutations(append(base[:i], base[i+1:]...)) {
			permutations = append(permutations, append([]int64{n}, perm...))
		}
	}
	return
}

func RunAmplifiersWithPhaseSettings(input string, phaseSettings []int64) int64 {

	// Create amplificators and give them phaser settings
	amplifiers := make([]computer.Computer, len(phaseSettings))
	for i, phaseSetting := range phaseSettings {
		amp, err := computer.InitComputer(input)
		if err != nil {
			panic(err)
		}
		amplifiers[i] = amp

		amp.Input(BigInt(phaseSetting))
		amp.Run()
	}

	// Connect amplifiers
	amplifiers[0].Input(BigInt(0))
	for i := 1; i < len(amplifiers); i++ {
		go ConnectAmplifiers(amplifiers[i-1], amplifiers[i])
	}
	out := <-amplifiers[len(amplifiers)-1].Output()
	return out.Int64()
}

func RunAmplifiersWithPhaseSettingsFeedback(input string, phaseSettings []int64) int64 {

	// Create amplificators and give them phaser settings
	amplifiers := make([]computer.Computer, len(phaseSettings))
	for i, phaseSetting := range phaseSettings {
		amp, err := computer.InitComputer(input)
		if err != nil {
			panic(err)
		}
		amplifiers[i] = amp

		amp.Input(BigInt(phaseSetting))
		amp.Run()
	}

	// Connect amplifiers
	amplifiers[0].Input(BigInt(0))
	for i := 1; i < len(amplifiers); i++ {
		go ConnectAmplifiers(amplifiers[i-1], amplifiers[i])
	}
	// Feedback to first, return value that can't be inputed (amplifiers halted)
	for out := range amplifiers[len(amplifiers)-1].Output() {
		if amplifiers[0].Input(out) != nil {
			return out.Int64()
		}
	}
	panic("should have returned a value")
}

func ConnectAmplifiers(output, input computer.Computer) {
	for out := range output.Output() {
		err := input.Input(out)
		if err != nil {
			panic(err)
		}
	}
}
