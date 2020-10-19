package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestSpaceMap_BestLocation(t *testing.T) {
	tests := []struct {
		input        []string
		bestLocation Point
		asteroids    int
	}{
		{
			input: []string{
				"###",
				"###",
				"###",
			},
			bestLocation: Point{1, 1},
			asteroids:    8,
		},
		{
			input: []string{
				".#..#",
				".....",
				"#####",
				"....#",
				"...##",
			},
			bestLocation: Point{3, 4},
			asteroids:    8,
		},
		{
			input: []string{
				"......#.#.",
				"#..#.#....",
				"..#######.",
				".#.#.###..",
				".#..#.....",
				"..#....#.#",
				"#..#....#.",
				".##.#..###",
				"##...#..#.",
				".#....####",
			},
			bestLocation: Point{5, 8},
			asteroids:    33,
		},
		{
			input: []string{
				"#.#...#.#.",
				".###....#.",
				".#....#...",
				"##.#.#.#.#",
				"....#.#.#.",
				".##..###.#",
				"..#...##..",
				"..##....##",
				"......#...",
				".####.###.",
			},
			bestLocation: Point{1, 2},
			asteroids:    35,
		},
		{
			input: []string{
				".#..#..###",
				"####.###.#",
				"....###.#.",
				"..###.##.#",
				"##.##.#.#.",
				"....###..#",
				"..#.#..#.#",
				"#..#.#.###",
				".##...##.#",
				".....#.#..",
			},
			bestLocation: Point{6, 3},
			asteroids:    41,
		},
		{
			input: []string{
				".#..##.###...#######",
				"##.############..##.",
				".#.######.########.#",
				".###.#######.####.#.",
				"#####.##.#.##.###.##",
				"..#####..#.#########",
				"####################",
				"#.####....###.#.#.##",
				"##.#################",
				"#####.##.###..####..",
				"..######..##.#######",
				"####.##.####...##..#",
				".#####..#.######.###",
				"##...#.##########...",
				"#.##########.#######",
				".####.#.###.###.#.##",
				"....##.##.###..#####",
				".#.#.###########.###",
				"#.#.#.#####.####.###",
				"###.##.####.##.#..##",
			},
			bestLocation: Point{11, 13},
			asteroids:    210,
		},
	}
	for _, tt := range tests {
		t.Run("example", func(t *testing.T) {
			sm := NewSpaceMap(tt.input)
			gotLocation, gotAsteroidsInSight := sm.BestLocation()
			require.Equal(t, gotLocation, tt.bestLocation)
			require.Equal(t, gotAsteroidsInSight, tt.asteroids)
		})
	}
}

func TestDirection_AngleWith(t *testing.T) {
	tests := []struct {
		a, b  Direction
		angle float64
	}{
		{Direction{1, 1}, Direction{1, 1}, 0},
		{Direction{0, 1}, Direction{1, 1}, 45},
		{Direction{0, 1}, Direction{1, 0}, 90},
		{Direction{0, 1}, Direction{-1, 0}, 90},
	}
	for _, tt := range tests {
		t.Run(spew.Sprintf("%v, %v => %f", tt.a, tt.b, tt.angle), func(t *testing.T) {
			got := tt.a.AngleWith(tt.b)
			require.InDelta(t, tt.angle, got*(180/math.Pi), 1e-5)

			inverse := tt.b.AngleWith(tt.a)
			require.InDelta(t, got, inverse, 1e-6)
		})
	}
}

func TestSpaceMap_VaporizeAsteroids(t *testing.T) {

	input := []string{
		".#..##.###...#######",
		"##.############..##.",
		".#.######.########.#",
		".###.#######.####.#.",
		"#####.##.#.##.###.##",
		"..#####..#.#########",
		"####################",
		"#.####....###.#.#.##",
		"##.#################",
		"#####.##.###..####..",
		"..######..##.#######",
		"####.##.####...##..#",
		".#####..#.######.###",
		"##...#.##########...",
		"#.##########.#######",
		".####.#.###.###.#.##",
		"....##.##.###..#####",
		".#.#.###########.###",
		"#.#.#.#####.####.###",
		"###.##.####.##.#..##",
	}

	sm := NewSpaceMap(input)

	src := Point{11, 13}
	order := sm.VaporizeAsteroids(src)

	for i := range sm {
		for j := range sm[i] {
			if (Point{i, j}) != src {
				require.Equal(t, string('.'), string(sm.Get(Point{i, j})))
			}
		}
	}

	for pos, point := range map[int]Point{
		1:   {11, 12},
		2:   {12, 1},
		3:   {12, 2},
		10:  {12, 8},
		20:  {16, 0},
		50:  {16, 9},
		100: {10, 16},
		199: {9, 6},
		200: {8, 2},
		201: {10, 9},
		299: {11, 1},
	} {
		require.Equal(t, point, order[pos-1], "wrong %d", pos)
	}
}
