package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOpcode(t *testing.T) {
	assert.Equal(t, Opcode(01), NewOpcode(1301))
	assert.Equal(t, Opcode(99), NewOpcode(451599))
}

func TestOpcode_ArgNum(t *testing.T) {
	assert.Equal(t, 3, OpAdd.ArgNum())
	assert.Equal(t, 3, OpMul.ArgNum())
	assert.Equal(t, 3, OpLT.ArgNum())
	assert.Equal(t, 3, OpEq.ArgNum())

	assert.Equal(t, 2, OpJNZ.ArgNum())
	assert.Equal(t, 2, OpJZ.ArgNum())

	assert.Equal(t, 1, OpRelBase.ArgNum())

	assert.Equal(t, 0, Opcode(00).ArgNum())
	assert.Equal(t, 0, OpEnd.ArgNum())
}
