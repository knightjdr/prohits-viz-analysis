// Package main reads input files, reformats and concatenates
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pvfilereader/columnparser"
	"github.com/knightjdr/prohits-viz-analysis/pvfilereader/logmessage"
)

func main() {
	// command line flag arguments
	filesArg := flag.String("file", "", "Input files as comma-separated list")
	headerArg := flag.String("header", "", "Header columns to use as comma-separated list")
	logFile := flag.String("log", "", "Log file")
	flag.Parse()

	// exit if no file or header columns specified
	argsError := false
	if *filesArg == "" {
		logmessage.Write(*logFile, "No input file specified")
		argsError = true
	}
	if *headerArg == "" {
		logmessage.Write(*logFile, "No headers specified")
		argsError = true
	}
	if argsError == true {
		os.Exit(1)
	}

	// split filesArg to array of files
	files := strings.Split(*filesArg, ",")

	// split headerArg to array of column names
	columns := strings.Split(*headerArg, ",")

	// read files and write columns to new file
	columnparser.Readfile(files, columns, *logFile)
	fmt.Println("file:", files)
	fmt.Println("columns:", columns)
}
