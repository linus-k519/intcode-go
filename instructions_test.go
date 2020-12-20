package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, int64(10), Add(6, 4))
	assert.Equal(t, int64(0), Add(0, 0))
	assert.Equal(t, int64(5), Add(10, -5))
}

func TestMul(t *testing.T) {
	assert.Equal(t, int64(6), Mul(3, 2))
	assert.Equal(t, int64(0), Mul(5, 0))
	assert.Equal(t, int64(-4), Mul(2, -2))
	assert.Equal(t, int64(1), Mul(-1, -1))
}

func TestIn(t *testing.T) {
	in := "5"
	out := new(strings.Builder)
	assert.Equal(t, int64(5), In(out, strings.NewReader(in)))
	assert.NotEmpty(t, out.String())

	in = "abcd"
	out = new(strings.Builder)
	assert.Panics(t, func() { In(out, strings.NewReader(in)) })
}

func TestOut(t *testing.T) {
	out := new(strings.Builder)
	Out(out, 5)
	assert.Contains(t, out.String(), "5")
}

func TestJNZ(t *testing.T) {
	assert.Equal(t, 1337, JNZ(1, 1337, 42))
	assert.Equal(t, 0, JNZ(42, 0, 42))
	assert.Equal(t, 42, JNZ(0, 1337, 42))
}

func TestJZ(t *testing.T) {
	assert.Equal(t, 1337, JZ(0, 1337, 42))
	assert.Equal(t, 0, JZ(42, 1337, 0))
	assert.Equal(t, 42, JZ(1, 1337, 42))
}

func TestLT(t *testing.T) {
	assert.Equal(t, int64(1), LT(42, 1337))
	assert.Equal(t, int64(0), LT(1337, 42))
}

func TestEq(t *testing.T) {
	assert.Equal(t, int64(1), Eq(42, 42))
	assert.Equal(t, int64(0), Eq(1337, 42))
}

func TestRelBase(t *testing.T) {
	assert.Equal(t, int64(15), RelBase(5, 10))
	assert.Equal(t, int64(2), RelBase(-3, 5))
	assert.Equal(t, int64(-2), RelBase(3, -5))
	assert.Equal(t, int64(-8), RelBase(-3, -5))
}

func TestBoolToInt(t *testing.T) {
	assert.Equal(t, int64(1), boolToInt(true))
	assert.Equal(t, int64(0), boolToInt(false))
}

func TestIntToBool(t *testing.T) {
	assert.Equal(t, true, intToBool(5))
	assert.Equal(t, true, intToBool(-3))
	assert.Equal(t, false, intToBool(0))
}
