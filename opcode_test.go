package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOpcode(t *testing.T) {
	assert.Equal(t, opcode(01), newOpcode(1301))
	assert.Equal(t, opcode(99), newOpcode(451599))
}
