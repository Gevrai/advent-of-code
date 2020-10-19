package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_object_CountOrbits(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"empty", []string{}, 0},
		{"one", []string{"COM)B"}, 1},
		{"two sequential", []string{"COM)B", "B)A"}, 3},
		{"two parallel", []string{"COM)B", "COM)A"}, 2},
		{"given example", []string{
			"COM)B",
			"B)C",
			"C)D",
			"D)E",
			"E)F",
			"B)G",
			"G)H",
			"D)I",
			"E)J",
			"J)K",
			"K)L",
		}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, objectMap := CreateObjectTree(tt.input)

			count := 0
			for _, object := range objectMap {
				count += object.NumberOfParents()
			}
			require.Equal(t, tt.want, count)
		})
	}
}
