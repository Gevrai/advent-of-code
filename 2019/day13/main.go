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
	screen := NewMapScreen()

	arcade := Arcade{
		cp:     cp,
		screen: screen,
	}

	arcade.RunGame()
	arcade.DrawScreen()

	println("Part one:", arcade.CountTiles(Block))

	cp, err = computer.InitComputer(input[0])
	if err != nil {
		panic(err)
	}
	// Put some coins in!
	cp.Put(0, BigInt(2))
	screen = NewMapScreen()

	arcade = Arcade{
		cp:     cp,
		screen: screen,
	}

	score := arcade.RunGame()
	arcade.DrawScreen()

	println("Part two:", score)
}
