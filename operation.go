package main

const (
	OpAdd = 1 << Operation(iota)
	OpMul
	OpEnd = Operation(99)
)

type Operation int
