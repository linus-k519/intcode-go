package main

import (
	"fmt"
	"math"
)

// Mode defines the mode of an argument, i.e. whether it points to a position or
// is the value itself. See const declaration for concrete values.
type Mode uint8

const (
	// ModePosition position mode.
	ModePosition Mode = 0
	// ModeImmediate immediate mode.
	ModeImmediate Mode = 1
	// ModeRelativeBase relative base mode.
	ModeRelativeBase Mode = 2
)

// modeName contains the names of Mode's.
var modeName = map[Mode]string{
	ModePosition:     "Position",
	ModeImmediate:    "Immediate",
	ModeRelativeBase: "Relative Base",
}

// NewModeList creates a new mode list with num entries from val. Thus, the
// right-most digit is the first and the left-most digit is the last.
func NewModeList(val int64, num int) []Mode {
	val /= 1e2
	modes := make([]Mode, num)
	for i := 0; i < num; i++ {
		// The division results in the position mode (0), if no parameter is specified
		modes[i] = Mode((val / int64(math.Pow10(i))) % 10)
	}
	return modes
}

func (m Mode) String() string {
	text, ok := modeName[m]
	if !ok {
		text = fmt.Sprintf("MODE_%d", m)
	}
	return text
}
