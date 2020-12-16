package main

import (
	"flag"
	"fmt"
	log "github.com/linus-k519/llog"
	"io/ioutil"
	"os"
	"time"
)

func run(file string) {
	// Clean newlines and comments
	cleanedFile := clean(file)
	program := ParseProgram(cleanedFile)

	// Execute program
	startTime := time.Now()
	program.Exec()
	execTime := time.Since(startTime)

	if printExecutedProgram {
		fmt.Fprintln(outputFile, program.StringInstructions())
	}

	fmt.Println("-- STATS --")
	fmt.Println("Finished in", execTime)
	cpuInfo, totalOperations := program.CpuInfo()
	fmt.Println(cpuInfo)
	timePerOperation := fmt.Sprintf("%.3fµs", float64(execTime.Nanoseconds())/float64(totalOperations))
	fmt.Println("Total operations:", totalOperations)
	fmt.Println("Time per operation:", timePerOperation)
}

var (
	printExecutedProgram bool
	outputFile           *os.File
)

const version = "v5.1"

func main() {
	log.Config(0)

	flag.BoolVar(&printExecutedProgram, "pep", false, "Print executed program")
	outFile := flag.String("output", "", "File to print output to")
	flag.Parse()
	if *outFile != "" {
		var err error
		outputFile, err = os.OpenFile(*outFile, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			panic(err)
		}
	}

	cmd := flag.Arg(0)
	if cmd == "" {
		fmt.Println("INTCODE Computer", version)
		fmt.Println("Use \"intcode -help\" for help")
		return
	}

	filename := flag.Arg(1)
	if filename == "" {
		panic("Please specify a filename")
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	switch cmd {
	case "run":
		run(string(file))
	case "beauty":
		beautify(string(file))
	default:
		panic("Unknown command")
	}
}
