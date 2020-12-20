package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOpcode(t *testing.T) {
	assert.Equal(t, Opcode(01), NewOpcode(1301))
	assert.Equal(t, Opcode(99), NewOpcode(451599))
}
