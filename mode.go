package main

import (
	"fmt"
	"math"
)

type Mode uint8

const (
	ModePos Mode = 0
	ModeImm Mode = 1
	ModeRel Mode = 2
)

var ModeName = map[Mode]string{
	ModePos: "POS",
	ModeImm: "IMM",
	ModeRel: "REL",
}

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
	text, ok := ModeName[m]
	if !ok {
		text = fmt.Sprintf("MODE_%d", m)
	}
	return text
}
