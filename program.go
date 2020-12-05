package main

import (
	"strconv"
	"strings"
)

type Program []int

func (p Program) String() string {
	stringProgram := make([]string, len(p))
	for i, v := range p {
		stringProgram[i] = strconv.Itoa(v)
	}
	return strings.Join(stringProgram, ",")
}

func (p Program) Exec() {
	for index := 0; index < len(p); {
		instruct := Instruction{}
		instruct.Opcode = Operation(p[index])
		index++
		if instruct.Opcode == OpEnd {
			break
		}

		instruct.ScanArgs(&index, p)
		instruct.Exec()
	}
}
