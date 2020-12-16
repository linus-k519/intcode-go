package main

import "fmt"

const (
	ModePos Mode = 0
	ModeImm Mode = 1
	ModeRel Mode = 2
)

type Mode uint8

var ModeName = map[Mode]string{
	ModePos: "POS",
	ModeImm: "IMM",
	ModeRel: "REL",
}

func (m Mode) String() string {
	text, ok := ModeName[m]
	if !ok {
		text = fmt.Sprintf("MODE_%d", m)
	}
	return text
}
