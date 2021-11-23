package main

import (
	"fmt"
	"math"
	"strings"

	. "advent-of-code-2020/utils"
)

var SeaMonster = `
                  # 
#    ##    ##    ###
 #  #  #  #  #  #   
`

type Tile struct {
	id   int
	tile [][]byte
}

func NewTile(input string) (tile Tile) {
	parts := strings.Split(strings.TrimSpace(input), ":")
	tile.id = ParseInt(strings.Trim(parts[0], "Tile "))
	for _, r := range strings.Split(strings.TrimSpace(parts[1]), " ") {
		tile.tile = append(tile.tile, []byte(r))
	}
	tile.print()
	return
}

type Face int

const (
	North Face = iota
	East
	South
	West
)

func (f Face) String() string {
	return []string{"North", "East", "South", "West", "None"}[f]
}

type Coord struct{ x, y int }

func main() {
	DownloadDayInput(2020, 20, false)
	input := SplitNewLine(ReadInputFileRelative())

	tiles := map[int]Tile{}
	for _, l := range strings.Split(strings.Join(input, " "), "  ") {
		t := NewTile(l)
		tiles[t.id] = t
	}

	assembled := assembleImage(tiles)
	w, l := len(assembled), len(assembled[0])
	total := assembled[0][0].id *
		assembled[0][l-1].id *
		assembled[w-1][0].id *
		assembled[w-1][l-1].id
	AssertEqual(total, 8272903687921)
	println("Part 1:", total)

	seaMonster := seaMonsterCoords()
	nbSeaMonsters := countSeaMonsters(assembled, seaMonster)
	println("Part 2:", count(makeImage(assembled), '#')-nbSeaMonsters*len(seaMonster))
}

func countSeaMonsters(assembled [][]Tile, seaMonster []Coord) int {
	for i := 0; i < 4; i++ {
		image := makeImage(assembled)
		rotate(image, i)
		if count := findMonster(image, seaMonster); count > 0 {
			return count
		}
		flip(North, image)
		if count := findMonster(image, seaMonster); count > 0 {
			return count
		}
		flip(North, image)

		flip(East, image)
		if count := findMonster(image, seaMonster); count > 0 {
			return count
		}
	}
	return 0
}

func count(image [][]byte, b byte) (count int) {
	for _, row := range image {
		for _, c := range row {
			if b == c {
				count++
			}
		}
	}
	return count
}

func findMonster(image [][]byte, monster []Coord) (count int) {
	for i := range image {
		for j := range image[i] {
			if matchesMonster(image, Coord{i, j}, monster) {
				count++
			}
		}
	}
	return count
}

func matchesMonster(image [][]byte, startingCoord Coord, monster []Coord) bool {
	for _, m := range monster {
		x := startingCoord.x + m.x
		y := startingCoord.y + m.y
		if !inBounds(image, x, y) {
			return false
		}
		if image[x][y] != '#' {
			return false
		}
	}
	return true
}

func inBounds(image [][]byte, i, j int) bool {
	return len(image) > i && len(image[i]) > j
}

func seaMonsterCoords() (monster []Coord) {
	lines := strings.Split(SeaMonster, "\n")
	for j, l := range lines {
		if l != "" {
			for i, c := range l {
				if c == '#' {
					monster = append(monster, Coord{i, j})
				}
			}
		}
	}
	return monster
}

func makeImage(assembled [][]Tile) [][]byte {
	tileWidth, tileHeight := len(assembled[0][0].tile)-2, len(assembled[0][0].tile[0])-2

	image := make([][]byte, tileWidth*len(assembled))
	for i := range image {
		image[i] = make([]byte, tileHeight*len(assembled[0]))
	}

	for i := range image {
		for j := range image[i] {
			tile := assembled[i/tileWidth][j/tileHeight]
			image[i][j] = tile.tile[(j%tileHeight)+1][(i%tileWidth)+1] //eugh...
		}
	}
	return image
}

func assembleImage(tiles map[int]Tile) [][]Tile {

	start := len(tiles)
	last := 0

	// Start image with any tile at (0,0)
	image := map[Coord]Tile{}
	for id, tile := range tiles {
		image[Coord{0, 0}] = tile
		delete(tiles, id)
		break
	}

	for len(tiles) > 0 {
		if last == len(tiles) {
			panic("none found")
		}
		last = len(tiles)
		fmt.Printf("%d/%d\n", start-last, start)

		for coord, tile := range image {
			for face, edge := range tile.getEdges() {
				newLocation := coord.getCoord(face)
				if _, ok := image[newLocation]; ok {
					continue // there is already a tile here
				}
				// Try to find a tile that fits this face
				for id, candidate := range tiles {
					candidateFace, flipped := candidate.matchingEdge(edge)
					if candidateFace == nil {
						continue
					}
					// We found one! Position correctly
					if flipped {
						flip(*candidateFace, candidate.tile)
					}
					rotate(candidate.tile, int(face-(*candidateFace)+2)%4)
					// And place in image
					image[newLocation] = candidate
					delete(tiles, id)

					// Check if really match (face check matches opposite face now)
					AssertEqual(string(tile.getEdges()[face]), Reverse(string(candidate.getEdges()[(face+2)%4])))
					break
				}
			}
		}
	}

	min := Coord{math.MaxInt64, math.MaxInt64}
	max := Coord{math.MinInt64, math.MinInt64}
	for pos := range image {
		min.x = Min(min.x, pos.x)
		min.y = Min(min.y, pos.y)
		max.x = Max(max.x, pos.x)
		max.y = Max(max.y, pos.y)
	}

	tiledImage := make([][]Tile, max.x-min.x+1)
	for i := min.x; i <= max.x; i++ {
		tiledImage[i-min.x] = make([]Tile, max.y-min.y+1)
		for j := min.y; j <= max.y; j++ {
			tile, ok := image[Coord{i, j}]
			if !ok {
				panic(fmt.Sprintf("no tile at (%d,%d)", i, j))
			}
			// Double check all faces match correctly
			for face, edge := range tile.getEdges() {
				neigh, ok := image[Coord{i, j}.getCoord(face)]
				if !ok {
					continue
				}
				otherFace, flipped := neigh.matchingEdge(edge)
				AssertEqual(*otherFace, (face+2)%4)
				AssertEqual(flipped, false)
			}
			tiledImage[i-min.x][j-min.y] = tile
		}
	}
	return tiledImage
}

// returns edges oriented like
//    --->
//	  A	 |
//	  |	 v
//    <---
func (t Tile) getEdges() map[Face][]byte {
	var east, west []byte
	for _, r := range t.tile {
		east = append(east, r[len(r)-1])
		west = append(west, r[0])
	}
	return map[Face][]byte{
		North: t.tile[0],
		East:  east,
		South: reverseBytes(t.tile[len(t.tile)-1]),
		West:  reverseBytes(west),
	}
}

func (t Tile) matchingEdge(edge []byte) (face *Face, flipped bool) {
	reversedEdge := reverseBytes(edge)
	for face, e := range t.getEdges() {
		// match reverse means face is not flipped to match
		if string(reversedEdge) == string(e) {
			return &face, false
		}
		// exact match means face should be flipped to match
		if string(edge) == string(e) {
			return &face, true
		}
	}
	return nil, false
}

func rotate(matrix [][]byte, rotate int) {
	switch Mod(rotate, 4) {
	case 0:
		return
	case 1:
		transpose(matrix)
		flip(North, matrix)
	case 2:
		flip(North, matrix)
		flip(East, matrix)
	case 3:
		flip(North, matrix)
		transpose(matrix)
	default:
		panic("mod 4?")
	}
}

func transpose(matrix [][]byte) {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func flip(face Face, matrix [][]byte) {
	switch face {
	case North, South:
		for i := range matrix {
			matrix[i] = reverseBytes(matrix[i])
		}
	case East, West:
		for i := 0; i < len(matrix)/2; i++ {
			matrix[i], matrix[len(matrix)-i-1] = matrix[len(matrix)-i-1], matrix[i]
		}
	default:
		panic("euhh")
	}
}

func (t Tile) print() {
	fmt.Printf("Tile %d: \n%s\n", t.id, formatMatrix(t.tile))
}

func formatMatrix(tile [][]byte) string {
	sb := strings.Builder{}
	for _, row := range tile {
		sb.Write(row)
		sb.WriteString("\n")
	}
	return sb.String()
}

func (c Coord) getCoord(face Face) Coord {
	switch face {
	case North:
		return Coord{c.x, c.y - 1}
	case South:
		return Coord{c.x, c.y + 1}
	case West:
		return Coord{c.x - 1, c.y}
	case East:
		return Coord{c.x + 1, c.y}
	default:
		panic("euhhhh")
	}
}

func reverseBytes(b []byte) []byte {
	return []byte(Reverse(string(b)))
}
