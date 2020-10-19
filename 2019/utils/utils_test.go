package utils

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func Test_fuelNeeded(t *testing.T) {
	tests := []struct {
		mass int
		want int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.mass), func(t *testing.T) {
			if got := FuelNeeded(tt.mass); got != tt.want {
				t.Errorf("fuelNeeded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFuelNeededIncludingFuelMass(t *testing.T) {
	tests := []struct {
		mass int
		want int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.mass), func(t *testing.T) {
			if got := FuelNeededIncludingFuelMass(tt.mass); got != tt.want {
				t.Errorf("FuelNeededIncludingFuelMass() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter(t *testing.T) {
	tests := []struct {
		start, end int
		expected   []int
	}{
		{0, 4, []int{0, 1, 2, 3}},
		{4, 0, []int{4, 3, 2, 1}},
		{-4, 0, []int{-4, -3, -2, -1}},
		{0, -4, []int{0, -1, -2, -3}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("[%d,%d]", tt.start, tt.end), func(t *testing.T) {
			var ints []int
			Iter(tt.start, tt.end, func(i int) {
				ints = append(ints, i)
			})
			require.Equal(t, tt.expected, ints)
		})
	}
}
