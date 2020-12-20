package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
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
		modes := NewModeList(p.Ints[p.IP], opcodeArgNum[opcode])
		argIndexes := p.NewArgIndexList(p.IP+1, modes)

		p.ExecInstruction(opcode, argIndexes)
		p.OperationCount[opcode]++
		// TODO: Activate on debug
		//fmt.Println(opcode, modes, argIndexes)
		if p.Finish {
			return
		}
	}
}

// ExecInstruction executes an Instruction
func (p *Program) ExecInstruction(opcode Opcode, argIndexes []int) {
	argNum := len(argIndexes)
	oldIP := p.IP

	f := operationFunctions[opcode]
	if f == nil {
		fmt.Println("f is nil")
	}
	f(p, p.GetArgPointers(argIndexes))

	if p.IP == oldIP {
		// Move instruction pointer by one (for the opcode) plus number of args, if it
		// has not been changed by an instruction-pointer-move-function.
		p.IP += 1 + argNum
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
		case ModePosition:
			// The value of the argument is the position of the actual value.
			argIndexes[i] = int(p.Ints[startIndex+i])
		case ModeImmediate:
			// The value of the argument is the argument itself.
			argIndexes[i] = startIndex + i
		case ModeRelativeBase:
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
		Ints:         ints,
		InputReader:  os.Stdin,
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

func (p *Program) GetArgPointers(argIndexes []int) []*int64 {
	argPointers := make([]*int64, len(argIndexes))
	for i := 0; i < len(argIndexes); i++ {
		argPointers[i] = &p.Ints[argIndexes[i]]
	}
	return argPointers
}

var operationFunctions = map[Opcode]func(*Program, []*int64){
	opAdd:                Add,
	opMultiply:           Multiply,
	opInput:              Input,
	opOutput:             Output,
	opJumpNonZero:        JumpNonZero,
	opJumpZero:           JumpZero,
	opLessThan:           LessThan,
	opEqual:              Equal,
	opChangeRelativeBase: ChangeRelativeBase,
	opBitAnd:             BitAnd,
	opBitOr:              BitOr,
	opBitXor:             BitXor,
	opDivision:           Division,
	opModulo:             Modulo,
	opLeftShift:          LeftShift,
	opRightShift:         RightShift,
	opNegate:             Negate,
	opTimestamp:          Timestamp,
	opRandom:             Random,
	opSyscall:            Syscall,
	opEnd:                func(_ *Program, _ []*int64) {},
}

// Add performs an addition (a + b).
func Add(p *Program, args []*int64) {
	*args[2] = *args[0] + *args[1]
}

// Multiply performs a multiplication (a * b).
func Multiply(p *Program, args []*int64) {
	*args[2] = *args[0] * *args[1]
}

// Input performs an input. It prints an input message on w and then reads an input from r and returns it.
func Input(p *Program, args []*int64) {
	_, err := fmt.Fprint(p.OutputWriter, "Input: ")
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fscanf(p.InputReader, "%d", args[0])
	if err != nil {
		panic(err)
	}
}

// Output performs an output. It prints val to w.
func Output(p *Program, args []*int64) {
	fmt.Fprintln(p.OutputWriter, "Output", *args[0])
}

// JumpNonZero performs a jump non-zero. It returns jumpPos if val is non-zero and oldPos otherwise.
func JumpNonZero(p *Program, args []*int64) {
	if *args[0] != 0 {
		p.IP = int(*args[1])
	}
}

// JumpZero performs a jump zero. It returns jumpPos if val is zero and oldPos otherwise.
func JumpZero(p *Program, args []*int64) {
	if *args[0] == 0 {
		p.IP = int(*args[1])
	}
}

// LessThan performs a less than. It returns 1 if a < b and 0 if not.
func LessThan(p *Program, args []*int64) {
	*args[2] = boolToInt(*args[0] < *args[1])
}

// Equal performs an equal. It returns 1 if a == b and 0 if not.
func Equal(p *Program, args []*int64) {
	*args[2] = boolToInt(*args[0] == *args[1])
}

// ChangeRelativeBase performs a relative base change. It adds v to relBase and returns the new relBase.
func ChangeRelativeBase(p *Program, args []*int64) {
	p.RelBase += *args[0]
}

// -- Additional section --

// BitAnd performs a bitwise and (a & b).
func BitAnd(p *Program, args []*int64) {
	*args[2] = *args[0] & *args[1]
}

// BitOr performs a bitwise or (a | b).
func BitOr(p *Program, args []*int64) {
	*args[2] = *args[0] | *args[1]
}

// BitXor performs a bitwise xor (a ^ b).
func BitXor(p *Program, args []*int64) {
	*args[2] = *args[0] ^ *args[1]
}

// Division performs a integer division (a / b).
func Division(p *Program, args []*int64) {
	*args[2] = *args[0] / *args[1]
}

// Modulo performs modulo (a % b).
func Modulo(p *Program, args []*int64) {
	*args[2] = *args[0] % *args[1]
}

// LeftShift performs a left shift (a << b).
func LeftShift(p *Program, args []*int64) {
	*args[2] = *args[0] << *args[1]
}

// RightShift performs a right shift (a >> b).
func RightShift(p *Program, args []*int64) {
	*args[2] = *args[0] >> *args[1]
}

// Negate negates the value. Returns 1 if v == 0, and returns 0 otherwise.
func Negate(p *Program, args []*int64) {
	*args[1] = boolToInt(!intToBool(*args[0]))
}

// Timestamp returns the current unix timestamp.
func Timestamp(p *Program, args []*int64) {
	*args[0] = time.Now().Unix()
}

// Random return a random positive number.
func Random(p *Program, args []*int64) {
	*args[0] = rand.Int63()
}

// Syscall performs a syscall.
func Syscall(p *Program, args []*int64) {
	syscall.RawSyscall(uintptr(*args[0]), uintptr(*args[1]), uintptr(*args[1]), 0)
}

// boolToInt converts a bool to an int. It returns 1 if b is true, and 0 if b is false.
func boolToInt(b bool) int64 {
	if b {
		return 1
	} else {
		return 0
	}
}

// intToBool converts an int to a bool. It return false if v == 0, and true otherwise.
func intToBool(v int64) bool {
	if v == 0 {
		return false
	} else {
		return true
	}
}
