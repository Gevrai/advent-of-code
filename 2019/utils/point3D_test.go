package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPointFromInput(t *testing.T) {
	tests := []struct {
		input     string
		wantPoint Point3D
	}{
		{"<x=-1, y=0, z=2>", Point3D{-1, 0, 2}},
		{"<x=2, y=-10, z=-7>", Point3D{2, -10, -7}},
		{"<x=4, y=-8, z=8>", Point3D{4, -8, 8}},
		{"<x=3, y=5, z=-1.5>", Point3D{3, 5, -1}},
		{"<x=3,     y=5, z=-1>", Point3D{3, 5, -1}},
		{"<    x=3, y=5, z=-1>", Point3D{3, 5, -1}},
		{"<x=3, y=5, z=-1>", Point3D{3, 5, -1}},
		{"<x=3, y=5, z=-1      >", Point3D{3, 5, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			require.Equal(t, tt.wantPoint, NewPointFromInput(tt.input))
		})
	}
}
