package main

import (
	"fmt"
)

type Operation int8

const (
	OpAdd        Operation = 1
	OpMul        Operation = 2
	OpIn         Operation = 3
	OpOut        Operation = 4
	OpRelBaseOff Operation = 9
	OpEnd        Operation = 99
)

var OpName = map[Operation]string{
	OpAdd:        "ADD",
	OpMul:        "MUL",
	OpIn:         "IN",
	OpOut:        "OUT",
	OpEnd:        "END",
	OpRelBaseOff: "REL_BASE_OFF",
}

func (o Operation) numOfArgs() int {
	switch o {
	case OpIn, OpOut, OpRelBaseOff:
		return 1
	case OpAdd, OpMul:
		return 3
	default:
		return 0
	}
}

func (o Operation) String() string {
	text, ok := OpName[o]
	if !ok {
		text = fmt.Sprintf("OP_%02d", o)
	}
	return text
}
