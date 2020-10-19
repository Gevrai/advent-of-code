package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_validPasswordPart1(t *testing.T) {
	tests := []struct {
		password string
		valid    bool
	}{
		{"111111", true},
		{"112345", true},
		{"123455", true},

		{"11111", false},
		{"223450", false},
		{"123789", false},
	}
	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			require.Equal(t, tt.valid, validPasswordPart1(tt.password))
		})
	}
}

func Test_validPasswordPart2(t *testing.T) {
	tests := []struct {
		password string
		valid    bool
	}{
		{"111122", true},
		{"111111", false},
		{"112345", true},
		{"123455", true},

		{"12344", false},
		{"123444", false},
		{"223450", false},
		{"123789", false},
	}
	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			require.Equal(t, tt.valid, validPasswordPart2(tt.password))
		})
	}
}
