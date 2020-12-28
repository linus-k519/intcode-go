package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOpcode(t *testing.T) {
	assert.Equal(t, opcode(01), NewOpcode(1301))
	assert.Equal(t, opcode(99), NewOpcode(451599))
}
