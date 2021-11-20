package main

import (
	"strings"

	. "advent-of-code-2020/utils"
)

func main() {
	DownloadDayInput(2020, 20, false)
	input := SplitNewLine(ReadInputFileRelative())

	var tiles []Tile
	for _, l := range strings.Split(strings.Join(input, " "), "  ") {
		tiles = append(tiles, NewTile(l))
	}

	total := 1
	for _, t := range findCorners(tiles) {
		total *= t.id
	}
	AssertEqual(total, 8272903687921)
	println("Part 1:", total)
}

func findCorners(tiles []Tile) (corners []Tile) {
	for _, t1 := range tiles {
		count := 0
		for _, t2 := range tiles {
			if t1.id == t2.id {
				continue
			}
			if t1.hasMatchingEdge(t2) {
				count++
			}
		}
		if count == 2 {
			corners = append(corners, t1)
		}
	}
	return corners
}

type Tile struct {
	id   int
	tile [][]byte
}

type Face int

const (
	North Face = iota
	South
	East
	West
	None
)

func NewTile(input string) (tile Tile) {
	parts := strings.Split(strings.TrimSpace(input), ":")
	tile.id = ParseInt(strings.Trim(parts[0], "Tile "))
	for _, r := range strings.Split(strings.TrimSpace(parts[1]), " ") {
		tile.tile = append(tile.tile, []byte(r))
	}
	return
}

func (t Tile) getEdges() map[Face][]byte {
	var west, east []byte
	for _, r := range t.tile {
		west = append(west, r[0])
		east = append(east, r[len(r)-1])
	}
	return map[Face][]byte{
		North: t.tile[0],
		South: t.tile[len(t.tile)-1],
		West:  west,
		East:  east,
	}
}

func (t Tile) hasMatchingEdge(t2 Tile) bool {
	for _, e1 := range t.getEdges() {
		if t2.matchingEdge(e1) != None {
			return true
		}
	}
	return false
}

func (t Tile) matchingEdge(edge []byte) Face {
	for face, e := range t.getEdges() {
		if string(edge) == string(e) || string(edge) == Reverse(string(e)) {
			return face
		}
	}
	return None
}
