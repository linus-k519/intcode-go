package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Program struct {
	Instructs       []int64
	InstructPointer int
	RelBase         int64
	Cpu             map[Operation]uint
}

func (p *Program) Exec() {
	for p.InstructPointer = 0; p.InstructPointer < len(p.Instructs); {
		instruct := NewInstruction(p)
		p.Cpu[instruct.Opcode]++
		if instruct.Opcode == OpEnd {
			return
		}
		instruct.Exec()
	}
}

func (p *Program) StringInstructions() string {
	stringProgram := make([]string, len(p.Instructs))
	for i, v := range p.Instructs {
		stringProgram[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(stringProgram, ",")
}

func (p *Program) CpuInfo() (string, uint) {
	stringCpu := make([]string, len(p.Cpu))
	var sum uint = 0
	i := 0
	for key, value := range p.Cpu {
		stringCpu[i] = fmt.Sprintf("%3dx %s", value, key)
		sum += value
		i++
	}
	return strings.Join(stringCpu, "\n"), sum
}
