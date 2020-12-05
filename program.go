package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Program struct {
	Instructions []int
	Cpu          map[Operation]int
}

func (p *Program) Exec() {
	for index := 0; index < len(p.Instructions); {
		instruct := Instruction{}
		instruct.Opcode = Operation(p.Instructions[index])
		p.Cpu[instruct.Opcode]++
		index++
		if instruct.Opcode == OpEnd {
			break
		}

		instruct.ScanArgs(&index, p)
		instruct.Exec()
	}
}

func (p *Program) StringInstructions() string {
	stringProgram := make([]string, len(p.Instructions))
	for i, v := range p.Instructions {
		stringProgram[i] = strconv.Itoa(v)
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