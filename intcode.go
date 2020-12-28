package main

import (
	"flag"
	"fmt"
	log "github.com/linus-k519/logo"
	"io/ioutil"
	"os"
	"runtime"
)

const version = "v9.2"

var (
	outputFilename   string
	outputFile       *os.File
	trace            bool
	showStats        bool
	additionalMemory uint
)

func main() {
	log.Config(0)
	flags()

	outputFile = openOutputFile(outputFilename)
	defer outputFile.Close()

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
	p := NewWithAdditionalMemory(str, additionalMemory)
	// TODO set p.Trace
	p.Exec()

	// Print executed program
	if outputFile != nil {
		fmt.Fprintln(outputFile, p.Ints.String())
	}

	// Show stats
	if showStats {
		fmt.Println("Stats:")
		fmt.Println(p.Stats.String())
	}
}

func openOutputFile(filename string) *os.File {
	if filename == "" {
		return nil
	}
	if filename == "-" {
		return os.Stdout
	}
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	return file
}

func flags() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <flags> <filename>\n\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Flags:")
		flag.PrintDefaults()
	}
	flag.StringVar(&outputFilename, "output", "", "File to print the executed program to. Use '-' to print to console")
	flag.BoolVar(&trace, "trace", false, "Trace program execution via debug output")
	flag.BoolVar(&showStats, "stats", false, "Show showStats")
	flag.UintVar(&additionalMemory, "mem", 42, "Number of ints that are allocated for the memory in addition "+
		"to the program. If a memory address outside the allocated memory is requested, the memory is increased by that offset")
	flag.Parse()
}

func printInfo() {
	fmt.Printf("INTCODE Computer %s on %s %s/%s\n\n", version, runtime.Version(), runtime.GOARCH, runtime.GOOS)
	fmt.Printf("Use \"%s -help\" for help\n", os.Args[0])
}
