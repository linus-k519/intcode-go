package main

import (
	"fmt"
	log "github.com/linus-k519/llog"
	"math"
	"strings"
)

type Instruction struct {
	Program *Program
	Opcode  Operation
	Modes   []Mode
	Args    []*int64
}

func (instruct *Instruction) Exec() {
	switch instruct.Opcode {
	case OpAdd:
		*instruct.Args[2] = *instruct.Args[0] + *instruct.Args[1]
		log.Debug(fmt.Sprintf("@IP_%d: %d [@%d] %s %d [@%d] = %d [@%d]",
			instruct.Program.InstructPointer,
			*instruct.Args[0], instruct.Program.InstructPointer+1,
			instruct.Opcode,
			*instruct.Args[1], instruct.Program.InstructPointer+2,
			*instruct.Args[2], instruct.Program.InstructPointer+3))
	case OpMul:
		*instruct.Args[2] = *instruct.Args[0] * *instruct.Args[1]
		log.Debug(fmt.Sprintf("@IP_%d: %d [@%d] %s %d [@%d] = %d [@%d]",
			instruct.Program.InstructPointer,
			*instruct.Args[0], instruct.Program.InstructPointer+1,
			instruct.Opcode,
			*instruct.Args[1], instruct.Program.InstructPointer+2,
			*instruct.Args[2], instruct.Program.InstructPointer+3))
	case OpOut:
		fmt.Printf("Output (@IP_%d): %d\n", instruct.Program.InstructPointer, *instruct.Args[0])
	case OpIn:
		fmt.Printf("Input (@IP_%d): ", instruct.Program.InstructPointer)
		_, err := fmt.Scanf("%d", instruct.Args[0])
		if err != nil {
			panic(err)
		}
	case OpRelBaseOff:
		instruct.Program.RelBase += *instruct.Args[0]
	case 0:
		log.Warn("Ignored opcode 0")
	default:
		panic(fmt.Sprintf("@IP_%d: Unknown opcode: %s", instruct.Program.InstructPointer, instruct))
	}
}

func (instruct *Instruction) ScanArgs(program *Program) int {
	numOfArgs := len(instruct.Modes)

	finishInstructPointer := program.InstructPointer + numOfArgs
	if len(program.Instructs) <= finishInstructPointer {
		panic(fmt.Sprintln("Invalid program format. Missing arguments for opcode", instruct.Opcode))
	}

	instruct.Args = make([]*int64, numOfArgs)
	for i := 0; i < numOfArgs; i++ {
		switch instruct.Modes[i] {
		case ModePos:
			instruct.Args[i] = &program.Instructs[program.Instructs[i+1]]
		case ModeImm:
			instruct.Args[i] = &program.Instructs[i+1]
		case ModeRel:
			instruct.Args[i] = &program.Instructs[program.RelBase+program.Instructs[i+1]]
		default:
			panic(fmt.Sprintf("Unkown mode %d", instruct.Modes[i]))
		}
	}
	return numOfArgs
}

func (instruct *Instruction) ScanModeParams(program *Program) {
	mode := program.Instructs[program.InstructPointer] / 1e2
	numOfArgs := instruct.Opcode.numOfArgs()
	instruct.Modes = make([]Mode, numOfArgs)
	for i := 0; i < numOfArgs; i++ {
		// The division results in the position mode if no parameter is specified
		instruct.Modes[i] = Mode((mode / int64(math.Pow10(i))) % 10)
	}
}

func NewInstruction(program *Program) (*Instruction, int) {
	scannedInts := 0
	instruct := new(Instruction)
	instruct.Program = program
	instruct.Opcode = Operation(program.Instructs[program.InstructPointer] % 1e2)
	scannedInts++
	instruct.ScanModeParams(program)
	scannedInts += instruct.ScanArgs(program)
	return instruct, scannedInts
}

func (instruct *Instruction) String() string {
	argsString := make([]string, len(instruct.Args))
	for i, arg := range instruct.Args {
		argsString[i] = fmt.Sprintf("%d (%d)", arg, instruct.Modes[i])
	}
	return fmt.Sprintf("Op=%d Args=%s", instruct.Opcode, strings.Join(argsString, " "))
}
