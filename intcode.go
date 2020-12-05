package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	showExecTime := flag.Bool("time", false, "Show execution time")
	flag.Parse()

	timeStart := time.Now()

	args := os.Args[1:]
	if len(args) <= 0 || args[0] == "" {
		panic("Please specify a filename")
	}
	filename := args[0]
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	cleanedFile := clean(string(file))
	program := ParseInstructions(cleanedFile)

	program.Exec()
	fmt.Println(program)

	if *showExecTime {
		fmt.Println(time.Since(timeStart))
	}
}
