package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Program struct {
	Instructions    []int64
	InstructPointer int
	Cpu             map[Operation]int
}

func (p *Program) Exec() {
	for p.InstructPointer = 0; p.InstructPointer < len(p.Instructions); {
		instruct, instructPointerDelta := NewInstruction(p)
		p.Cpu[instruct.Opcode]++
		if instruct.Opcode == OpEnd {
			return
		}
		instruct.Exec()
		p.InstructPointer += instructPointerDelta
	}
}

func (p *Program) StringInstructions() string {
	stringProgram := make([]string, len(p.Instructions))
	for i, v := range p.Instructions {
		stringProgram[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(stringProgram, ",")
}

func (p *Program) CpuInfo() ([]string, int) {
	stringCpu := make([]string, len(p.Cpu))

	i := 0
	sum := 0
	for key, value := range p.Cpu {
		stringCpu[i] = fmt.Sprint(key, ":", value)
		sum += value
		i++
	}
	return stringCpu, sum
}