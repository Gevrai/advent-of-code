package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"advent-of-code-2020/utils"
	"github.com/stretchr/testify/require"
)

func TestTile_rotate(t *testing.T) {
	tests := []struct {
		name                      string
		base, once, twice, thrice string
	}{
		{
			name:   "single",
			base:   "#",
			once:   "#",
			twice:  "#",
			thrice: "#",
		}, {
			name: "2x2",
			base: `#.
				   ..
`,
			once: `.#
				   ..
`,
			twice: `..
				    .#
`,
			thrice: `..
				     #.
`,
		}, {
			name: "3x3",
			base: `#..
				   ..#
				   ...
`,
			once: `..#
				   ...
				   .#.
`,
			twice: `...
				    #..
				    ..#
`,
			thrice: `.#.
				     ...
				     #..
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			testTile := createTile(tt.base)

			rotate(testTile.tile, 0)
			requireEqualTile(t, tt.base, testTile)

			rotate(testTile.tile, 1)
			requireEqualTile(t, tt.once, testTile)
			rotate(testTile.tile, 1)
			requireEqualTile(t, tt.twice, testTile)
			rotate(testTile.tile, 1)
			requireEqualTile(t, tt.thrice, testTile)
			rotate(testTile.tile, 1)
			requireEqualTile(t, tt.base, testTile)

			rotate(testTile.tile, 2)
			requireEqualTile(t, tt.twice, testTile)
			rotate(testTile.tile, 2)
			requireEqualTile(t, tt.base, testTile)

			rotate(testTile.tile, 3)
			requireEqualTile(t, tt.thrice, testTile)
			rotate(testTile.tile, 3)
			requireEqualTile(t, tt.twice, testTile)
			rotate(testTile.tile, 3)
			requireEqualTile(t, tt.once, testTile)
			rotate(testTile.tile, 3)
			requireEqualTile(t, tt.base, testTile)

			rotate(testTile.tile, -1)
			requireEqualTile(t, tt.thrice, testTile)
			rotate(testTile.tile, -3)
			requireEqualTile(t, tt.base, testTile)
		})
	}
}

func createTile(input string) (tile Tile) {
	input = strings.TrimSpace(input)
	for _, row := range strings.Split(input, "\n") {
		tile.tile = append(tile.tile, []byte(strings.TrimSpace(row)))
	}
	return tile
}

func requireEqualTile(t *testing.T, base string, tile Tile) {
	require.Equal(t, createTile(base), tile)
}

func Test_makeImage(t *testing.T) {

	input := utils.SplitNewLine(example)

	tiles := map[int]Tile{}
	for _, l := range strings.Split(strings.Join(input, " "), "  ") {
		t := NewTile(l)
		tiles[t.id] = t
	}

	assembled := assembleImage(tiles)
	assembled[0][0].print()
	assembled[0][1].print()
	assembled[0][2].print()
	image := makeImage(assembled)
	fmt.Println(formatMatrix(image))

	// Test any possible orientation matches
	for i := 0; i < 4; i++ {
		expectedImage := createTile(expectedImage).tile
		rotate(expectedImage, i)
		if reflect.DeepEqual(image, expectedImage) {
			return
		}

		flip(North, expectedImage)
		if reflect.DeepEqual(image, expectedImage) {
			return
		}
		flip(North, expectedImage)

		flip(East, expectedImage)
		if reflect.DeepEqual(image, expectedImage) {
			return
		}
	}

	require.FailNow(t, "no orientation matches")
}

const example = `
Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...
`

const expectedImage = `
.#.#..#.##...#.##..#####
###....#.#....#..#......
##.##.###.#.#..######...
###.#####...#.#####.#..#
##.#....#.##.####...#.##
...########.#....#####.#
....#..#...##..#.#.###..
.####...#..#.....#......
#..#.##..#..###.#.##....
#.####..#.####.#.#.###..
###.#.#...#.######.#..##
#.####....##..########.#
##..##.#...#...#.#.#.#..
...#..#..#.#.##..###.###
.#.#....#.##.#...###.##.
###.#...#..#.##.######..
.#.#.###.##.##.#..#.##..
.####.###.#...###.#..#.#
..#.#..#..#.#.#.####.###
#..####...#.#.#.###.###.
#####..#####...###....##
#.##..#..#...#..####...#
.#.###..##..##..####.##.
...###...##...#...#..###`
