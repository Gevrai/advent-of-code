package main

import (
	"advent-of-code-2019/utils"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

const DEBUG = false

func TestBody_Update(t *testing.T) {
	tests := []struct {
		input  []Body
		steps  int
		energy int64
	}{
		{
			input: []Body{
				{pos: utils.Point3D{-1, 0, 2}},
				{pos: utils.Point3D{2, -10, -7}},
				{pos: utils.Point3D{4, -8, 8}},
				{pos: utils.Point3D{3, 5, -1}},
			},
			steps:  10,
			energy: 179,
		},
		{
			input: []Body{
				{pos: utils.Point3D{-8, -10, 0}},
				{pos: utils.Point3D{5, 5, 10}},
				{pos: utils.Point3D{2, -7, 3}},
				{pos: utils.Point3D{9, -8, -3}},
			},
			steps:  100,
			energy: 1940,
		},
		{
			input:  NewSystem(utils.ReadInputFileRelative()).bodies,
			steps:  1000,
			energy: 11384,
		},
		{
			input:  NewSystem(utils.ReadInputFileRelative()).bodies,
			steps:  10000,
			energy: 116132,
		},
		{
			input:  NewSystem(utils.ReadInputFileRelative()).bodies,
			steps:  1e7,
			energy: 195319,
		},
	}
	for _, tt := range tests {
		t.Run("example", func(t *testing.T) {
			system := System{bodies: tt.input}

			if DEBUG {
				println("Step ", 0)
				println(system.SPrint())
			}

			for i := 0; i < tt.steps; i++ {
				system.Update()
				if DEBUG {
					println("Step ", i+1)
					println(system.SPrint())
				}
			}
			require.Equal(t, tt.energy, system.Energy())
		})
	}
}

func TestBody_UpdateFast(t *testing.T) {
	tests := []struct {
		input  []Body
		steps  int
		energy int64
	}{
		{
			input: []Body{
				{pos: utils.Point3D{-1, 0, 2}},
				{pos: utils.Point3D{2, -10, -7}},
				{pos: utils.Point3D{4, -8, 8}},
				{pos: utils.Point3D{3, 5, -1}},
			},
			steps:  10,
			energy: 179,
		},
		{
			input: []Body{
				{pos: utils.Point3D{-8, -10, 0}},
				{pos: utils.Point3D{5, 5, 10}},
				{pos: utils.Point3D{2, -7, 3}},
				{pos: utils.Point3D{9, -8, -3}},
			},
			steps:  100,
			energy: 1940,
		},
		{
			input:  NewSystem(utils.ReadInputFileRelative()).bodies,
			steps:  1000,
			energy: 11384,
		},
		{
			input:  NewSystem(utils.ReadInputFileRelative()).bodies,
			steps:  1e7,
			energy: 195319,
		},
	}
	for _, tt := range tests {
		t.Run("example", func(t *testing.T) {
			system := System{bodies: tt.input}

			if DEBUG {
				println("Step ", 0)
				println(system.SPrint())
			}

			for i := 0; i < tt.steps; i++ {
				system.UpdateFast()
				if DEBUG {
					println("Step ", i+1)
					println(system.SPrint())
				}
			}
			require.Equal(t, tt.energy, system.Energy())
		})
	}
}

func BenchmarkBody_Update(b *testing.B) {

	system := NewSystem(utils.ReadInputFileRelative())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		system.Update()
	}
	b.StopTimer()

	println(fmt.Sprintf("Energy was %d after %d steps", system.Energy(), b.N))
}

func BenchmarkBody_UpdateFast(b *testing.B) {

	system := NewSystem(utils.ReadInputFileRelative())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		system.UpdateFast()
	}
	b.StopTimer()

	println(fmt.Sprintf("Energy was %d after %d steps", system.Energy(), b.N))
}
