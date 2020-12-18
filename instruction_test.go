package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstruction_Exec(t *testing.T) {
	instructs := []int64{101, 5, 1, 3, 99}
	p := Program{
		Instructs:       instructs,
		InstructPointer: 0,
		RelBase:         0,
		OperationCount:  map[Operation]uint{},
	}
	instruct := NewInstruction(&p)

	instruct.ScanModeParams()
	assert.Equal(t, []Mode{ModeImm, ModePos, ModePos}, instruct.Modes)

	instruct.ScanArgs()
	assert.Equal(t, []*int64{&instructs[1], &instructs[1], &instructs[3]}, instruct.Args)

	instruct.Exec()
	assert.Equal(t, []int64{101, 5, 1, 10, 99}, p.Instructs)
}
