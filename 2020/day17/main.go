package main

import (
	"math"

	. "advent-of-code-2020/utils"
)

func main() {
	DownloadDayInput(2020, 17, false)
	input := SplitNewLine(ReadInputFileRelative())

	field := NewField(10)
	for i, l := range input {
		for j, r := range l {
			if r == '#' {
				field.spawn(j, i, 0)
			}
		}
	}

	for i := 0; i < 6; i++ {
		newField := NewField(len(field))
		field.foreach(func(x, y, z int) {
			if field.shouldLive(x, y, z) {
				newField.spawn(x, y, z)
			}
		})
		field = newField
	}

	println("Part 1:", len(field))

	{
		field := NewField4D(10)
		for i, l := range input {
			for j, r := range l {
				if r == '#' {
					field.spawn(j, i, 0, 0)
				}
			}
		}

		for i := 0; i < 6; i++ {
			newField := NewField4D(len(field))
			field.foreach(func(x, y, z, w int) {
				if field.shouldLive(x, y, z, w) {
					newField.spawn(x, y, z, w)
				}
			})
			field = newField
		}
		println("Part 2:", len(field))
	}
}

type Point struct {
	x, y, z int
}

type Point4D struct {
	x, y, z, w int
}

type Field map[Point]struct{}
type Field4D map[Point4D]struct{}

func NewField(len int) Field {
	return make(map[Point]struct{}, len)
}

func NewField4D(len int) Field4D {
	return make(map[Point4D]struct{}, len)
}

func (f Field) Print() {
	min, max := f.bounds()

	for z := min.z; z <= max.z; z++ {
		println()
		println("z=", z)
		for y := min.y; y <= max.y; y++ {
			for x := min.x; x <= max.x; x++ {
				if f.isAlive(x, y, z) {
					print("#")
				} else {
					print(".")
				}
			}
			println()
		}
	}

}

func (f Field) bounds() (min, max Point) {

	min.x = math.MaxInt64
	min.y = math.MaxInt64
	min.z = math.MaxInt64

	max.x = math.MinInt64
	max.y = math.MinInt64
	max.z = math.MinInt64

	for p := range f {
		if p.x < min.x {
			min.x = p.x
		} else if p.x > max.x {
			max.x = p.x
		}
		if p.y < min.y {
			min.y = p.y
		} else if p.y > max.y {
			max.y = p.y
		}
		if p.z < min.z {
			min.z = p.z
		} else if p.z > max.z {
			max.z = p.z
		}
	}
	return
}

func (f Field4D) bounds() (min, max Point4D) {

	min.x = math.MaxInt64
	min.y = math.MaxInt64
	min.z = math.MaxInt64
	min.w = math.MaxInt64

	max.x = math.MinInt64
	max.y = math.MinInt64
	max.z = math.MinInt64
	max.w = math.MinInt64

	for p := range f {
		if p.x < min.x {
			min.x = p.x
		} else if p.x > max.x {
			max.x = p.x
		}
		if p.y < min.y {
			min.y = p.y
		} else if p.y > max.y {
			max.y = p.y
		}
		if p.z < min.z {
			min.z = p.z
		} else if p.z > max.z {
			max.z = p.z
		}
		if p.w < min.w {
			min.w = p.w
		} else if p.w > max.w {
			max.w = p.w
		}
	}
	return
}

func (f Field) foreach(do func(x, y, z int)) {
	min, max := f.bounds()
	for i := min.x - 1; i <= max.x+1; i++ {
		for j := min.y - 1; j <= max.y+1; j++ {
			for k := min.z - 1; k <= max.z+1; k++ {
				do(i, j, k)
			}
		}
	}
}

func (f Field4D) foreach(do func(x, y, z, w int)) {
	min, max := f.bounds()
	for i := min.x - 1; i <= max.x+1; i++ {
		for j := min.y - 1; j <= max.y+1; j++ {
			for k := min.z - 1; k <= max.z+1; k++ {
				for l := min.w - 1; l <= max.w+1; l++ {
					do(i, j, k, l)
				}
			}
		}
	}
}

var neighbors Field
var neighbors4D Field4D

func init() {
	neighbors = NewField(27)
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				neighbors[Point{x, y, z}] = struct{}{}
			}
		}
	}
	delete(neighbors, Point{0, 0, 0})

	neighbors4D = NewField4D(81)
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					neighbors4D[Point4D{x, y, z, w}] = struct{}{}
				}
			}
		}
	}
	delete(neighbors4D, Point4D{0, 0, 0, 0})
}

func (f Field) isAlive(x, y, z int) bool {
	_, isAlive := f[Point{x, y, z}]
	return isAlive
}

func (f Field4D) isAlive(x, y, z, w int) bool {
	_, isAlive := f[Point4D{x, y, z, w}]
	return isAlive
}

func (f Field) spawn(x, y, z int) {
	f[Point{x, y, z}] = struct{}{}
}

func (f Field4D) spawn(x, y, z, w int) {
	f[Point4D{x, y, z, w}] = struct{}{}
}

func (f Field) shouldLive(x, y, z int) bool {

	countNeighbors := 0
	for n := range neighbors {
		if f.isAlive(x+n.x, y+n.y, z+n.z) {
			countNeighbors++
			if countNeighbors > 3 {
				return false // bust
			}
		}
	}

	if f.isAlive(x, y, z) {
		return countNeighbors == 2 || countNeighbors == 3
	} else {
		return countNeighbors == 3
	}
}

func (f Field4D) shouldLive(x, y, z, w int) bool {

	countNeighbors := 0
	for n := range neighbors4D {
		if f.isAlive(x+n.x, y+n.y, z+n.z, w+n.w) {
			countNeighbors++
			if countNeighbors > 3 {
				return false // bust
			}
		}
	}

	if f.isAlive(x, y, z, w) {
		return countNeighbors == 2 || countNeighbors == 3
	} else {
		return countNeighbors == 3
	}
}
