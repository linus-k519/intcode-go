package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewModeList(t *testing.T) {
	assert.Equal(t, []Mode{3, 2, 1}, NewModeList(12301, 3))
	assert.Equal(t, []Mode{}, NewModeList(12301, 0))
	assert.Equal(t, []Mode{3, 2, 1, 0}, NewModeList(12301, 4))
}

func TestPosition(t *testing.T) {
	p := Program{Ints: []int64{4}}
	assert.Equal(t, 4, Position(&p, 0))
}

func TestImmediate(t *testing.T) {
	p := Program{Ints: []int64{4}}
	assert.Equal(t, 0, Immediate(&p, 0))
}

func TestRelativeBase(t *testing.T) {
	p := Program{Ints: []int64{4}, RelBase: 42}
	assert.Equal(t, 46, RelativeBase(&p, 0))
}
