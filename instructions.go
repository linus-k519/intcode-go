package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Add sets arg[2] to arg[0] + arg[1]
func Add(p *Program, argIndexes []int) {
	p.Set(argIndexes[2], p.Get(argIndexes[0])+p.Get(argIndexes[1]))
}

// Multiply sets arg[2] to arg[0] * arg[1]
func Multiply(p *Program, argIndexes []int) {
	p.Set(argIndexes[2], p.Get(argIndexes[0])*p.Get(argIndexes[1]))
}

// Input prints an input message on Program.DebugWriter and then reads an input
// from Program.InputReader to arg[0].
func Input(p *Program, argIndexes []int) {
	// Show input prompt
	if p.InputReader == os.Stdin {
		fmt.Fprint(p.DebugWriter, "Input: ")
	}

	// Read input value
	var value int64
	_, err := fmt.Fscanf(p.InputReader, "%d", &value)
	if err != nil {
		panic(err)
	}
	p.Set(argIndexes[0], value)
}

// Output prints arg[0] to Program.OutputWriter.
func Output(p *Program, argIndexes []int) {
	fmt.Fprintln(p.OutputWriter, p.Get(argIndexes[0]))
}

// JumpNonZero sets Program.IP to arg[1], if arg[0] is non-zero.
func JumpNonZero(p *Program, argIndexes []int) {
	if p.Get(argIndexes[0]) != 0 {
		p.IP = int(p.Get(argIndexes[1]))
		p.MoveIP = false
	}
}

// JumpZero sets Program.IP to arg[1], if arg[0] is zero.
func JumpZero(p *Program, argIndexes []int) {
	if p.Get(argIndexes[0]) == 0 {
		p.IP = int(p.Get(argIndexes[1]))
		p.MoveIP = false
	}
}

// LessThan sets arg[2] to 1, if arg[0] < arg[1]. Otherwise sets arg[2] to 0.
func LessThan(p *Program, argIndexes []int) {
	val := boolToInt(p.Get(argIndexes[0]) < p.Get(argIndexes[1]))
	p.Set(argIndexes[2], val)
}

// Equal sets arg[2] to 1, if arg[0] == arg[1]. Otherwise sets arg[2] to 0.
func Equal(p *Program, argIndexes []int) {
	val := boolToInt(p.Get(argIndexes[0]) == p.Get(argIndexes[1]))
	p.Set(argIndexes[2], val)
}

// AddRelativeBase adds arg[0] to Program.RelBase.
func AddRelativeBase(p *Program, argIndexes []int) {
	p.RelBase += p.Get(argIndexes[0])
}

// -- Additional section --

// BitAnd performs a bitwise and (a & b).
func BitAnd(p *Program, argIndexes []int) {
	p.Ints[argIndexes[2]] = p.Ints[argIndexes[0]] & p.Ints[argIndexes[1]]
}

// BitOr performs a bitwise or (a | b).
func BitOr(p *Program, argIndexes []int) {
	p.Ints[argIndexes[2]] = p.Ints[argIndexes[0]] | p.Ints[argIndexes[1]]
}

// BitXor performs a bitwise xor (a ^ b).
func BitXor(p *Program, argIndexes []int) {
	p.Ints[argIndexes[2]] = p.Ints[argIndexes[0]] ^ p.Ints[argIndexes[1]]
}

// Division performs a integer division (a / b).
func Division(p *Program, argIndexes []int) {
	p.Ints[argIndexes[2]] = p.Ints[argIndexes[0]] / p.Ints[argIndexes[1]]
}

// Modulo performs modulo (a % b).
func Modulo(p *Program, argIndexes []int) {
	p.Ints[argIndexes[2]] = p.Ints[argIndexes[0]] % p.Ints[argIndexes[1]]
}

// LeftShift performs a left shift (a << b).
func LeftShift(p *Program, argIndexes []int) {
	p.Ints[argIndexes[2]] = p.Ints[argIndexes[0]] << p.Ints[argIndexes[1]]
}

// RightShift performs a right shift (a >> b).
func RightShift(p *Program, argIndexes []int) {
	p.Ints[argIndexes[2]] = p.Ints[argIndexes[0]] >> p.Ints[argIndexes[1]]
}

// Negate negates the value. Returns 1 if v == 0, and returns 0 otherwise.
func Negate(p *Program, argIndexes []int) {
	p.Ints[argIndexes[1]] = boolToInt(!intToBool(p.Ints[argIndexes[0]]))
}

// Timestamp returns the current unix timestamp.
func Timestamp(p *Program, argIndexes []int) {
	p.Ints[argIndexes[0]] = time.Now().Unix()
}

// Random return a random positive number.
func Random(p *Program, argIndexes []int) {
	p.Ints[argIndexes[0]] = rand.Int63()
}

// Absolute calculates the positive value of argIndexes[0] and saves it into argIndexes[1]
func Absolute(p *Program, argIndexes []int) {
	if p.Ints[argIndexes[0]] < 0 {
		p.Ints[argIndexes[1]] = -p.Ints[argIndexes[0]]
	} else {
		p.Ints[argIndexes[1]] = p.Ints[argIndexes[0]]
	}
}

// Syscall performs a syscall.
func Syscall(_ *Program, _ []int) {
	//syscall.RawSyscall(uintptr(*argIndexes[0]), uintptr(*argIndexes[1]), uintptr(*argIndexes[1]), 0)
}

// End sets Program.Finish to true to end the program.
func End(p *Program, _ []int) {
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
