package main

import (
	"fmt"
	"math"
	"strings"
)

type Instruction struct {
	Program *Program
	Opcode  Operation
	Modes   []Mode
	Args    []*int64
}

func (instruct *Instruction) Exec() (end bool) {
	switch instruct.Opcode {
	case OpAdd:
		*instruct.Args[2] = *instruct.Args[0] + *instruct.Args[1]
	case OpMul:
		*instruct.Args[2] = *instruct.Args[0] * *instruct.Args[1]
	case OpIn:
		fmt.Printf("Input (@IP_%d): ", instruct.Program.InstructPointer)
		_, err := fmt.Scanf("%d", instruct.Args[0])
		if err != nil {
			panic(err)
		}
	case OpOut:
		fmt.Printf("Output (@IP_%d): %d\n", instruct.Program.InstructPointer, *instruct.Args[0])
	case OpJNZ:
		if *instruct.Args[0] != 0 {
			instruct.Program.InstructPointer = int(*instruct.Args[1])
			return
		}
	case OpJZ:
		if *instruct.Args[0] == 0 {
			instruct.Program.InstructPointer = int(*instruct.Args[1])
			return
		}
	case OpLT:
		if *instruct.Args[0] < *instruct.Args[1] {
			*instruct.Args[2] = 1
		}
	case OpEq:
		if *instruct.Args[0] == *instruct.Args[1] {
			*instruct.Args[2] = 1
		}
	case OpRelBase:
		instruct.Program.RelBase += *instruct.Args[0]
	case OpEnd:
		end = true
	default:
		panic(fmt.Sprintf("@IP_%d: Unknown opcode: %s", instruct.Program.InstructPointer, instruct))
	}
	instruct.Program.InstructPointer += 1 + len(instruct.Args)
	return end
}

func (instruct *Instruction) ScanArgs() int {
	argNum := len(instruct.Modes)

	finishInstructPointer := instruct.Program.InstructPointer + argNum
	if len(instruct.Program.Instructs) <= finishInstructPointer {
		panic(fmt.Sprintln("Invalid program format. Missing arguments for opcode", instruct.Opcode))
	}

	instruct.Args = make([]*int64, argNum)
	for i := 0; i < argNum; i++ {
		switch instruct.Modes[i] {
		case ModePos:
			instruct.Args[i] = &instruct.Program.Instructs[instruct.Program.Instructs[instruct.Program.InstructPointer+i+1]]
		case ModeImm:
			instruct.Args[i] = &instruct.Program.Instructs[instruct.Program.InstructPointer+i+1]
		case ModeRel:
			instruct.Args[i] = &instruct.Program.Instructs[instruct.Program.RelBase+instruct.Program.Instructs[instruct.Program.InstructPointer+i+1]]
		default:
			panic(fmt.Sprintf("Unkown mode %d", instruct.Modes[i]))
		}
	}
	return argNum
}

func (instruct *Instruction) ScanModeParams() {
	mode := instruct.Program.Instructs[instruct.Program.InstructPointer] / 1e2
	numOfArgs := instruct.Opcode.numOfArgs()
	instruct.Modes = make([]Mode, numOfArgs)
	for i := 0; i < numOfArgs; i++ {
		// The division results in the position mode (0), if no parameter is specified
		instruct.Modes[i] = Mode((mode / int64(math.Pow10(i))) % 10)
	}
}

func NewInstruction(program *Program) *Instruction {
	instruct := new(Instruction)
	instruct.Program = program
	instruct.Opcode = Operation(program.Instructs[program.InstructPointer] % 1e2)
	instruct.ScanModeParams()
	instruct.ScanArgs()
	return instruct
}

func (instruct *Instruction) String() string {
	argsString := make([]string, len(instruct.Args))
	for i, arg := range instruct.Args {
		argsString[i] = fmt.Sprintf("%d (%d)", arg, instruct.Modes[i])
	}
	return fmt.Sprintf("Op=%d Args=%s", instruct.Opcode, strings.Join(argsString, " "))
}
