package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Memory []int

/// Creates new memory from the given slice.
func NewMemory(mem []int) Memory {
	return Memory(mem)
}

func (mem *Memory) slice() []int {
	return []int(*mem)
}

/// Loads a value from the given address.
func (mem *Memory) Load(addr int) int {
	return mem.slice()[addr]
}

/// Stores the given value at the given address.
func (mem *Memory) Store(addr, val int) {
	mem.slice()[addr] = val
}

/// Returns a cloned memory instance which can be mutated independently.
func (mem *Memory) Clone() Memory {
	clone := make([]int, len(mem.slice()))
	copy(clone, mem.slice())
	return Memory(clone)
}

func splitComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i, c := range data {
		switch c {
		case ',', ' ', '\r', '\n':
			return i + 1, data[:i], nil
		}
	}

	if !atEOF {
		return 0, nil, nil
	}

	return 0, data, bufio.ErrFinalToken
}

/// Reads comma-separated program opcodes from the given reader.
func ReadMemory(reader io.Reader) (mem Memory, err error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(splitComma)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		val, err := strconv.Atoi(text)
		if err != nil {
			return nil, err
		}

		mem = append(mem, val)
	}

	return Memory(mem), nil
}

const (
	// Add value at [1] and value at [2], store sum at [3].
	// (3 arguments)
	OpcodeAdd = 1
	// Multiply value at [2] and value at [2], store product at [3].
	// (3 arguments)
	OpcodeMul = 2
	// Immediate program termination.
	// (0 arguments)
	OpcodeHalt = 99
)

type Program struct {
	init Memory
	mem  Memory
	pc   int
}

/// Runs the program in the given memory and returns the modified memory.
func NewProgram(mem Memory) Program {
	return Program{mem, mem.Clone(), 0}
}

/// Resets the program to its initial state.
func (p *Program) Reset() {
	p.mem = p.init.Clone()
	p.pc = 0
}

// Runs the program by interpreting all opcodes until `99` (Halt), returning the
// output of the program (value at address 0 after termination).
func (p *Program) Run() int {
	for {
		val := p.nextValue()
		switch val {
		case OpcodeAdd:
			v1 := p.mem.Load(p.nextValue())
			v2 := p.mem.Load(p.nextValue())
			p.mem.Store(p.nextValue(), v1+v2)
		case OpcodeMul:
			v1 := p.mem.Load(p.nextValue())
			v2 := p.mem.Load(p.nextValue())
			p.mem.Store(p.nextValue(), v1*v2)
		case OpcodeHalt:
			// Reset the program counter to the halt instruction. This will cause the
			// next invocation of Run() to halt immediately, unless Reset() is called.
			// Also, this has the nice side effect of leaving the PC on the halt
			// instruction instead of advancing it.
			p.pc--
			return p.mem.Load(0)
		default:
			panic(fmt.Sprintf("invalid opcode: %v", val))
		}
	}
}

// Advances the program counter and returns the next value.
//
// This value can either be an operator (opcode) or a parameter. The number of
// parameters depends on the instruction.
func (p *Program) nextValue() int {
	val := p.mem.Load(p.pc)
	p.pc++
	return val
}

func RunWith(p *Program, noun, verb int) int {
	p.Reset()
	p.mem.Store(1, noun)
	p.mem.Store(2, verb)
	return p.Run()
}

func main() {
	mem, err := ReadMemory(os.Stdin)
	if err != nil {
		panic(err)
	}

	program := NewProgram(mem)
	fmt.Println("Part 1:", RunWith(&program, 12, 2))

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			result := RunWith(&program, noun, verb)

			if result == 19690720 {
				fmt.Println("Part 2:", 100*noun+verb)
				return
			}
		}
	}
}
