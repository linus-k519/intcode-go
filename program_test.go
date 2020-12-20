package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProgram_resetRegisters(t *testing.T) {
	p := Program{
		Ints:           nil,
		IP:             42,
		RelBase:        42,
		Input:          nil,
		Output:         nil,
		Finish:         false,
		OperationCount: nil,
	}
	p.resetRegisters()

	assert.NotNil(t, p.OperationCount)
	assert.Equal(t, int64(0), p.RelBase)
	assert.Equal(t, 0, p.IP)
}

func TestProgram_clean(t *testing.T) {
	assert.Equal(t, "1,2,3,5", clean("1,2,3,5"))
	assert.Equal(t, "1,2,3,5", clean("\n1,2\n,\n3,5"))
	assert.Equal(t, "1,2,3,5", clean("1 2 3 5"))
	assert.Equal(t, "1,2,3,5", clean("1,2,3,\n# Hallo\n,5"))
	assert.Equal(t, "1,2,3,5", clean("1 \n,\n,\n 2 \n 3,, 5"))
	assert.Equal(t, "1,2,3,5", clean("1 \n,\n,#Hallo\n 2 \n 3,, 5"))
}

func TestNewProgram(t *testing.T) {
	assert.Equal(t, []int64{1, 2, 3, 5}, NewProgram("1 \n,\n,#Hallo\n 2 \n 3,, 5").Ints)
}

func TestProgram_StringInts(t *testing.T) {
	p := Program{}
	assert.Equal(t, "", p.StringInts())
	p.Ints = []int64{1, 2, 3, 5}
	assert.Equal(t, "1,2,3,5", p.StringInts())
}

func TestProgram_ExecInstruction(t *testing.T) {
	p := Program{
		Ints:           []int64{1, 2, 3, 0, 99},
		IP:             0,
		RelBase:        0,
		Input:          nil,
		Output:         nil,
		Finish:         false,
		OperationCount: nil,
	}
	p.ExecInstruction(OpAdd, []int{1, 2, 3})
	assert.Equal(t, int64(5), p.Ints[3])
}

func TestProgram_NewArgIndexList(t *testing.T) {
	p := Program{
		Ints:    []int64{1, 0, 3, -38, 99},
		RelBase: 42,
	}
	argIndexes := p.NewArgIndexList(1, []Mode{ModePos, ModeImm, ModeRel})
	assert.Equal(t, int64(1), p.Ints[argIndexes[0]])
	assert.Equal(t, int64(3), p.Ints[argIndexes[1]])
	assert.Equal(t, int64(99), p.Ints[argIndexes[2]])
}
