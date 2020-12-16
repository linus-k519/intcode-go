package main

import (
	"fmt"
	"strconv"
	"strings"
)

func beautify(file string) {
	file = clean(file)
	program := ParseProgram(file)

	for i := 0; i < len(program.Instructs); {
		var line []string
		opcode := Operation(program.Instructs[i] % 1e2)
		line = append(line, strconv.Itoa(int(opcode)))
		i++
		argNum := opcode.numOfArgs()
		for j := 0; j < argNum; j++ {
			line = append(line, strconv.FormatInt(program.Instructs[i+j], 10))
		}
		i += argNum
		fmt.Println(strings.Join(line, ","))
		if opcode == OpEnd {
			program.Instructs = program.Instructs[i:]
			fmt.Println(program.StringInstructions())
			return
		}
	}
}
