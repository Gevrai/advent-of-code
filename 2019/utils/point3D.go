package utils

import (
	"fmt"
	"regexp"
	"strconv"
)

type Point3D struct {
	X, Y, Z int64
}

//const floatRegexp = `-?\d*.?\d*`
const intRegexp = `-?\d*.?\d*`

var regexpInput = regexp.MustCompile(`<\W*x=(` + intRegexp + `),\W*y=(` + intRegexp + `),\W*z=(` + intRegexp + `)\W*>`)

func NewPointFromInput(input string) (p Point3D, err error) {
	groups := regexpInput.FindAllStringSubmatch(input, 1)
	if len(groups) != 1 || len(groups[0]) != 4 {
		return p, fmt.Errorf("invalid input %q", input)
	}
	for i, f := range []*int64{&p.X, &p.Y, &p.Z} {
		*f, err = strconv.ParseInt(groups[0][i+1], 10, 64)
		if err != nil {
			return p, fmt.Errorf("invalid float for argument %d: %v", i, err)
		}
	}
	return
}
