package main

import (
	"fmt"
	"math/rand"
	"syscall"
	"time"
)

// Add performs an addition (a + b).
func Add(_ *Program, args []*int64) {
	*args[2] = *args[0] + *args[1]
}

// Multiply performs a multiplication (a * b).
func Multiply(_ *Program, args []*int64) {
	*args[2] = *args[0] * *args[1]
}

// Input performs an input. It prints an input message on Program.InputWriter and
// then reads an input from Program.InputReader.
func Input(p *Program, args []*int64) {
	if p.InputWriter != nil {
		fmt.Fprint(p.InputWriter, "Input: ")
	}

	_, err := fmt.Fscanf(p.InputReader, "%d", args[0])
	if err != nil {
		panic(err)
	}
}

// Output performs an output. It prints val to w.
func Output(p *Program, args []*int64) {
	fmt.Fprintln(p.OutputWriter, "Output:", *args[0])
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
func LessThan(_ *Program, args []*int64) {
	*args[2] = boolToInt(*args[0] < *args[1])
}

// Equal performs an equal. It returns 1 if a == b and 0 if not.
func Equal(_ *Program, args []*int64) {
	*args[2] = boolToInt(*args[0] == *args[1])
}

// ChangeRelativeBase performs a relative base change. It adds v to relBase and returns the new relBase.
func ChangeRelativeBase(p *Program, args []*int64) {
	p.RelBase += *args[0]
}

// -- Additional section --

// BitAnd performs a bitwise and (a & b).
func BitAnd(_ *Program, args []*int64) {
	*args[2] = *args[0] & *args[1]
}

// BitOr performs a bitwise or (a | b).
func BitOr(_ *Program, args []*int64) {
	*args[2] = *args[0] | *args[1]
}

// BitXor performs a bitwise xor (a ^ b).
func BitXor(_ *Program, args []*int64) {
	*args[2] = *args[0] ^ *args[1]
}

// Division performs a integer division (a / b).
func Division(_ *Program, args []*int64) {
	*args[2] = *args[0] / *args[1]
}

// Modulo performs modulo (a % b).
func Modulo(_ *Program, args []*int64) {
	*args[2] = *args[0] % *args[1]
}

// LeftShift performs a left shift (a << b).
func LeftShift(_ *Program, args []*int64) {
	*args[2] = *args[0] << *args[1]
}

// RightShift performs a right shift (a >> b).
func RightShift(_ *Program, args []*int64) {
	*args[2] = *args[0] >> *args[1]
}

// Negate negates the value. Returns 1 if v == 0, and returns 0 otherwise.
func Negate(_ *Program, args []*int64) {
	*args[1] = boolToInt(!intToBool(*args[0]))
}

// Timestamp returns the current unix timestamp.
func Timestamp(_ *Program, args []*int64) {
	*args[0] = time.Now().Unix()
}

// Random return a random positive number.
func Random(_ *Program, args []*int64) {
	*args[0] = rand.Int63()
}

// Absolute calculates the positive value of args[0] and saves it into args[1]
func Absolute(_ *Program, args []*int64) {
	if *args[0] < 0 {
		*args[1] = -*args[0]
	} else {
		*args[1] = *args[0]
	}
}

// Syscall performs a syscall.
func Syscall(_ *Program, args []*int64) {
	syscall.RawSyscall(uintptr(*args[0]), uintptr(*args[1]), uintptr(*args[1]), 0)
}

// End sets Program.Finish to true to end the program.
func End(p *Program, _ []*int64) {
	p.Finish = true
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
