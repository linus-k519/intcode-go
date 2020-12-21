package main

import (
	"flag"
	"fmt"
	log "github.com/linus-k519/llog"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"
)

func run(file string) {
	// Clean newlines and comments
	program := NewProgram(file, additionalMemory)

	// Execute program
	startTime := time.Now()
	program.Exec()
	execTime := time.Since(startTime)

	if outputFile != nil {
		fmt.Fprintln(outputFile, program.StringInts())
	}

	if showStats {
		fmt.Println(NewStats(program.OperationCount, execTime).String())
	}
}

var (
	outputFile       *os.File
	outputFilename   string
	trace            bool
	showStats        bool
	additionalMemory uint
)

const version = "v9.2"

func main() {
	log.Config(0)
	flags()

	outputFile = openOutputFile()
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

	run(string(programFile))
}

func openOutputFile() *os.File {
	if outputFilename == "" {
		return nil
	}

	if strings.ToLower(outputFilename) == "stdout" {
		return os.Stdout
	}
	outFile, err := os.OpenFile(outputFilename, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	return outFile
}

func flags() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <flags> <filename>\n\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Flags:")
		flag.PrintDefaults()
	}
	flag.StringVar(&outputFilename, "output", "", "File to print the executed program to. Use 'stdout' to print to console")
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
