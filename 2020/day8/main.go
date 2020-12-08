package main

import (
	"fmt"
	"strings"

	. "advent-of-code-2020/utils"
)

func main() {
	DownloadDayInput(2020, 8, false)
	input := SplitNewLine(ReadInputFileRelative())

	acc, _ := execute(input)
	println("Part 1:", acc)

	for i := range input {
		current := input[i]
		op := strings.Split(input[i], " ")[0]
		switch op {
		case "nop":
			input[i] = strings.Replace(current, "nop", "jmp", 1)
		case "acc":
			continue
		case "jmp":
			input[i] = strings.Replace(current, "jmp", "nop", 1)
		default:
			panic("Unknown op " + op)
		}

		acc, err := execute(input)
		if err == nil {
			println("Part 2:", acc)
		}
		input[i] = current
	}

}

func execute(input []string) (acc int, err error) {

	seen := map[int]bool{}
	for i := 0; i < len(input); i++ {
		parts := strings.Split(input[i], " ")

		if seen[i] {
			return acc, fmt.Errorf("rerunning %d", i)
		}
		seen[i] = true

		op := strings.TrimSpace(parts[0])
		val := ParseInt(parts[1], 10)

		switch op {
		case "nop":
			// noop
		case "acc":
			acc += val
		case "jmp":
			i += val - 1
		default:
			panic("Unknown op " + op)
		}
	}
	return acc, nil
}
