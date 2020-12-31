package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInts_String(t *testing.T) {
	i := ints{}
	assert.Equal(t, "", i.String())
	i = []int64{1, 2, 3, 5}
	assert.Equal(t, "1,2,3,5", i.String())
}
