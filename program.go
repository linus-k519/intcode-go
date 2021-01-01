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
	Ints ints
	// IP is the instruction pointer. It is the index in the Ints array of the
	// instruction that is currently being executed.
	IP int
	// MoveIP indicates whether IP should be moved after executing the instruction.
	MoveIP bool
	// RelBase is the value of the relative base register.
	RelBase int64
	// InputReader is the io.Reader, from which the input for the program is read.
	InputReader io.Reader
	// DebugWriter is the io.Writer where input prompts are written to.
	DebugWriter io.Writer
	// OutputWriter is the io.Writer, in which output of the program is written.
	OutputWriter io.Writer
	// Finish indicates whether the program has finished running.
	Finish bool
	// Stats contains detailed information about the program execution.
	Stats stats
	// Debug indicates whether showDebug outputs should be shown.
	Debug bool
}

// Exec executes a program.
func (p *Program) Exec() {
	p.Stats.start()
	for p.IP = 0; p.IP < len(p.Ints); {
		p.MoveIP = true
		// Parse current instruction
		op := newOpcode(p.Ints[p.IP])
		modes := NewModeList(p.Ints[p.IP], opcodes[op].ArgNum)
		argIndexes := p.newArgIndexList(p.IP+1, modes)
		// Execute instruction
		p.execInstruction(op, argIndexes)
		if p.Finish {
			break
		}
	}
	p.Stats.stop()
}

// execInstruction executes an instruction.
func (p *Program) execInstruction(op opcode, argIndexes []int) {
	// Get function of opcode and execute it
	opInfo := opcodes[op]
	if opInfo.Fn == nil {
		fmt.Fprintln(p.DebugWriter, "Unknown opcode", op.String())
	}
	opInfo.Fn(p, argIndexes)
	if p.Debug {
		p.debugInstruction(op, argIndexes)
	}
	if p.Stats.Activated {
		// Increment operations count
		p.Stats.Operations[op]++
	}
	if p.MoveIP {
		// Move instruction pointer by one for the opcode plus the number of arguments
		p.IP += 1 + len(argIndexes)
	}
}

func (p *Program) debugInstruction(op opcode, argIndexes []int) {
	// Print instruction pointer
	fmt.Fprintf(p.DebugWriter, "IP %3d: ", p.IP)

	// Print opcode and args
	args := make(ints, len(argIndexes))
	for i, index := range argIndexes {
		args[i] = p.Ints[index]
	}
	fmt.Fprintf(p.DebugWriter, "%-60s", op.String()+" args ["+args.String()+"]")

	// Print raw integers
	raw := p.Ints[p.IP : p.IP+opcodes[op].ArgNum+1]
	fmt.Fprintln(p.DebugWriter, " (Raw: "+raw.String()+")")
}

// newArgIndexList returns a list of indexes of the arguments starting by
// startIndex in Ints. They are evaluated according to the specific Mode.
// The number of len(modes) arguments will be returned.
func (p *Program) newArgIndexList(startIndex int, modes []Mode) []int {
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

// New parses a new program of the provided string. For running the program
// additionalMemory many int64s are allocated in addition to the program memory.
// Comment lines starting with an '#' will be ignored, spaces will be converted
// into commas and multiple commas and newlines will be ignored.
func New(intsStr string, additionalMemory uint) *Program {
	intsStr = clean(intsStr)
	intsStrArr := strings.Split(intsStr, ",")

	// Parse each value to a number
	intsArr := make(ints, len(intsStrArr)+int(additionalMemory))
	for i, v := range intsStrArr {
		var err error
		intsArr[i], err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
	}
	return &Program{
		Ints:         intsArr,
		InputReader:  os.Stdin,
		DebugWriter:  os.Stderr,
		OutputWriter: os.Stdout,
	}
}

// clean removes comment lines starting with a '#', converts spaces into commas,
// and removes multiple commas and newlines.
func clean(str string) string {
	// Remove comment lines starting with #
	commentRegex := regexp.MustCompile("(?m)\\s*#.*$")
	str = commentRegex.ReplaceAllString(str, "")

	// Convert spaces and newlines into commas
	str = strings.ReplaceAll(str, "\n", ",")
	str = strings.ReplaceAll(str, " ", ",")

	// Remove possible first empty value (due to a comment at the beginning of the program)
	beginProgramRegex := regexp.MustCompile("^,")
	str = beginProgramRegex.ReplaceAllString(str, "")

	// Remove multiple commas
	multipleCommas := regexp.MustCompile("[,]+")
	str = multipleCommas.ReplaceAllString(str, ",")
	return str
}

func (p *Program) Get(index int) int64 {
	p.increaseMemoryIfNecessary(index)
	if p.Stats.Activated {
		p.Stats.MemoryAccesses["Get"]++
	}
	return p.Ints[index]
}

func (p *Program) Set(index int, value int64) {
	p.increaseMemoryIfNecessary(index)
	if p.Stats.Activated {
		p.Stats.MemoryAccesses["Set"]++
	}
	p.Ints[index] = value
}

func (p *Program) increaseMemoryIfNecessary(index int) {
	if index >= len(p.Ints) {
		// Memory address is out of range -> allocate more memory
		p.increaseMemory(index + 1 - len(p.Ints))
	}
}

// increaseMemory increases the memory of Program.Ints by delta.
func (p *Program) increaseMemory(newSize int) {
	difference := newSize - len(p.Ints)
	if difference <= 0 {
		return
	}

	percentage := (float64(difference) / float64(len(p.Ints))) * 100
	if p.Debug {
		fmt.Fprintf(p.DebugWriter, "Increasing memory by %d ints (%.4f%%)\n", difference, percentage)
	}

	// Make new array of newSize, copy the elements and assign it to program
	intsLarge := make(ints, newSize)
	copy(intsLarge, p.Ints)
	p.Ints = intsLarge
}
