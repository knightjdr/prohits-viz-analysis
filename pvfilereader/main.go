// Package main reads input files, reformats and concatenates
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pvfilereader/columnparser"
	"github.com/knightjdr/prohits-viz-analysis/pvfilereader/filter"
	"github.com/knightjdr/prohits-viz-analysis/pvfilereader/logmessage"
)

func main() {
	// command line flag arguments
	abundance := flag.String("abundance", "", "Abundance column")
	bait := flag.String("bait", "", "Bait column")
	baitList := flag.String("baitlist", "", "List of baits")
	control := flag.String("control", "", "Control column")
	fileList := flag.String("file", "", "Input files as comma-separated list")
	logFile := flag.String("log", "", "Log file")
	prey := flag.String("prey", "", "Prey column")
	preyList := flag.String("preylist", "", "List of preys")
	primaryFilter := flag.Float64("primary", 0, "Primary filter")
	score := flag.String("score", "", "Score column")
	scoreType := flag.String("scoreType", "lte", "Score type")
	flag.Parse()

	// exit if required args are missing
	argsError := false
	if *abundance == "" {
		logmessage.Write(*logFile, "No abundance column specified")
		argsError = true
	}
	if *bait == "" {
		logmessage.Write(*logFile, "No bait column specified")
		argsError = true
	}
	if *fileList == "" {
		logmessage.Write(*logFile, "No input file specified")
		argsError = true
	}
	if *prey == "" {
		logmessage.Write(*logFile, "No prey column specified")
		argsError = true
	}
	if *score == "" {
		logmessage.Write(*logFile, "No score column specified")
		argsError = true
	}
	if argsError == true {
		os.Exit(1)
	}

	// split filesArg to array of files
	files := strings.Split(*fileList, ",")

	// create mapping of column names to keys
	columnMap := map[string]string{
		"abundance": *abundance,
		"bait":      *bait,
		"control":   *control,
		"prey":      *prey,
		"score":     *score,
	}

	// split bait and prey lists
	baits := strings.Split(*baitList, ",")
	preys := strings.Split(*preyList, ",")

	// read needed columns from files
	parsedColumns, err := columnparser.Readfile(files, columnMap, *logFile)

	// filter rows
	filtered, err := filter.Data(parsedColumns, primaryFilter, baits, preys, *scoreType)
	fmt.Println(filtered)
}
