package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProgram_clean(t *testing.T) {
	assert.Equal(t, "1,2,3,5", clean("1,2,3,5"))
	assert.Equal(t, "1,2,3,5", clean("\n1,2\n,\n3,5"))
	assert.Equal(t, "1,2,3,5", clean("1 2 3 5"))
	assert.Equal(t, "1,2,3,5", clean("1,2,3,\n# Hallo\n,5"))
	assert.Equal(t, "1,2,3,5", clean("1 \n,\n,\n 2 \n 3,, 5"))
	assert.Equal(t, "1,2,3,5", clean("1 \n,\n,#Hallo\n 2 \n 3,, 5"))
}

func TestNewProgram(t *testing.T) {
	p := New("1 \n,\n,#Hallo\n 2 \n 3,, 5", 0)
	assert.Equal(t, ints{1, 2, 3, 5}, p.Ints)
	assert.Equal(t, 0, p.IP)
	assert.Equal(t, int64(0), p.RelBase)
	assert.False(t, p.Finish)
}

func TestProgram_ExecInstruction(t *testing.T) {
	p := Program{
		Ints:         []int64{1, 2, 3, 0, 99},
		IP:           0,
		RelBase:      0,
		InputReader:  nil,
		OutputWriter: nil,
		Finish:       false,
	}
	p.execInstruction(1, []int{1, 2, 3})
	assert.Equal(t, int64(5), p.Ints[3])
}

func TestProgram_NewArgIndexList(t *testing.T) {
	p := Program{
		Ints:    []int64{1, 0, 3, -38, 99},
		RelBase: 42,
	}
	argIndexes := p.newArgIndexList(1, []Mode{0, 1, 2})
	assert.Equal(t, int64(1), p.Ints[argIndexes[0]])
	assert.Equal(t, int64(3), p.Ints[argIndexes[1]])
	assert.Equal(t, int64(99), p.Ints[argIndexes[2]])
}

func BenchmarkProgram_Copy(b *testing.B) {
	old := []int64{1, 2, 3}
	for i := 0; i < b.N; i++ {
		intsLarge := make([]int64, len(old)+420)
		copy(intsLarge, old)
		old = intsLarge

		if len(old) > 4096*2 {
			old = []int64{1, 2, 3}
		}
	}
}

func BenchmarkProgram_Extend(b *testing.B) {
	old := []int64{1, 2, 3}
	for i := 0; i < b.N; i++ {
		old = append(old, make([]int64, len(old)+420)...)

		if len(old) > 4096*2 {
			old = []int64{1, 2, 3}
		}
	}
}
