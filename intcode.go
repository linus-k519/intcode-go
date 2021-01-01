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
	executedProgramFilename string
	executedProgramFile     *os.File
	outputFilename          string
	outputFile              *os.File
	inputFilename           string
	inputFile               *os.File
	showDebug               bool
	showStats               bool
	additionalMemory        uint
)

func main() {
	flags()
	openFiles()
	defer executedProgramFile.Close()
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

// runProgram creates a new program, executes it, prints the executed program and shows stats
func runProgram(str string) {
	// Create a new program and execute it
	p := New(str, additionalMemory)
	p.InputReader = inputFile
	p.Debug = showDebug
	if showStats {
		p.Stats = newStats()
	}
	p.Exec()

	// Print executed program
	if executedProgramFile != nil {
		fmt.Fprintln(executedProgramFile, p.Ints.String())
	}

	// Show stats
	if showStats {
		fmt.Fprintln(os.Stderr, "Stats:")
		fmt.Fprintln(os.Stderr, p.Stats.String())
	}
}

// openFiles opens the executed program file, the output file and the input file
// specified by their filename variable.
func openFiles() {
	// Open executed program file
	if executedProgramFilename == "" {
		executedProgramFile = nil
	} else if executedProgramFilename == "-" {
		executedProgramFile = os.Stdout
	} else {
		var err error
		executedProgramFile, err = os.OpenFile(executedProgramFilename, os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			panic(err)
		}
	}

	// Open output file
	if outputFilename == "" {
		outputFile = os.Stdout
	} else {
		var err error
		outputFile, err = os.OpenFile(outputFilename, os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			panic(err)
		}
	}

	// Open input file
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

// flags parsed the program arguments into the variables.
func flags() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <flags> <filename>\n\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Flags:")
		flag.PrintDefaults()
	}
	flag.StringVar(&executedProgramFilename, "executed-program", "", "File to print the executed program to. Use '-' to print to console")
	flag.StringVar(&inputFilename, "input", "", "File to read input values from")
	flag.StringVar(&outputFilename, "output", "", "File to print output values to")
	flag.BoolVar(&showDebug, "showDebug", false, "Trace program execution via showDebug output")
	flag.BoolVar(&showStats, "stats", false, "Show statistics about execution duration and memory accesses")
	flag.UintVar(&additionalMemory, "mem", 42, "Number of ints that are allocated for the memory in addition "+
		"to the program. If a memory address outside the allocated memory is requested, the memory is increased by that offset")
	flag.Parse()
}

// printInfo prints the version number and help for this intcode interpreter.
func printInfo() {
	fmt.Printf("INTCODE Computer %s on %s %s/%s\n\n", version, runtime.Version(), runtime.GOARCH, runtime.GOOS)
	fmt.Printf("Use \"%s -help\" for help\n", os.Args[0])
}
