package main

import . "advent-of-code-2019/utils"

type ArcadeAI struct {
	previousBall *Point
	nextPaddle *Point
}

func (ai *ArcadeAI) JoystickInput(screen Screen) (input int64) {

	paddles := screen.FindTiles(HorizontalPaddle)
	if len(paddles) == 0 {
		return 0
	}
	paddle := paddles[0]

	balls := screen.FindTiles(Ball)
	if len(balls) == 0 {
		return 0
	}
	newBall := balls[0]

	if ai.nextPaddle != nil {
		if paddle.X < expectedCollision.X {
			input = 1
		}
		if paddle.X > expectedCollision.X {
			input = -1
		}
	}

	if ai.previousBall != nil {
		v := ai.previousBall.DirectionTo(newBall)
		nextPosition :=
		if newBall.Add(v)
		println(v.String())
		expectedCollision := ai.CollidesAtY(paddle.Y)
		if paddle.X < expectedCollision.X {
			input = 1
		}
		if paddle.X > expectedCollision.X {
			input = -1
		}
	}

	ai.previousBall = &newBall
	return input
}

func (ai *ArcadeAI) CollidesAtY(y int) Point {
	return Point{0, 0}
}
