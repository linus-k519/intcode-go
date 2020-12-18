package main

import (
	"regexp"
	"strconv"
	"strings"
)

func clean(file string) string {
	commentRegex := regexp.MustCompile("(?m)\\s*#.*$")
	file = commentRegex.ReplaceAllString(file, "")
	file = strings.ReplaceAll(file, " ", ",")
	file = strings.ReplaceAll(file, "\n", "")
	return file
}

func ParseProgram(file string) *Program {
	stringInstructions := strings.Split(file, ",")
	intInstructions := make([]int64, len(stringInstructions))
	for i, instruction := range stringInstructions {
		var err error
		intInstructions[i], err = strconv.ParseInt(instruction, 10, 64)
		if err != nil {
			panic(err)
		}
	}
	return &Program{Instructs: intInstructions, OperationCount: map[Operation]uint{}}
}
