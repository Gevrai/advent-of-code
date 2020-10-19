package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPointFromInput(t *testing.T) {
	tests := []struct {
		input     string
		wantPoint Point3D
		wantErr   bool
	}{
		{"<x=-1, y=0, z=2>", Point3D{-1, 0, 2}, false},
		{"<x=2, y=-10, z=-7>", Point3D{2, -10, -7}, false},
		{"<x=4, y=-8, z=8>", Point3D{4, -8, 8}, false},
		{"<x=3, y=5, z=-1.5>", Point3D{3, 5, -1}, false},
		{"<x=3,     y=5, z=-1>", Point3D{3, 5, -1}, false},
		{"<    x=3, y=5, z=-1>", Point3D{3, 5, -1}, false},
		{"<x=3, y=5, z=-1>", Point3D{3, 5, -1}, false},
		{"<x=3, y=5, z=-1      >", Point3D{3, 5, -1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			gotPoint, err := NewPointFromInput(tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.wantPoint, gotPoint)
		})
	}
}
