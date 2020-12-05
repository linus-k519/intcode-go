package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type Program []int

func (p Program) String() string {
	stringProgram := make([]string, len(p))
	for i, v := range p {
		stringProgram[i] = strconv.Itoa(v)
	}
	return strings.Join(stringProgram, ",")
}

func ExecInstructions(program Program) {
	for i := 0; i < len(program); i++ {
		instruct := Instruction{}
		instruct.Opcode = Operation(program[i])
		if instruct.Opcode == OpEnd {
			break
		}
		if len(program) < i + len(instruct.Args) {
			panic(fmt.Sprintln("Invalid program format. Missing arguments for opcode", instruct.Opcode))
		}
		for j := 0; j < len(instruct.Args); j++ {
			i++
			instruct.Args[j] = &program[program[i]]
		}
		instruct.Exec()
	}
}

func main() {
	showExecTime := flag.Bool("time", false, "Show execution time")
	flag.Parse()

	timeStart := time.Now()

	args := os.Args[1:]
	if len(args) <= 0 || args[0] == "" {
		panic("Please specify a filename")
	}
	filename := args[0]
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	cleanedFile := clean(string(file))
	program := ParseInstructions(cleanedFile)

	ExecInstructions(program)
	fmt.Println(program.String())

	if *showExecTime {
		fmt.Println(time.Since(timeStart))
	}
}
