package computer

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_getMode(t *testing.T) {
	tests := []struct {
		code int
		pos  int
		want ParameterMode
	}{
		{10100, 1, 1},
		{10100, 2, 0},
		{10100, 3, 1},
		{10100, 4, 0},
		{10100, 5, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d pos %d", tt.code, tt.pos), func(t *testing.T) {
			require.Equal(t, tt.want, getMode(tt.code, tt.pos))
		})
	}
}
