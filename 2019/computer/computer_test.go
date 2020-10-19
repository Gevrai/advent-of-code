package computer

import (
	"github.com/stretchr/testify/require"
	"math/big"
	"strings"
	"testing"
)

func TestStripFromInput(t *testing.T) {
	tests := []struct {
		strip string
		want  Intcode
	}{
		{"1,2", Intcode{*big.NewInt(1), *big.NewInt(2)}},
	}
	for _, tt := range tests {
		t.Run(tt.strip, func(t *testing.T) {
			got, err := intcodeFromInput(tt.strip)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_computer_Run(t *testing.T) {
	tests := []struct {
		name          string
		intcode       string
		expectedStrip string
		output        int
	}{
		// DAY 2 TESTING
		{
			name:          "HALT test",
			intcode:       "99",
			expectedStrip: "99",
			output:        99,
		},
		{
			name:          "ADD test",
			intcode:       "1,0,0,0,99",
			expectedStrip: "2,0,0,0,99",
			output:        2,
		},
		{
			name:          "MULT test",
			intcode:       "2,3,0,3,99",
			expectedStrip: "2,3,0,6,99",
			output:        2,
		},
		{
			name:          "MULT test",
			intcode:       "2,4,4,5,99,0",
			expectedStrip: "2,4,4,5,99,9801",
			output:        2,
		},
		{
			name:          "ADD, MULT test",
			intcode:       "1,1,1,4,99,5,6,0,99",
			expectedStrip: "30,1,1,4,2,5,6,0,99",
			output:        30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c, err := InitComputer(tt.intcode)
			require.NoError(t, err)

			doneSignal := c.Run()
			require.Equal(t, tt.output, <-doneSignal)

			if tt.expectedStrip != "" {
				expected, err := intcodeFromInput(tt.expectedStrip)
				require.NoError(t, err)
				require.Equal(t, expected, c.(*computer).intcode)
			}
		})
	}
}

func Test_computer_RunWithInputOutput(t *testing.T) {
	tests := []struct {
		name                 string
		intcode              string
		input                []int64
		expectedOutput       []int64
		expectedOutputBigInt string
	}{
		{
			name:           "jump, position mode, input 0",
			intcode:        "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			input:          []int64{0},
			expectedOutput: []int64{0},
		},
		{
			name:           "jump, position mode, input 1",
			intcode:        "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			input:          []int64{1},
			expectedOutput: []int64{1},
		},
		{
			name:           "jump, position mode, input -2",
			intcode:        "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			input:          []int64{-2},
			expectedOutput: []int64{1},
		},
		{
			name:           "jump, position mode, input 35",
			intcode:        "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			input:          []int64{35},
			expectedOutput: []int64{1},
		},
		{
			name:           "jump, immediate mode, input 0",
			intcode:        "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			input:          []int64{0},
			expectedOutput: []int64{0},
		},
		{
			name:           "jump, immediate mode, input 1",
			intcode:        "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			input:          []int64{1},
			expectedOutput: []int64{1},
		},
		{
			name:           "jump, immediate mode, input -2",
			intcode:        "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			input:          []int64{-2},
			expectedOutput: []int64{1},
		},
		{
			name:           "jump, immediate mode, input 35",
			intcode:        "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			input:          []int64{35},
			expectedOutput: []int64{1},
		},
		{
			name:           "jump compare test day 5, input 7",
			intcode:        "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			input:          []int64{7},
			expectedOutput: []int64{999},
		},
		{
			name:           "jump compare test day 5, input 8",
			intcode:        "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			input:          []int64{8},
			expectedOutput: []int64{1000},
		},
		{
			name:           "jump compare test day 5, input 9",
			intcode:        "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			input:          []int64{9},
			expectedOutput: []int64{1001},
		},
		{
			name:           "quine",
			intcode:        "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99",
			input:          []int64{},
			expectedOutput: []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			name:           "output big 16-digit number",
			intcode:        "1102,34915192,34915192,7,4,7,99,0",
			input:          []int64{},
			expectedOutput: []int64{1219070632396864},
		},
		{
			name:           "output middle number small",
			intcode:        "104,12345,99",
			input:          []int64{},
			expectedOutput: []int64{12345},
		},
		{
			name:           "output middle number medium",
			intcode:        "104,1125899906842624,99",
			input:          []int64{},
			expectedOutput: []int64{1125899906842624},
		},
		{
			name:                 "output middle number big",
			intcode:              "104,1125899906842624999999999999999,99",
			input:                []int64{},
			expectedOutputBigInt: "1125899906842624999999999999999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() { require.Nil(t, recover()) }()

			c, err := InitComputer(tt.intcode)
			require.NoError(t, err)

			DEBUG = true

			go func() {
				for _, val := range tt.input {
					err := c.Input(*big.NewInt(val))
					require.NoError(t, err)
				}
			}()

			c.Run()

			if len(tt.expectedOutput) > 0 {
				var output []int64
				for val := range c.Output() {
					output = append(output, val.Int64())
				}
				require.Equal(t, tt.expectedOutput, output)
			} else {
				require.NotEmpty(t, tt.expectedOutputBigInt)
				var expected []big.Int
				for _, n := range strings.Split(tt.expectedOutputBigInt, ",") {
					i, ok := (&big.Int{}).SetString(n, 10)
					require.True(t, ok)
					require.NotNil(t, i)
					expected = append(expected, *i)
				}
				var output []big.Int
				for val := range c.Output() {
					output = append(output, val)
				}
				require.Equal(t, expected, output)
			}
		})
	}
}
