package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func compile() {
	// Read file
	args := os.Args[1:]
	if len(args) <= 0 || args[0] == "" {
		panic("Please specify a filename")
	}
	filename := args[0]
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Clean newlines and comments
	cleanedFile := clean(string(file))
	program := ParseInstructions(cleanedFile)

	// Execute program
	startTime := time.Now()
	program.Exec()
	execTime := time.Since(startTime)

	fmt.Println(program.StringInstructions())

	fmt.Println("FINISHED in", execTime)
	cpuInfo, totalOperations := program.CpuInfo()
	fmt.Println(cpuInfo)
	timePerOperation := fmt.Sprintf("%.3fns", float64(execTime.Nanoseconds()) / float64(totalOperations))
	fmt.Println("Total operations:", totalOperations)
	fmt.Println("Time per operation:", timePerOperation)
}

func main() {
	compile()
}
