package main

import (
	"strconv"
)

// Opcode is a two-digit operation code, like 99(END) or 01(ADD). See const
// declaration for concrete values.
type Opcode int8

const (
	// OpAdd adds to numbers.
	OpAdd Opcode = 1
	// OpMul multiplies two number.
	OpMul Opcode = 2
	// OpIn inputs a number.
	OpIn Opcode = 3
	// OpOut outputs a number.
	OpOut Opcode = 4
	// OpJNZ jump non-zero.
	OpJNZ Opcode = 5
	// OpJZ jump zero.
	OpJZ Opcode = 6
	// OpLT checks less than.
	OpLT Opcode = 7
	// OpEq checks equality.
	OpEq Opcode = 8
	// OpRelBase sets the relative base register.
	OpRelBase Opcode = 9
	// OpEnd ends the program.
	OpEnd Opcode = 99
)

// NewOpcode extracts an Opcode from an instruction value (Such as 01 from 12201).
func NewOpcode(val int64) Opcode {
	return Opcode(val % 1e2)
}

// ArgNum returns the number of arguments of an Opcode.
func (o Opcode) ArgNum() int {
	switch o {
	case OpIn, OpOut, OpRelBase:
		return 1
	case OpJNZ, OpJZ:
		return 2
	case OpAdd, OpMul, OpLT, OpEq:
		return 3
	default:
		// OpEnd in particular
		return 0
	}
}

// opName contains the names of the Opcode's.
var opName = map[Opcode]string{
	OpAdd:     "Add",
	OpMul:     "Multiply",
	OpIn:      "Input",
	OpOut:     "Output",
	OpJNZ:     "Jump non-zero",
	OpJZ:      "Jump Zero",
	OpLT:      "Less than",
	OpEq:      "Equal",
	OpEnd:     "End",
	OpRelBase: "Relative Base Offset",
}

func (o Opcode) String() string {
	text, ok := opName[o]
	if !ok {
		text = strconv.Itoa(int(o))
	}
	return text
}
