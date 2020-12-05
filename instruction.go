package main

import (
	"fmt"
)

type Instruction struct {
	Opcode Operation
	Args   [3]*int
}

const (
	IN1 = iota
	IN2
	OUT
)

func (i *Instruction) Exec() {
	switch i.Opcode {
	case OpAdd:
		*i.Args[OUT] = *i.Args[IN1] + *i.Args[IN2]
	case OpMul:
		*i.Args[OUT] = *i.Args[IN1] * *i.Args[IN2]
	default:
		panic(fmt.Sprintln("Can not execute opcode", i.Opcode))
	}
}
