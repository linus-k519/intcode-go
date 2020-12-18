package main

import (
	"strconv"
	"strings"
)

type Program struct {
	Instructs       []int64
	InstructPointer int
	RelBase         int64
	OperationCount  map[Operation]uint
}

func (p *Program) Exec() {
	for p.InstructPointer = 0; p.InstructPointer < len(p.Instructs); {
		instruct := NewInstruction(p)
		p.OperationCount[instruct.Opcode]++
		end := instruct.Exec()
		if end {
			return
		}
	}
}

func (p *Program) StringInstructions() string {
	stringProgram := make([]string, len(p.Instructs))
	for i, v := range p.Instructs {
		stringProgram[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(stringProgram, ",")
}
