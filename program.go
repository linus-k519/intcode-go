package main

import (
	"fmt"
	log "github.com/linus-k519/llog"
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
	// InputReader is the io.Reader, from which the input for the program is read.
	InputReader io.Reader
	// OutputWriter is the io.Writer, in which output of the program is written.
	OutputWriter io.Writer
	// Finish indicates whether the program has finished running.
	Finish bool
	// OperationCount contains the number of executions of each Opcode.
	OperationCount map[Opcode]uint
}

// Exec executes a program.
func (p *Program) Exec() {
	for p.IP = 0; p.IP < len(p.Ints); {
		// Parse current instruct
		opcode := NewOpcode(p.Ints[p.IP])
		modes := NewModeList(p.Ints[p.IP], Opcodes[opcode].ArgNum)
		argIndexes := p.NewArgIndexList(p.IP+1, modes)

		p.ExecInstruction(opcode, argIndexes)
		p.OperationCount[opcode]++
		if p.Finish {
			return
		}
	}
}

// ExecInstruction executes an Instruction
func (p *Program) ExecInstruction(opcode Opcode, argIndexes []int) {
	argNum := len(argIndexes)
	oldIP := p.IP

	if int(opcode) >= len(Opcodes) {
		panic(fmt.Sprint("Unknown opcode ", opcode))
	}
	opInfo := Opcodes[opcode]
	args := p.GetArgPointers(argIndexes)

	opInfo.Fn(p, args)
	//printArgs(opcode, args)

	if p.IP == oldIP {
		// If the instruction pointer has not been changed by an
		// instruction-pointer-move-function, move it by one (for the opcode) plus number
		// of args
		p.IP += 1 + argNum
	}
}

func printArgs(opcode Opcode, args []*int64) {
	argsStr := make([]string, len(args))
	for i, v := range args {
		argsStr[i] = strconv.FormatInt(*v, 10)
	}
	log.Debug("Instruction:", opcode, "["+strings.Join(argsStr, " ")+"]")
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
		if int(modes[i]) > len(Modes) {
			panic(fmt.Sprintf("Unkown mode %s", modes[i]))
		}
		info := Modes[modes[i]]
		argIndexes[i] = info.Fn(p, startIndex+i)
	}
	return argIndexes
}

// NewProgram parses a new program of the provided string. Comment lines starting
// with an '#' will be ignored, spaces will be converted into commas and multiple
// commas and newlines will be ignored.
func NewProgram(str string, additionalMemory uint) *Program {
	str = clean(str)
	intsStr := strings.Split(str, ",")
	ints := make([]int64, len(intsStr)+int(additionalMemory))
	for i, v := range intsStr {
		var err error
		ints[i], err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
	}
	return &Program{
		Ints:           ints,
		InputReader:    os.Stdin,
		OutputWriter:   os.Stdout,
		OperationCount: map[Opcode]uint{},
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

// increaseMemory increases the memory of Program.Ints by delta.
func (p *Program) increaseMemory(delta uint) {
	// Reserve a little bit more memory right away, because the reservation is very
	// slow and future requests could access later addresses
	delta += uint(float64(len(p.Ints))*0.05) + 1

	percentage := (float64(delta) / float64(len(p.Ints))) * 100
	log.Info(fmt.Sprintf("Increasing memory by %d ints (%.4f%%)", delta, percentage))

	// Request new array of delta, copy the elements and assign it to program
	intsMem := make([]int64, delta+uint(len(p.Ints)))
	copy(intsMem, p.Ints)
	p.Ints = intsMem
}

func (p *Program) GetArgPointers(argIndexes []int) []*int64 {
	argPointers := make([]*int64, len(argIndexes))
	for i := 0; i < len(argIndexes); i++ {
		if argIndexes[i] >= len(p.Ints) {
			// Memory address is out of range -> allocate more memory
			p.increaseMemory(uint(argIndexes[i] + 1 - len(p.Ints)))
		}
		argPointers[i] = &p.Ints[argIndexes[i]]
	}
	return argPointers
}
