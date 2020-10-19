package computer

import (
	"fmt"
	"math"
	"math/big"
)

type Intcode []big.Int

type Opcode int

const (
	Add             Opcode = 1
	Mult            Opcode = 2
	Input           Opcode = 3
	Output          Opcode = 4
	JmpIfTrue       Opcode = 5
	JmpIfFalse      Opcode = 6
	LessThan        Opcode = 7
	Equals          Opcode = 8
	ModRelativeBase Opcode = 9
	Halt            Opcode = 99
)

func (o Opcode) String() string {
	switch o {
	case Add:
		return "Add"
	case Mult:
		return "Mult"
	case Input:
		return "Input"
	case Output:
		return "Output"
	case JmpIfTrue:
		return "Jump if True"
	case JmpIfFalse:
		return "Jump if False"
	case LessThan:
		return "Less than"
	case Equals:
		return "Equals"
	case ModRelativeBase:
		return "Modify Relative Base"
	case Halt:
		return "Halt"
	default:
		return "Unknown opcode"
	}
}

type ParameterMode int

const (
	Position  ParameterMode = 0
	Immediate ParameterMode = 1
	Relative  ParameterMode = 2
)

type computer struct {
	instrPointer int // program counter
	intcode      Intcode
	relativeBase int

	input  chan big.Int
	output chan big.Int
}

func (c *computer) Fetch(ptr int) big.Int {
	c.ensureAddressable(ptr)
	return c.intcode[ptr]
}

func (c *computer) Put(ptr int, newValue big.Int) {
	c.ensureAddressable(ptr)
	c.intcode[ptr] = newValue
}

func (c *computer) Output() chan big.Int {
	return c.output
}

func (c *computer) Input(val big.Int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	c.input <- val
	return nil
}

func (c *computer) Run() (done chan big.Int) {

	done = make(chan big.Int, 1)
	go func() {
		for {
			opcode := c.opcode()
			c.debugPrint("Exec op %s", opcode.String())
			switch opcode {
			case Halt:
				c.halt(done)
				return
			case Add:
				c.alu(func(a, b big.Int) (i big.Int) { return *i.Add(&a, &b) })
			case Mult:
				c.alu(func(a, b big.Int) (i big.Int) { return *i.Mul(&a, &b) })
			case Input:
				c.inputValue()
			case Output:
				c.outputValue()
			case JmpIfTrue:
				c.jumpIf(true)
			case JmpIfFalse:
				c.jumpIf(false)
			case LessThan:
				c.compare(func(a, b big.Int) bool { return a.Cmp(&b) == -1 })
			case Equals:
				c.compare(func(a, b big.Int) bool { return a.Cmp(&b) == 0 })
			case ModRelativeBase:
				c.modRelativeBase()
			default:
				panic(fmt.Sprintf("got unknown op code %d", opcode))
			}
		}
	}()
	return done
}

func (c computer) debugPrint(format string, args ...string) {
	if DEBUG {
		fmt.Printf(format+"\n", args)
	}
}

func (c *computer) instruction() int {
	c.ensureAddressable(c.instrPointer)
	return int(c.intcode[c.instrPointer].Uint64())
}

func (c *computer) opcode() Opcode {
	return Opcode(c.instruction() % 100)
}

func (c *computer) getParameter(pos int) big.Int {
	mode := getMode(c.instruction(), pos)

	switch mode {
	case Position:
		pos := c.Fetch(c.instrPointer + pos)
		return c.Fetch(int(pos.Int64()))
	case Immediate:
		return c.Fetch(c.instrPointer + pos)
	case Relative:
		pos := c.Fetch(c.instrPointer + pos)
		return c.Fetch(c.relativeBase + int(pos.Int64()))
	default:
		panic(fmt.Sprintf("unknown mode %d", mode))
	}
}

func (c *computer) write(pos int, val big.Int) {
	mode := getMode(c.instruction(), pos)

	switch mode {
	case Position:
		pos := c.Fetch(c.instrPointer + pos)
		c.Put(int(pos.Int64()), val)
	case Immediate:
		c.Put(c.instrPointer+pos, val)
	case Relative:
		pos := c.Fetch(c.instrPointer + pos)
		c.Put(c.relativeBase+int(pos.Int64()), val)
	default:
		panic(fmt.Sprintf("unknown mode %d", mode))
	}
}

func getMode(code, pos int) ParameterMode {
	return ParameterMode((code / (100 * int(math.Pow10(pos-1)))) % 10)
}

// Stop execution and return first value on intcode
func (c *computer) halt(done chan big.Int) {
	close(c.output)
	close(c.input)
	done <- c.intcode[0]
	close(done)
}

func (c *computer) alu(f func(big.Int, big.Int) big.Int) {
	res := f(c.getParameter(1), c.getParameter(2))
	c.write(3, res)
	c.instrPointer += 4
}

func (c *computer) inputValue() {
	c.write(1, <-c.input)
	c.instrPointer += 2
}

func (c *computer) outputValue() {
	c.output <- c.getParameter(1)
	c.instrPointer += 2
}

func (c *computer) jumpIf(b bool) {
	param := c.getParameter(1)
	isTrue := param.Cmp(big.NewInt(0)) != 0
	if isTrue == b {
		param := c.getParameter(2)
		c.instrPointer = int(param.Int64())
	} else {
		c.instrPointer += 3
	}
}

func (c *computer) compare(f func(big.Int, big.Int) bool) {
	if f(c.getParameter(1), c.getParameter(2)) {
		c.write(3, *big.NewInt(1))
	} else {
		c.write(3, *big.NewInt(0))
	}
	c.instrPointer += 4
}

func (c *computer) modRelativeBase() {
	param := c.getParameter(1)
	c.relativeBase += int(param.Int64())
	c.instrPointer += 2
}

func (c *computer) ensureAddressable(address int) {
	missingMemSpace := address - (len(c.intcode) - 1)
	if missingMemSpace <= 0 {
		return
	}
	expansion := make(Intcode, missingMemSpace)
	c.intcode = append(c.intcode, expansion...)
}
