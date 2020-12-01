package computer

import (
	"fmt"
	"math/big"
	"strings"
)

var DEBUG = false

type Computer interface {
	Run() (doneSignal chan big.Int)
	Put(pc int, newValue big.Int)
	Output() chan big.Int
	Input(big.Int) error
	OnExpectInput(func())
}

func InitComputer(input string) (Computer, error) {
	strip, err := intcodeFromInput(input)
	if err != nil {
		return nil, err
	}
	return &computer{
		instrPointer: 0,
		intcode:      strip,
		relativeBase: 0,
		input:        make(chan big.Int, 2),
		output:       make(chan big.Int, 2),
	}, nil
}

func intcodeFromInput(input string) (intcode Intcode, err error) {
	for _, elem := range strings.Split(input, ",") {
		i, ok := (&big.Int{}).SetString(elem, 10)
		if !ok || i == nil {
			return nil, fmt.Errorf("%s is not a valid big Int representation", elem)
		}
		intcode = append(intcode, *i)
	}
	return
}
