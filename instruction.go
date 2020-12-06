package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Instruction struct {
	Opcode Operation
	Modes []Mode
	Args   []*int64
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
		panic(fmt.Sprintln("Unknown opcode", instruct.Opcode))
	}
}

func (instruct *Instruction) ScanArgs(program *Program) int {
	numOfArgs := len(instruct.Modes)

	finishInstructPointer := program.InstructPointer + numOfArgs
	if len(program.Instructions) <= finishInstructPointer {
		panic(fmt.Sprintln("Invalid program format. Missing arguments for opcode", instruct.Opcode))
	}

	instruct.Args = make([]*int64, numOfArgs)
	for i := 0; i < numOfArgs; i++ {
		switch instruct.Modes[i] {
		case ModePos:
			instruct.Args[i] = &program.Instructions[program.Instructions[i + 1]]
		case ModeImm:
			instruct.Args[i] = &program.Instructions[i + 1]
		default:
			panic(fmt.Sprintf("Unkown mode %d", instruct.Modes[i]))
		}
	}
	return numOfArgs
}

func (instruct *Instruction) ScanModeParams(program *Program) {
	mode := program.Instructions[program.InstructPointer] / 1e2
	numOfArgs := GetNumOfArgs(instruct.Opcode)
	instruct.Modes = make([]Mode, numOfArgs)
	for i := 0; i < numOfArgs; i++ {
		// The division results in the position mode if no parameter is specified
		instruct.Modes[i] = Mode((mode / int64(math.Pow10(i))) % 10)
	}
}

func NewInstruction(program *Program) (*Instruction, int) {
	scannedInts := 0
	instruct := new(Instruction)
	instruct.Opcode = Operation(program.Instructions[program.InstructPointer] % 1e2)
	scannedInts++
	instruct.ScanModeParams(program)
	scannedInts += instruct.ScanArgs(program)
	return instruct, scannedInts
}

func (instruct *Instruction) String() string {
	argsString := make([]string, len(instruct.Args))
	for i, v := range instruct.Args {
		argsString[i] = strconv.FormatInt(*v, 10)
	}
	return fmt.Sprintf("Op %d Args %s", instruct.Opcode, strings.Join(argsString, " "))
}