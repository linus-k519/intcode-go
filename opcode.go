package main

import (
	"strconv"
)

// Opcode is a two-digit operation code, like 99(END) or 01(ADD). See const
// declaration for concrete values.
type Opcode int8

// NewOpcode extracts an Opcode from an instruction value (Such as 01 from 12201).
func NewOpcode(val int64) Opcode {
	return Opcode(val % 1e2)
}

type OpcodeInfo struct {
	Name   string
	ArgNum int
	Fn     func(p *Program, args []*int64)
}

var Opcodes = map[Opcode]OpcodeInfo{
	1: {
		Name:   "Add",
		ArgNum: 3,
		Fn:     Add,
	},
	2: {
		Name:   "Multiply",
		ArgNum: 3,
		Fn:     Multiply,
	},
	3: {
		Name:   "Input",
		ArgNum: 1,
		Fn:     Input,
	},
	4: {
		Name:   "Output",
		ArgNum: 1,
		Fn:     Output,
	},
	5: {
		Name:   "Jump non-zero",
		ArgNum: 2,
		Fn:     JumpNonZero,
	},
	6: {
		Name:   "Jump zero",
		ArgNum: 2,
		Fn:     JumpZero,
	},
	7: {
		Name:   "Less than",
		ArgNum: 3,
		Fn:     LessThan,
	},
	8: {
		Name:   "Equal",
		ArgNum: 3,
		Fn:     Equal,
	},
	9: {
		Name:   "Change relative base",
		ArgNum: 1,
		Fn:     ChangeRelativeBase,
	},
	10: {
		Name:   "Bitwise And",
		ArgNum: 3,
		Fn:     BitAnd,
	},
	11: {
		Name:   "Bitwise Or",
		ArgNum: 3,
		Fn:     BitOr,
	},
	12: {
		Name:   "Bitwise Xor",
		ArgNum: 3,
		Fn:     BitXor,
	},
	13: {
		Name:   "Division",
		ArgNum: 3,
		Fn:     Division,
	},
	14: {
		Name:   "Modulo",
		ArgNum: 3,
		Fn:     Modulo,
	},
	15: {
		Name:   "Left Shift",
		ArgNum: 3,
		Fn:     LeftShift,
	},
	16: {
		Name:   "Right shift",
		ArgNum: 3,
		Fn:     RightShift,
	},
	17: {
		Name:   "Negate",
		ArgNum: 2,
		Fn:     Negate,
	},
	18: {
		Name:   "Timestamp",
		ArgNum: 1,
		Fn:     Timestamp,
	},
	19: {
		Name:   "Random",
		ArgNum: 1,
		Fn:     Random,
	},
	20: {
		Name:   "Absolute",
		ArgNum: 2,
		Fn:     Absolute,
	},
	80: {
		Name:   "Syscall",
		ArgNum: 3,
		Fn:     Syscall,
	},
	99: {
		Name:   "End",
		ArgNum: 0,
		Fn:     End,
	},
}

func (o Opcode) String() string {
	opInfo, ok := Opcodes[o]
	var text string
	if ok {
		text = opInfo.Name
	} else {
		text = strconv.Itoa(int(o))
	}
	return text
}
