package utils

import (
	"io/ioutil"
	"path"
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
