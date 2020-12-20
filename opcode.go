package main

import (
	"strconv"
)

// Opcode is a two-digit operation code, like 99(END) or 01(ADD). See const
// declaration for concrete values.
type Opcode int8

const (
	opAdd Opcode = iota + 1
	opMultiply
	opInput
	opOutput
	opJumpNonZero
	opJumpZero
	opLessThan
	opEqual
	opChangeRelativeBase
	opBitAnd
	opBitOr
	opBitXor
	opDivision
	opModulo
	opLeftShift
	opRightShift
	opNegate
	opTimestamp
	opRandom
	opSyscall Opcode = 80
	opEnd     Opcode = 99
)

// NewOpcode extracts an Opcode from an instruction value (Such as 01 from 12201).
func NewOpcode(val int64) Opcode {
	return Opcode(val % 1e2)
}

var opcodeArgNum = map[Opcode]int{
	opAdd:                3,
	opMultiply:           3,
	opInput:              1,
	opOutput:             1,
	opJumpNonZero:        2,
	opJumpZero:           2,
	opLessThan:           3,
	opEqual:              3,
	opEnd:                0,
	opChangeRelativeBase: 1,
	opBitAnd:             3,
	opBitOr:              3,
	opBitXor:             3,
	opDivision:           3,
	opModulo:             3,
	opLeftShift:          3,
	opRightShift:         3,
	opNegate:             2,
	opTimestamp:          1,
	opRandom:             1,
	opSyscall:            3,
}

// opName contains the names of the Opcode's.
var opName = map[Opcode]string{
	opAdd:                "Add",
	opMultiply:           "Multiply",
	opInput:              "InputReader",
	opOutput:             "OutputWriter",
	opJumpNonZero:        "Jump non-zero",
	opJumpZero:           "Jump Zero",
	opLessThan:           "Less than",
	opEqual:              "Equal",
	opEnd:                "End",
	opChangeRelativeBase: "Relative Base Offset",
	opBitAnd:             "Bitwise And",
	opBitOr:              "Bitwise Or",
	opBitXor:             "Bitwise Xor",
	opDivision:           "Division",
	opModulo:             "Modulo",
	opLeftShift:          "Left shift",
	opRightShift:         "Right shift",
	opNegate:             "Negate",
	opTimestamp:          "Timestamp",
	opRandom:             "Random",
	opSyscall:            "Syscall",
}

func (o Opcode) String() string {
	text, ok := opName[o]
	if !ok {
		text = strconv.Itoa(int(o))
	}
	return text
}
