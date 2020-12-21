package main

import (
	"fmt"
	"math"
)

// Mode defines the mode of an argument, i.e. whether it points to a Position or
// is the value itself. See const declaration for concrete values.
type Mode uint8

type ModeInfo struct {
	Name string
	Fn   func(*Program, int) int
}

var Modes = [...]ModeInfo{
	0: {
		Name: "Position",
		Fn:   Position,
	},
	1: {
		Name: "Immediate",
		Fn:   Immediate,
	},
	2: {
		Name: "Relative Base",
		Fn:   RelativeBase,
	},
}

func Position(program *Program, index int) int {
	return int(program.Ints[index])
}

func Immediate(_ *Program, index int) int {
	return index
}

func RelativeBase(program *Program, index int) int {
	return int(program.RelBase + program.Ints[index])
}

// NewModeList creates a new mode list with num entries from val. Thus, the
// right-most digit is the first and the left-most digit is the last.
func NewModeList(val int64, num int) []Mode {
	val /= 1e2
	modes := make([]Mode, num)
	for i := 0; i < num; i++ {
		// The division results in the Position mode (0), if no parameter is specified
		modes[i] = Mode((val / int64(math.Pow10(i))) % 10)
	}
	return modes
}

func (m Mode) String() string {
	text := ""
	if int(m) < len(Modes) {
		text = Modes[m].Name
	} else {
		text = fmt.Sprintf("MODE_%d", m)
	}
	return text
}
