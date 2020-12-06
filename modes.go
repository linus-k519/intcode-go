package main

import "strconv"

const (
	ModePos Mode = 0
	ModeImm Mode = 1
)

type Mode int

var ModeName = map[Mode]string{
	ModePos: "POS",
	ModeImm: "IMM",
}

func (m Mode) String() string {
	text, ok := ModeName[m]
	if !ok {
		text = strconv.Itoa(int(m))
	}
	return text
}