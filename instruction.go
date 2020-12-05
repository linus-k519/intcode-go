package main

import (
	"fmt"
)

type Instruction struct {
	Opcode Operation
	Args   []*int
}

func (instruct *Instruction) Exec() {
	switch instruct.Opcode {
	case OpAdd:
		*instruct.Args[2] = *instruct.Args[0] + *instruct.Args[1]
	case OpMul:
		*instruct.Args[2] = *instruct.Args[0] * *instruct.Args[1]
	case OpOut:
		fmt.Println("Output:", *instruct.Args[0])
	case OpIn:
		fmt.Print("Input: ")
		_, err := fmt.Scanf("%d", instruct.Args[0])
		if err != nil {
			panic(err)
		}
	default:
		panic(fmt.Sprintln("Can not execute opcode", instruct.Opcode))
	}
}

func (instruct *Instruction) ScanArgs(index *int, program Program) {
	numOfArgs := GetNumOfArgs(instruct.Opcode)

	finishIndex := *index + numOfArgs
	if len(program) <= finishIndex {
		panic(fmt.Sprintln("Invalid program format. Missing arguments for opcode", instruct.Opcode))
	}

	for ; *index < finishIndex; *index++ {
		instruct.Args = append(instruct.Args, &program[program[*index]])
	}
}