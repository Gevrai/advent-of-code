package utils

import (
	"fmt"
	"regexp"
	"strconv"
)

type Point3D struct {
	X, Y, Z int64
}

const intGroup = `\s*(-?\d*)\s*`

var regexpInput = regexp.MustCompile(
	`<\s*x=` + intGroup +
		`,\s*y=` + intGroup +
		`,\s*z=` + intGroup +
		`>`)

func NewPointFromInput(input string) (p Point3D) {
	groups := regexpInput.FindAllStringSubmatch(input, 1)
	if len(groups) != 1 || len(groups[0]) != 4 {
		panic(fmt.Errorf("invalid input %q", input))
	}
	var err error
	for i, f := range []*int64{&p.X, &p.Y, &p.Z} {
		*f, err = strconv.ParseInt(groups[0][i+1], 10, 64)
		if err != nil {
			panic(fmt.Errorf("invalid float for argument %d: %v", i, err))
		}
	}
	return
}

func (p Point3D) AsVector() Vector3D {
	return Vector3D{
		X: p.X,
		Y: p.Y,
		Z: p.Z,
	}
}
