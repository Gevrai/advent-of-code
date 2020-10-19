package utils

import (
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
