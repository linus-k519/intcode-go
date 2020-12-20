package main

import (
	"fmt"
	"io"
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
	// OperationCount contains the number of executions of each Opcode.
	OperationCount map[Opcode]uint
}

// init inits/resets registers of a Program.
func (p *Program) init() {
	p.IP = 0
	p.RelBase = 0
	p.OperationCount = map[Opcode]uint{}
}

// Exec executes a program.
func (p *Program) Exec() {
	p.init()
	for p.IP = 0; p.IP < len(p.Ints); {
		// Parse current instruct
		opcode := NewOpcode(p.Ints[p.IP])
		modes := NewModeList(p.Ints[p.IP], opcode.ArgNum())
		argIndexes := p.NewArgIndexList(p.IP, modes)
		instruct := &Instruction{
			Opcode:     opcode,
			ArgIndexes: argIndexes,
			Modes:      modes,
			ArgNum:     len(argIndexes),
		}

		p.ExecInstruction(instruct)
		p.OperationCount[instruct.Opcode]++
	}
}

// ExecInstruction executes an Instruction
func (p *Program) ExecInstruction(instruct *Instruction) {
	moveIP := func() {
		// Move instruction pointer by one (for the opcode) plus number of args.
		p.IP += 1 + instruct.ArgNum
	}

	switch instruct.Opcode {
	case OpAdd:
		p.Ints[instruct.ArgIndexes[2]] = Add(p.Ints[instruct.ArgIndexes[0]], p.Ints[instruct.ArgIndexes[1]])
		moveIP()
	case OpMul:
		p.Ints[instruct.ArgIndexes[2]] = Mul(p.Ints[instruct.ArgIndexes[0]], p.Ints[instruct.ArgIndexes[1]])
		moveIP()
	case OpIn:
		p.Ints[instruct.ArgIndexes[0]] = In(p.Output, p.Input)
		moveIP()
	case OpOut:
		Out(p.Output, p.Ints[instruct.ArgIndexes[0]])
		moveIP()
	case OpJNZ:
		p.IP = JNZ(p.Ints[instruct.ArgIndexes[0]], int(p.Ints[instruct.ArgIndexes[1]]), p.IP)
	case OpJZ:
		p.IP = JZ(p.Ints[instruct.ArgIndexes[0]], int(p.Ints[instruct.ArgIndexes[1]]), p.IP)
	case OpLT:
		p.Ints[instruct.ArgIndexes[2]] = LT(p.Ints[instruct.ArgIndexes[0]], p.Ints[instruct.ArgIndexes[1]])
		moveIP()
	case OpEq:
		p.Ints[instruct.ArgIndexes[2]] = Eq(p.Ints[instruct.ArgIndexes[0]], p.Ints[instruct.ArgIndexes[1]])
		moveIP()
	case OpRelBase:
		p.RelBase = RelBase(p.Ints[instruct.ArgIndexes[0]], p.RelBase)
		moveIP()
	case OpEnd:
		moveIP()
		return
	}
}

// NewArgIndexList returns a list of indexes of the arguments starting by
// startIndex in Ints. They are evaluated according to the specific Mode.
// The number of len(modes) arguments will be returned.
func (p *Program) NewArgIndexList(startIndex int, modes []Mode) []int {
	argNum := len(modes)
	endIndex := startIndex + argNum
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

// StringInts converts Ints to a comma-separated string.
func (p *Program) StringInts() string {
	intsStr := make([]string, len(p.Ints))
	for i, v := range p.Ints {
		intsStr[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(intsStr, ",")
}

// NewProgram parses a new program of the provided string. Comment lines starting
// with an '#' will be ignored, spaces will be converted into commas, newlines and
// multiple commas will be ignored.
func NewProgram(str string) *Program {
	str = clean(str)
	intsStr := strings.Split(str, ",")
	Ints := make([]int64, len(intsStr))
	for i, v := range intsStr {
		var err error
		Ints[i], err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
	}
	return &Program{Ints: Ints, OperationCount: map[Opcode]uint{}}
}

// clean removes comment lines starting with a '#', converts spaces into commas,
// removes newlines and removes multiple commas.
func clean(str string) string {
	// Remove comment lines starting with #
	commentRegex := regexp.MustCompile("(?m)\\s*#.*$")
	str = commentRegex.ReplaceAllString(str, "")

	// Convert spaces into commas and remove newlines
	str = strings.ReplaceAll(str, " ", ",")
	str = strings.ReplaceAll(str, "\n", "")

	// Remove multiple commas
	multipleCommas := regexp.MustCompile(",+")
	str = multipleCommas.ReplaceAllString(str, ",")
	return str
}
