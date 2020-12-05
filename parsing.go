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

func ParseInstructions(file string) Program {
	stringInstructions := strings.Split(file, ",")
	intInstructions := make([]int, len(stringInstructions))
	for i, instruction := range stringInstructions {
		var err error
		intInstructions[i], err = strconv.Atoi(instruction)
		if err != nil {
			panic(err)
		}
	}
	return intInstructions
}
