//nolint
package main

import (
	"fmt"
	"strconv"
	"strings"

	. "advent-of-code-2020/utils"
)

func main() {
	DownloadDayInput(2020, 23, false)
	input := ReadInputFileRelative()

	{
		q := NewCQueue(input, len(input))
		for i := 1; i <= 100; i++ {
			q.print()
			move(q)
		}
		q.print()
		q.inc(q.find(1))

		result := ""
		for i := 1; i < len(q.elems); i++ {
			result += strconv.Itoa(q.peak(i))
		}
		AssertEqual("97342568", result)
		println("Part 1:", result)
	}

	{
		q := NewCQueue(input, 1000000)
		for i := len(q.elems); i < 1000000; i++ {
			q.elems = append(q.elems, i+1)
		}
		for i := 1; i <= 100000; i++ {
			if i%1000 == 0 {
				fmt.Printf("%d %%\n", i*100/10000000)
			}
			move(q)
		}
		one := q.find(1)
		result := q.peak(one+1) * q.peak(one+2)

		println("Part 2:", result)
	}
}

func move(q *CQueue) {
	c1 := q.pick(1)
	c2 := q.pick(1)
	c3 := q.pick(1)

	destLbl := q.peak(0)
	for {
		destLbl--
		if destLbl < 0 {
			destLbl = Max(q.elems...)
		}
		if q.find(destLbl) != -1 {
			break // we found the label
		}
	}
	pos := q.find(destLbl)
	q.put(pos+1, c3)
	q.put(pos+1, c2)
	q.put(pos+1, c1)
	q.inc(1)
}

type CQueue struct {
	curr  int
	elems []int
}

func NewCQueue(input string, cap int) *CQueue {
	q := &CQueue{
		elems: make([]int, 0, cap),
	}
	for _, c := range strings.TrimSpace(input) {
		q.elems = append(q.elems, ParseInt(string(c)))
	}
	return q
}

func (q *CQueue) inc(i int) {
	q.curr = q.pos(i)
}

func (q *CQueue) find(elem int) int {
	for i, e := range q.elems {
		if e == elem {
			return Mod(i-q.curr, len(q.elems))
		}
	}
	return -1
}

func (q *CQueue) pos(i int) int {
	return (q.curr + i) % len(q.elems)
}

func (q *CQueue) peak(i int) int {
	return q.elems[q.pos(i)]
}

func (q *CQueue) pick(i int) int {
	pos := q.pos(i)
	e := q.elems[pos]
	if pos < q.curr {
		q.curr--
	}
	copy(q.elems[pos:], q.elems[pos+1:])
	q.elems = q.elems[:len(q.elems)-1]
	return e
}

func (q *CQueue) put(i, elem int) {
	pos := q.pos(i)
	if pos <= q.curr {
		q.curr++
	}
	q.elems = append(q.elems, 0)
	copy(q.elems[pos+1:], q.elems[pos:])
	q.elems[pos] = elem
}

func (q *CQueue) print() {
	sb := strings.Builder{}
	for i, e := range q.elems {
		if q.curr == i {
			sb.WriteString("(")
		}
		sb.WriteString(strconv.Itoa(e))
		if q.curr == i {
			sb.WriteString(")")
		}
		sb.WriteString(" ")
	}
	fmt.Println(sb.String())
}
