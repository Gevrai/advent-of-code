package utils

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"os/exec"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

func ReadInputFileRelative(filename ...string) string {
	if len(filename) == 0 {
		filename = []string{inputFile}
	}
	// Relative path to where function is defined
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("ReadInputFile caller returned not ok")
	}
	dir, _ := path.Split(file)
	inputFile := path.Join(dir, filename[0])
	content, err := ioutil.ReadFile(inputFile)
	PanicIfError(err)
	return string(content)
}

func SplitNewLine(content string) []string {
	lines := strings.Split(content, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		return lines[:len(lines)-1]
	}
	return lines
}

func SplitEmptySlice(content []string) (parts [][]string) {
	parts = append(parts, []string{})
	for _, l := range content {
		if strings.TrimSpace(l) == "" {
			parts = append(parts, []string{})
			continue
		} else {
			parts[len(parts)-1] = append(parts[len(parts)-1], l)
		}
	}
	return parts
}

func DownloadDayInput(year, day int, force bool) {
	// Relative path to where function is defined
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("ReadInputFile caller returned not ok")
	}
	dir, _ := path.Split(file)
	inputFile := path.Join(dir, "input.txt")

	if _, err := os.Stat(inputFile); force || os.IsNotExist(err) {
		out, err := exec.Command("curl",
			fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day),
			"--cookie", fmt.Sprintf("session=%s", ReadCookie()),
		).Output()
		PanicIfError(err)
		PanicIfError(ioutil.WriteFile(inputFile, out, os.ModePerm))
	}
}

func ReadCookie() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("ReadCookie caller returned not ok")
	}
	dir, _ := path.Split(file)
	cookieFile := path.Join(dir, "cookie")

	content, err := ioutil.ReadFile(cookieFile)
	PanicIfError(err)
	return string(content)
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func AssertEqual(a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		panic(fmt.Sprintf("%v != %v", a, b))
	}
}

func ParseInt(s string) int {
	return ParseIntBase(s, 10)
}

func ParseIntBase(s string, base int) int {
	i, err := strconv.ParseInt(strings.TrimSpace(s), base, 64)
	PanicIfError(err)
	return int(i)
}

func Map[T, S any, TL ~[]T](list TL, mapper func(T) S) []S {
	if list == nil {
		return nil
	}
	out := make([]S, len(list))
	for i := range list {
		out[i] = mapper(list[i])
	}
	return out
}

func Trims(input string, trims ...string) string {
	for _, t := range trims {
		input = strings.Trim(input, t)
	}
	return input
}

func ReverseSlice(slice interface{}) {
	if slice == nil {
		return
	}
	swap := reflect.Swapper(slice)
	l := reflect.ValueOf(slice).Len()
	for i := 0; i < l/2; i++ {
		swap(i, l-i-1)
	}
}

func Reverse(input string) string {
	out := make([]byte, len(input))
	for i := range input {
		out[len(input)-1-i] = input[i]
	}
	return string(out)
}

func Mod(a, b int) int {
	return (a%b + b) % b
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

func Abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func Min(l ...int) int {
	min := math.MaxInt64
	for _, i := range l {
		if i < min {
			min = i
		}
	}
	return min
}

func Max(l ...int) int {
	max := math.MinInt64
	for _, i := range l {
		if i > max {
			max = i
		}
	}
	return max
}

func Count(inputs []string, condition func(string) bool) int {
	count := 0
	for _, s := range inputs {
		if condition(s) {
			count++
		}
	}
	return count
}

func Filter(inputs []string, condition func(string) bool) []string {
	var output []string
	for _, s := range inputs {
		if condition(s) {
			output = append(output, s)
		}
	}
	return output
}

func TransposeSquare(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func Sum(list []int) (count int) {
	for _, i := range list {
		count += i
	}
	return count
}

func Mult(list []int) (count int) {
	if len(list) == 0 {
		return 0
	}
	count = 1
	for _, i := range list {
		count *= i
	}
	return count
}

func Must[T any](t T, err error) T {
	PanicIfError(err)
	return t
}
