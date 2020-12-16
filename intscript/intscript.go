package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func isResult(arr []string) bool {
	return len(arr) > 0
}

var (
	newVarExp    = regexp.MustCompile("var (\\D+) = (\\d+)")
	updateVarExp = regexp.MustCompile("(\\D+) = (\\d+)")
)

func main() {
	file, err := ioutil.ReadFile("intscript/test.is")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		commentExp := regexp.MustCompile("#.*")
		if line == "" || commentExp.MatchString(line) {
			continue
		}

		newVarResult := newVarExp.FindStringSubmatch(line)
		updateVarResult := updateVarExp.FindStringSubmatch(line)

		if isResult(newVarResult) {
			varName := newVarResult[1]
			varValue := newVarResult[2]
			fmt.Println("Set", varName, "=", varValue)
		} else if isResult(updateVarResult) {
			varName := newVarResult[1]
			varValue := newVarResult[2]
			fmt.Println("Update", varName, "=", varValue)
		} else {
			fmt.Println("Unknown line", line)
		}
	}
}
