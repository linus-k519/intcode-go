package main

type Operation int

const (
	OpAdd Operation = 1
	OpMul Operation = 2
	OpIn  Operation = 3
	OpOut Operation = 4
	OpEnd Operation = 99
)

var OpName = map[Operation]string{
	OpAdd: "ADD",
	OpMul: "MUL",
	OpIn: "IN",
	OpOut: "OUT",
	OpEnd: "END",
}

func GetNumOfArgs(o Operation) int {
	switch o {
	case OpIn, OpOut:
		return 1
	case OpAdd, OpMul:
		return 3
	default:
		return 0
	}
}

func (o Operation) String() string {
	text, ok := OpName[o]
	if !ok {
		text = "UNKNOWN"
	}
	return text
}
