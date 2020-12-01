package main

import (
	"advent-of-code-2019/computer"
	. "advent-of-code-2019/utils"
)

type Arcade struct {
	cp           computer.Computer
	screen       Screen
	currentScore int64

	ai ArcadeAI
}

func (a *Arcade) RunGame() (score int64) {
	done := a.cp.Run()

	a.cp.OnExpectInput(func() {
		joystickInput := a.ai.JoystickInput(a.screen)
		err := a.cp.Input(BigInt(joystickInput))
		if err != nil {
			println("couldn't input joystick:", err)
		}
	})
	for {
		select {
		case <-done:
			return a.currentScore
		case x := <-a.cp.Output():
			y := <-a.cp.Output()
			param := <-a.cp.Output()

			if x.Int64() == -1 && y.Int64() == 0 {
				// Show score
				a.currentScore = param.Int64()
				if a.CountTiles(Block) == 0 {
					return a.currentScore
				}
			} else {
				// Paint a tile
				p := Point{
					X: int(x.Int64()),
					Y: int(y.Int64()),
				}
				tile := Tile(param.Int64())
				a.screen.Paint(p, tile)
			}
		}
	}
}

func (a *Arcade) DrawScreen() {
	println(a.screen.Show())
}

func (a *Arcade) CountTiles(tile Tile) (nbTiles int) {
	return len(a.screen.FindTiles(tile))
}
