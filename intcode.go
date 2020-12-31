package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

const version = "v9.3"

var (
	outputFilename   string
	outputFile       *os.File
	inputFilename    string
	inputFile        *os.File
	debug            bool
	showStats        bool
	additionalMemory uint
)

func main() {
	flags()
	openFiles()
	defer outputFile.Close()
	defer inputFile.Close()

	programFilename := flag.Arg(0)
	if programFilename == "" {
		printInfo()
		return
	}
	programFile, err := ioutil.ReadFile(programFilename)
	if err != nil {
		panic(err)
	}

	runProgram(string(programFile))
}

func runProgram(str string) {
	// Create a new program and execute it
	p := New(str, additionalMemory)
	p.InputReader = inputFile
	p.Debug = debug
	if showStats {
		p.Stats = newStats()
	}
	p.Exec()

	// Print executed program
	if outputFile != nil {
		fmt.Fprintln(outputFile, p.Ints.String())
	}

	// Show stats
	if showStats {
		fmt.Fprintln(os.Stderr, "Stats:")
		fmt.Fprintln(os.Stderr, p.Stats.String())
	}
}

func openFiles() {
	if outputFilename == "" {
		outputFile = nil
	} else if outputFilename == "-" {
		outputFile = os.Stdout
	} else {
		var err error
		outputFile, err = os.OpenFile(outputFilename, os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			panic(err)
		}
	}

	if inputFilename == "" {
		inputFile = os.Stdin
	} else {
		var err error
		inputFile, err = os.Open(inputFilename)
		if err != nil {
			panic(err)
		}
	}
}

func flags() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <flags> <filename>\n\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Flags:")
		flag.PrintDefaults()
	}
	flag.StringVar(&outputFilename, "output", "", "File to print the executed program to. Use '-' to print to console")
	flag.StringVar(&inputFilename, "input", "", "File to read input values from")
	flag.BoolVar(&debug, "debug", false, "Trace program execution via debug output")
	flag.BoolVar(&showStats, "stats", false, "Show statistics about execution duration and memory accesses")
	flag.UintVar(&additionalMemory, "mem", 42, "Number of ints that are allocated for the memory in addition "+
		"to the program. If a memory address outside the allocated memory is requested, the memory is increased by that offset")
	flag.Parse()
}

func printInfo() {
	fmt.Printf("INTCODE Computer %s on %s %s/%s\n\n", version, runtime.Version(), runtime.GOARCH, runtime.GOOS)
	fmt.Printf("Use \"%s -help\" for help\n", os.Args[0])
}
