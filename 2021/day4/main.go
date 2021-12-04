package main

import (
	"strings"

	. "advent-of-code-2021/utils"
)

var id = 0

func main() {
	DownloadDayInput(2021, 4, false)
	input := SplitEmptySlice(SplitNewLine(ReadInputFileRelative()))

	var draws []int
	for _, n := range strings.Split(input[0][0], ",") {
		draws = append(draws, ParseInt(n))
	}

	boards := make([]BingoBoard, len(input[1:]))
	for i, l := range input[1:] {
		boards[i] = NewBingoBoard(l)
	}

	last, board := play(draws, boards)
	println("Part 1:", last*board.sumUnmarked())

	for len(boards) > 1 {
		_, board := play(draws, boards)
		// Remove it from list, and play again and again
		for i, b := range boards {
			if board.id == b.id {
				boards = append(boards[:i], boards[i+1:]...)
				break
			}
		}
	}
	last, board = play(draws, boards)
	println("Part 2:", last*board.sumUnmarked())
}

type BingoBoard struct {
	id     int
	board  [][]int
	marked [][]bool
}

func NewBingoBoard(input []string) (b BingoBoard) {
	b.board = make([][]int, len(input))
	b.marked = make([][]bool, len(input))
	for i, r := range input {
		b.board[i] = []int{}
		for _, n := range strings.Split(r, " ") {
			if n != "" {
				b.board[i] = append(b.board[i], ParseInt(n))
			}
		}
		b.marked[i] = make([]bool, len(b.board[i]))
	}
	TransposeSquare(b.board)
	b.forEach(func(i, j, n int) {
		if b.board[i][j] != n {
			panic("")
		}
	})
	b.id = id
	id++
	return b
}

func play(draws []int, boards []BingoBoard) (lastDraw int, board BingoBoard) {
	for _, draw := range draws {
		for _, b := range boards {
			b.mark(draw)
			if b.won() {
				return draw, b
			}
		}
	}
	panic("no winner")
}

func (b *BingoBoard) forEach(fn func(i, j, n int)) {
	for i, r := range b.board {
		for j, n := range r {
			fn(i, j, n)
		}
	}
}

func (b *BingoBoard) mark(number int) {
	b.forEach(func(i, j, n int) {
		if n == number {
			b.marked[i][j] = true
		}
	})
}

func (b *BingoBoard) sumUnmarked() (sum int) {
	b.forEach(func(i, j, n int) {
		if !b.marked[i][j] {
			sum += n
		}
	})
	return sum
}

func (b *BingoBoard) won() bool {
	// Rows
	for i := range b.marked {
		allMarked := true
		for j := range b.marked[i] {
			if !b.marked[i][j] {
				allMarked = false
				break
			}
		}
		if allMarked {
			return true
		}
	}
	// Columns
	for j := range b.marked[0] {
		allMarked := true
		for i := range b.marked {
			if !b.marked[i][j] {
				allMarked = false
				break
			}
		}
		if allMarked {
			return true
		}
	}
	return false
}
