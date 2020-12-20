package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Program contains the program itself (the Ints) and various registers, which
// control the program flow.
type Program struct {
	// Ints is the actual program, i.e. an array of the instructions with the
	// corresponding arguments.
	Ints []int64
	// IP is the instruction pointer. It is the index in the Ints array of the
	// instruction that is currently being executed.
	IP int
	// RelBase is the value of the relative base register.
	RelBase int64
	// Input is the io.Reader, from which the input for the program is read.
	Input io.Reader
	// Output is the io.Writer, in which output of the program is written.
	Output io.Writer
	// Finish indicates whether the program has finished running.
	Finish bool
	// OperationCount contains the number of executions of each Opcode.
	OperationCount map[Opcode]uint
}

// resetRegisters resets registers of a Program.
func (p *Program) resetRegisters() {
	p.IP = 0
	p.RelBase = 0
	p.Finish = false
	p.OperationCount = map[Opcode]uint{}
}

// Exec executes a program.
func (p *Program) Exec() {
	p.resetRegisters()
	for p.IP = 0; p.IP < len(p.Ints); {
		// Parse current instruct
		opcode := NewOpcode(p.Ints[p.IP])
		modes := NewModeList(p.Ints[p.IP], opcode.ArgNum())
		argIndexes := p.NewArgIndexList(p.IP+1, modes)

		p.ExecInstruction(opcode, argIndexes)
		p.OperationCount[opcode]++
		// TODO: Activate on debug
		fmt.Println(opcode, modes, argIndexes)
		if p.Finish {
			return
		}
	}
}

// ExecInstruction executes an Instruction
func (p *Program) ExecInstruction(opcode Opcode, argIndexes []int) {
	argNum := len(argIndexes)
	moveIP := func() {
		// Move instruction pointer by one (for the opcode) plus number of args.
		p.IP += 1 + argNum
	}

	switch opcode {
	case OpAdd:
		p.Ints[argIndexes[2]] = Add(p.Ints[argIndexes[0]], p.Ints[argIndexes[1]])
		moveIP()
	case OpMul:
		p.Ints[argIndexes[2]] = Mul(p.Ints[argIndexes[0]], p.Ints[argIndexes[1]])
		moveIP()
	case OpIn:
		p.Ints[argIndexes[0]] = In(p.Output, p.Input)
		moveIP()
	case OpOut:
		Out(p.Output, p.Ints[argIndexes[0]])
		moveIP()
	case OpJNZ:
		p.IP = JNZ(p.Ints[argIndexes[0]], int(p.Ints[argIndexes[1]]), p.IP)
	case OpJZ:
		p.IP = JZ(p.Ints[argIndexes[0]], int(p.Ints[argIndexes[1]]), p.IP)
	case OpLT:
		p.Ints[argIndexes[2]] = LT(p.Ints[argIndexes[0]], p.Ints[argIndexes[1]])
		moveIP()
	case OpEq:
		p.Ints[argIndexes[2]] = Eq(p.Ints[argIndexes[0]], p.Ints[argIndexes[1]])
		moveIP()
	case OpRelBase:
		p.RelBase = RelBase(p.Ints[argIndexes[0]], p.RelBase)
		moveIP()
	case OpEnd:
		p.Finish = true
		moveIP()
	default:
		panic(fmt.Sprintf("Could not execute opcode %s", opcode.String()))
	}
}

// NewArgIndexList returns a list of indexes of the arguments starting by
// startIndex in Ints. They are evaluated according to the specific Mode.
// The number of len(modes) arguments will be returned.
func (p *Program) NewArgIndexList(startIndex int, modes []Mode) []int {
	argNum := len(modes)
	endIndex := startIndex + argNum - 1
	if endIndex >= len(p.Ints) {
		panic("Invalid program format. Not enough arguments")
	}

	argIndexes := make([]int, argNum)
	for i := 0; i < argNum; i++ {
		switch modes[i] {
		case ModePos:
			// The value of the argument is the position of the actual value.
			argIndexes[i] = int(p.Ints[startIndex+i])
		case ModeImm:
			// The value of the argument is the argument itself.
			argIndexes[i] = startIndex + i
		case ModeRel:
			// The value of the argument combined with the relative base register results in the position of the actual value.
			argIndexes[i] = int(p.Ints[startIndex+i] + p.RelBase)
		default:
			panic(fmt.Sprintf("Unkown mode %s", modes[i]))
		}
	}
	return argIndexes
}

// NewProgram parses a new program of the provided string. Comment lines starting
// with an '#' will be ignored, spaces will be converted into commas and multiple
// commas and newlines will be ignored.
func NewProgram(str string) *Program {
	str = clean(str)
	intsStr := strings.Split(str, ",")
	ints := make([]int64, len(intsStr))
	for i, v := range intsStr {
		var err error
		ints[i], err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
	}
	return &Program{
		Ints:   ints,
		Input:  os.Stdin,
		Output: os.Stdout,
	}
}

// clean removes comment lines starting with a '#', converts spaces into commas,
// and removes multiple commas and newlines.
func clean(str string) string {
	// Remove comment lines starting with #
	commentRegex := regexp.MustCompile("(?m)\\s*#.*$")
	str = commentRegex.ReplaceAllString(str, "")

	// Convert spaces and newlines into commas
	str = strings.ReplaceAll(str, " ", ",")
	str = strings.ReplaceAll(str, "\n", "")

	// Remove multiple commas
	multipleCommas := regexp.MustCompile("[,]+")
	str = multipleCommas.ReplaceAllString(str, ",")
	return str
}

// StringInts converts Ints to a comma-separated string.
func (p *Program) StringInts() string {
	ints := make([]string, len(p.Ints))
	for i, v := range p.Ints {
		ints[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(ints, ",")
}
