package main

import (
	"time"

	"advent-of-code-2019/computer"
	. "advent-of-code-2019/utils"
)

type Arcade struct {
	cp           computer.Computer
	screen       Screen
	currentScore int64

	ball   Point
	paddle Point
}

func (a *Arcade) RunGame() (score int64) {
	done := a.cp.Run()

	a.cp.OnExpectInput(func() {
		time.Sleep(1 * time.Millisecond)
		a.DrawScreen()
		input := 0
		if a.paddle.X < a.ball.X {
			input = 1
			a.paddle.X += 1
		}
		if a.paddle.X > a.ball.X {
			input = -1
			a.paddle.X -= 1
		}
		err := a.cp.Input(BigInt(input))
		if err != nil {
			println("couldn't input joystick:", err)
		}
	})
	for {
		select {
		case <-done:
			// Game over
			return a.currentScore
		case x := <-a.cp.Output():
			y := <-a.cp.Output()
			param := <-a.cp.Output()

			if x.Int64() == -1 && y.Int64() == 0 {
				// Won the game !
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
				switch tile {
				case Ball:
					a.ball = p
				case HorizontalPaddle:
					a.paddle = p
				}
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
