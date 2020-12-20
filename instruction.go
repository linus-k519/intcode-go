package main

import (
	"fmt"
	"io"
)

type Instruction struct {
	Opcode     Opcode
	ArgIndexes []int
	Modes      []Mode
	ArgNum     int
}

// Add performs an 'addition'. It adds the numbers a and b and return the result of the addition.
func Add(a, b int64) int64 {
	return a + b
}

// Mul performs a 'multiplication'. It multiplies the numbers a and b and return the result of the multiplication.
func Mul(a, b int64) int64 {
	return a * b
}

// In performs an 'input'.It prints an input message on w and then reads an input from r and returns it.
func In(w io.Writer, r io.Reader) int64 {
	// If w == nil, there will be an error. But this is ok if you don't want an input query text
	fmt.Fprint(w, "Input: ")

	var val int64
	_, err := fmt.Fscanf(r, "%d", &val)
	if err != nil {
		panic(err)
	}
	return val
}

// Out performs an output. It prints val to w.
func Out(w io.Writer, val int64) {
	fmt.Fprintf(w, "Output: %d", val)
}

// JNZ performs a 'jump non-zero'. It returns jumpPos if val is non-zero and oldPos otherwise.
func JNZ(val int64, jumpPos int, oldPos int) int {
	if val != 0 {
		return jumpPos
	} else {
		return oldPos
	}
}

// JZ performs a 'jump zero'. It returns jumpPos if val is zero and oldPos otherwise.
func JZ(val int64, jumpPos int, oldPos int) int {
	if val == 0 {
		return jumpPos
	} else {
		return oldPos
	}
}

// LT performs a 'less than'. It returns 1 if a < b and 0 if not.
func LT(a, b int64) int64 {
	return boolToInt(a < b)
}

// Eq performs a 'equal'. It returns 1 if a == b and 0 if not.
func Eq(a, b int64) int64 {
	return boolToInt(a == b)
}

// RelBase performs a 'relative base offset'. It adds v to relBase and returns the new relBase.
func RelBase(v, relBase int64) int64 {
	return relBase + v
}

// boolToInt converts a bool to an int. It returns 1 if b is true, and 0 if b is false.
func boolToInt(b bool) int64 {
	if b {
		return 1
	} else {
		return 0
	}
}
