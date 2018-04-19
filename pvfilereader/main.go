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
	baitArg := flag.String("bait", "", "Bait column")
	controlArg := flag.String("control", "", "Control column")
	filesArg := flag.String("file", "", "Input files as comma-separated list")
	logFile := flag.String("log", "", "Log file")
	preyArg := flag.String("prey", "", "Prey column")
	flag.Parse()

	// exit if required args are missing
	argsError := false
	if *filesArg == "" {
		logmessage.Write(*logFile, "No input file specified")
		argsError = true
	}
	if *baitArg == "" {
		logmessage.Write(*logFile, "No bait column")
		argsError = true
	}
	if *preyArg == "" {
		logmessage.Write(*logFile, "No prey column")
		argsError = true
	}
	if argsError == true {
		os.Exit(1)
	}

	// split filesArg to array of files
	files := strings.Split(*filesArg, ",")

	// creating mapping of column names to keys
	columnMap := map[string]string{
		"bait":    *baitArg,
		"control": *controlArg,
		"prey":    *preyArg,
	}

	// read files and write columns to new file
	columnparser.Readfile(files, columnMap, *logFile)
	fmt.Println("file:", files)
}
