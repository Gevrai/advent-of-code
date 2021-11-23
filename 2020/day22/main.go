package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	. "advent-of-code-2020/utils"
)

var playerRegexp = regexp.MustCompile(`Player (\d):`)

const example1 = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

const example2 = `Player 1:
43
19

Player 2:
2
29
14
`

func main() {
	DownloadDayInput(2020, 22, false)

	{
		input := SplitNewLine(ReadInputFileRelative())
		player1, player2 := createDecks(input)
		winner, points := playCombat(player1, player2)
		println(fmt.Sprintf("Part 1: Winner is player %d with %d points", winner, points))
	}

	{
		input := SplitNewLine(ReadInputFileRelative())
		player1, player2 := createDecks(input)
		winner, points := playRecursiveCombat(player1, player2)
		println(fmt.Sprintf("Part 2: Winner is player %d with %d points", winner, points))
	}
}

func playRecursiveCombat(player1, player2 Deck) (winner, points int) {
	states := map[string]bool{}
	for player1.len() > 0 && player2.len() > 0 {
		// Check previous states
		s := player1.state() + "-" + player2.state()
		if states[s] {
			return 1, player1.points()
		}
		states[s] = true

		// Play round
		c1 := player1.pop()
		c2 := player2.pop()
		var winner int
		switch {
		case c1 <= player1.len() && c2 <= player2.len():
			winner, _ = playRecursiveCombat(player1.copy(c1), player2.copy(c2))
		case c1 > c2:
			winner = 1
		case c2 > c1:
			winner = 2
		default:
			panic("equal cards...")
		}

		if winner == 1 {
			player1.bury(c1)
			player1.bury(c2)
		} else {
			player2.bury(c2)
			player2.bury(c1)
		}
	}

	if player2.len() == 0 {
		// Player 1 won!
		return 1, player1.points()
	} else {
		// Player 2 won!
		return 2, player2.points()
	}
}

func playCombat(player1, player2 Deck) (winner, points int) {
	for player1.len() > 0 && player2.len() > 0 {
		c1 := player1.pop()
		c2 := player2.pop()

		switch {
		case c1 > c2:
			player1.bury(c1)
			player1.bury(c2)
		case c2 > c1:
			player2.bury(c2)
			player2.bury(c1)
		default:
			panic("equal cards...")
		}
	}
	if player2.len() == 0 {
		// Player 1 won!
		return 1, player1.points()
	} else {
		// Player 2 won!
		return 2, player2.points()
	}
}

func createDecks(input []string) (player1, player2 Deck) {
	var current *Deck
	for _, l := range input {
		if l == "" {
			continue
		}
		if match := playerRegexp.FindStringSubmatch(l); match != nil {
			switch ParseInt(match[1]) {
			case 1:
				current = &player1
			case 2:
				current = &player2
			default:
				panic("too many players")
			}
			continue
		}
		current.bury(ParseInt(l))
	}
	return player1, player2
}

type Deck struct {
	cards []int
}

func (d *Deck) len() int {
	return len(d.cards)
}

func (d *Deck) pop() (card int) {
	card, d.cards = d.cards[0], d.cards[1:]
	return card
}

func (d *Deck) bury(card int) {
	d.cards = append(d.cards, card)
}

func (d *Deck) points() (points int) {
	ReverseSlice(d.cards) // lazyness...
	for i, card := range d.cards {
		points += (i + 1) * card
	}
	ReverseSlice(d.cards)
	return points
}

func (d *Deck) state() string {
	sb := strings.Builder{}
	for _, card := range d.cards {
		if sb.Len() > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(strconv.Itoa(card))
	}
	return sb.String()
}

func (d *Deck) copy(maxCards int) Deck {
	cards := make([]int, Min(len(d.cards), maxCards))
	copy(cards, d.cards)
	return Deck{
		cards: cards,
	}
}
