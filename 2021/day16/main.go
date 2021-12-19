package main

import (
	"fmt"
	"strings"

	. "advent-of-code-2021/utils"
)

type Stream struct {
	input string
	curr  int
}

func NewStreamFromHexadecimal(input string) *Stream {
	sb := strings.Builder{}
	for _, c := range strings.TrimSpace(input) {
		_, err := fmt.Fprintf(&sb, "%04b", ParseInt(string(c), 16))
		PanicIfError(err)
	}
	return &Stream{
		input: sb.String(),
		curr:  0,
	}
}

func (s *Stream) HasNext() bool {
	return s.curr < len(s.input)
}

func (s *Stream) Rx(n int) string {
	s.curr += n
	if s.curr > len(s.input) {
		panic(fmt.Sprintf("can't read %d", n))
	}
	return s.input[s.curr-n : s.curr]
}

func (s *Stream) RxInt(n int) int {
	rx := s.Rx(n)
	i := ParseInt(rx, 2)
	return i
}

func (s *Stream) RxLiteral() int {
	sb := strings.Builder{}
	for {
		next := s.Rx(5)
		sb.WriteString(next[1:])
		if next[0] != '1' {
			break
		}
	}
	return ParseInt(sb.String(), 2)
}

type Packet struct {
	version int
	typeID  int

	literal    int
	subpackets []*Packet
}

func main() {
	DownloadDayInput(2021, 16, false)
	input := ReadInputFileRelative()

	stream := NewStreamFromHexadecimal(input)

	packet := NewPacket(stream)
	println("Part 1:", packet.SumPacketVersions())
	println("Part 2:", packet.Eval())
}

func NewPacket(s *Stream) *Packet {
	packet := &Packet{}
	packet.version = s.RxInt(3)
	packet.typeID = s.RxInt(3)

	switch packet.typeID {
	case 4: // literal
		packet.literal = s.RxLiteral()
	default: // operator
		switch s.Rx(1) {
		case "0":
			subpacketLength := s.curr + s.RxInt(15)
			for s.curr < subpacketLength {
				packet.subpackets = append(packet.subpackets, NewPacket(s))
			}
		case "1":
			nbSubpackets := s.RxInt(11)
			for i := 0; i < nbSubpackets; i++ {
				packet.subpackets = append(packet.subpackets, NewPacket(s))
			}
		}
	}
	return packet
}

func (p *Packet) SumPacketVersions() int {
	count := p.version
	for _, s := range p.subpackets {
		count += s.SumPacketVersions()
	}
	return count
}

func (p *Packet) Eval() int {
	var subevals []int
	for _, subp := range p.subpackets {
		subevals = append(subevals, subp.Eval())
	}
	switch p.typeID {
	case 0: // sum
		return Sum(subevals)
	case 1: // product
		return Mult(subevals)
	case 2: // minimum
		return Min(subevals...)
	case 3: // maximum
		return Max(subevals...)
	case 4: // literal
		return p.literal
	case 5: // greater-than
		if subevals[0] > subevals[1] {
			return 1
		}
		return 0
	case 6: // less-than
		if subevals[0] < subevals[1] {
			return 1
		}
		return 0
	case 7: // equal
		if subevals[0] == subevals[1] {
			return 1
		}
		return 0
	default:
		panic(p.typeID)
	}
}
