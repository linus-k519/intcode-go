package main

import (
	"fmt"
	"io"
	"math/rand"
	"syscall"
	"time"
)

// Add performs an addition (a + b).
func Add(a, b int64) int64 {
	return a + b
}

// Mul performs a multiplication (a * b).
func Mul(a, b int64) int64 {
	return a * b
}

// In performs an input. It prints an input message on w and then reads an input from r and returns it.
func In(w io.Writer, r io.Reader) int64 {
	_, err := fmt.Fprint(w, "Input: ")
	if err != nil {
		panic(err)
	}

	var val int64
	_, err = fmt.Fscanf(r, "%d", &val)
	if err != nil {
		panic(err)
	}
	return val
}

// Out performs an output. It prints val to w.
func Out(w io.Writer, val int64) {
	fmt.Fprintf(w, "Output: %d", val)
}

// JNZ performs a jump non-zero. It returns jumpPos if val is non-zero and oldPos otherwise.
func JNZ(val int64, jumpPos int, oldPos int) int {
	if val != 0 {
		return jumpPos
	} else {
		return oldPos
	}
}

// JZ performs a jump zero. It returns jumpPos if val is zero and oldPos otherwise.
func JZ(val int64, jumpPos int, oldPos int) int {
	if val == 0 {
		return jumpPos
	} else {
		return oldPos
	}
}

// LT performs a less than. It returns 1 if a < b and 0 if not.
func LT(a, b int64) int64 {
	return boolToInt(a < b)
}

// Eq performs an equal. It returns 1 if a == b and 0 if not.
func Eq(a, b int64) int64 {
	return boolToInt(a == b)
}

// RelBase performs a relative base change. It adds v to relBase and returns the new relBase.
func RelBase(v, relBase int64) int64 {
	return relBase + v
}

// -- Additional section --

// BitAnd performs a bitwise and (a & b).
func BitAnd(a, b int64) int64 {
	return a & b
}

// BitOr performs a bitwise or (a | b).
func BitOr(a, b int64) int64 {
	return a | b
}

// BitXor performs a bitwise xor (a ^ b).
func BitXor(a, b int64) int64 {
	return a ^ b
}

// Div performs a integer division (a / b).
func Div(a, b int64) int64 {
	return a / b
}

// Mod performs modulo (a % b).
func Mod(a, b int64) int64 {
	return a % b
}

// LShift performs a left shift (a << b).
func LShift(a, b int64) int64 {
	return a << b
}

// RShift performs a right shift (a >> b).
func RShift(a, b int64) int64 {
	return a >> b
}

// Negate negates the value. Returns 1 if v == 0, and returns 0 otherwise.
func Negate(v int64) int64 {
	return boolToInt(!intToBool(v))
}

// Timestamp returns the current unix timestamp.
func Timestamp() int64 {
	return time.Now().Unix()
}

// Random return a random positive number.
func Random() int64 {
	return rand.Int63()
}

// Syscall performs a syscall.
func Syscall(a, b, c int64) {
	syscall.RawSyscall(uintptr(a), uintptr(b), uintptr(c), 0)
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
