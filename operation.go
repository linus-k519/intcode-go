package main

import (
	"fmt"
)

// Operation is a two-digit operation code, like 99(END) or 01(ADD). See const declaration for values.
type Operation int8

const (
	// OpAdd adds two numbers.
	OpAdd Operation = 1
	// OpMul multiplies two numbers.
	OpMul Operation = 2
	// OpIn takes a number as input.
	OpIn Operation = 3
	// OpOut outputs a number.
	OpOut Operation = 4
	// OpJNZ jumps to second arg, if first arg is non-zero.
	OpJNZ Operation = 5
	// OpJZ jumps to second arg, if first arg is zero.
	OpJZ Operation = 6
	// OpLT sets third arg to 1, if first arg is less than second arg.
	OpLT Operation = 7
	// OpEq sets third arg to 1, if first arg is equal to second arg.
	OpEq Operation = 8
	// OpRelBase sets the relative base value.
	OpRelBase Operation = 9
	// OpEnd ends the program.
	OpEnd Operation = 99
)

var OpName = map[Operation]string{
	OpAdd:     "ADD",
	OpMul:     "MULTIPLY",
	OpIn:      "INPUT",
	OpOut:     "OUTPUT",
	OpJNZ:     "JUMP NON-ZERO",
	OpJZ:      "JUMP ZERO",
	OpLT:      "LESS THAN",
	OpEq:      "EQUAL",
	OpEnd:     "END",
	OpRelBase: "RELATIVE BASE OFFSET",
}

func (o Operation) numOfArgs() int {
	switch o {
	case OpIn, OpOut, OpRelBase:
		return 1
	case OpJNZ, OpJZ:
		return 2
	case OpAdd, OpMul, OpLT, OpEq:
		return 3
	default:
		return 0
	}
}

func (o *Operation) String() string {
	text, ok := OpName[*o]
	if !ok {
		text = fmt.Sprintf("OP_%02d", o)
	}
	return text
}

func (o *Operation) MarshalJSON() ([]byte, error) {
	return []byte(o.String()), nil
}
