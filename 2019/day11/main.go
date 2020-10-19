package main

import (
	"advent-of-code-2019/computer"
	"advent-of-code-2019/robot"
	. "advent-of-code-2019/utils"
)

func main() {
	input := ReadInputFileRelative()

	cp, err := computer.InitComputer(input[0])
	if err != nil {
		panic(err)
	}
	bot := robot.NewRobot(cp)
	hull := NewMapHull()
	bot.SetPeeker(func(point Point) Color {
		return hull.Peek(point)
	})
	bot.SetPainter(func(point Point, color Color) {
		hull.Paint(point, color)
	})
	bot.Run()

	println("Part one:", len(hull.(MapHull)))
	println(hull.Show())

	cp, err = computer.InitComputer(input[0])
	if err != nil {
		panic(err)
	}
	bot = robot.NewRobot(cp)
	hull = NewMapHull()
	hull.Paint(Point{0, 0}, WHITE)
	bot.SetPeeker(func(point Point) Color {
		return hull.Peek(point)
	})
	bot.SetPainter(func(point Point, color Color) {
		hull.Paint(point, color)
	})
	bot.Run()

	println("Part two:")
	println(hull.Show())
}
