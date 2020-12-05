package main

const (
	OpAdd Operation = 1
	OpMul Operation = 2
	OpIn  Operation = 3
	OpOut Operation = 4
	OpEnd Operation = 99
)

func GetNumOfArgs(operation Operation) int {
	switch operation {
	case OpIn, OpOut:
		return 1
	case OpAdd, OpMul:
		return 3
	default:
		return 0
	}
}

type Operation int
