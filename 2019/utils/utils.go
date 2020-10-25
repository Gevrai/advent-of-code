package utils

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"path"
	"reflect"
	"runtime"
	"strings"
)

func ReadInputFileRelative() (line []string) {
	// Relative path to where function is defined
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("ReadInputFile caller returned not ok")
	}
	dir, _ := path.Split(file)
	inputFile := path.Join(dir, "input.txt")
	return ReadInputFile(inputFile)
}

func ReadInputFile(filePath string) (lines []string) {

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err.Error())
	}

	for _, untrimmedLine := range strings.Split(string(content), "\n") {
		line := strings.TrimSpace(untrimmedLine)
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

func FuelNeeded(mass int) int {
	return mass/3 - 2
}

func FuelNeededIncludingFuelMass(mass int) int {
	fuel := FuelNeeded(mass)
	if fuel <= 0 {
		return 0
	}
	return fuel + FuelNeededIncludingFuelMass(fuel)
}

func BigInt(i interface{}) big.Int {
	if reflect.ValueOf(i).Kind() == reflect.Ptr {
		return BigInt(reflect.ValueOf(i).Elem().Interface())
	}
	switch i := i.(type) {
	case int:
		return *big.NewInt(int64(i))
	case int64:
		return *big.NewInt(int64(i))
	case uint64:
		return *big.NewInt(int64(i))
	default:
		panic(fmt.Errorf("could not convert type %s to big.Int", reflect.TypeOf(i).Name()))
	}
}

func GCD(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return (a / GCD(a, b)) * b
}

func AbsInt64(i int64) int64 {
	if i < 0 {
		return -i
	} else {
		return i
	}
}
