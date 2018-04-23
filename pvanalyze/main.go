// Package main takes input files and parameters and runs the analysis
package main

import (
	"flag"
	"os"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/errorcheck"
	"github.com/knightjdr/prohits-viz-analysis/filereader/columnparser"
	"github.com/knightjdr/prohits-viz-analysis/filereader/filter"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/transform"
)

func main() {
	// command line flag arguments
	abundance := flag.String("abundance", "", "Abundance column")
	analysisType := flag.String("analysisType", "", "Analysis type, one of: baitbait, correlation, dotplot or specificity")
	bait := flag.String("bait", "", "Bait column")
	baitList := flag.String("baitList", "", "List of baits")
	control := flag.String("control", "", "Control column")
	fileList := flag.String("fileList", "", "Input files as comma-separated list")
	logFile := flag.String("logFile", "", "Log file")
	prey := flag.String("prey", "", "Prey column")
	preyLength := flag.String("preyLength", "", "Prey length column")
	preyList := flag.String("preyList", "", "List of preys")
	primaryFilter := flag.Float64("primaryFilter", 0, "Primary filter")
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
		"abundance":  *abundance,
		"bait":       *bait,
		"control":    *control,
		"prey":       *prey,
		"preyLength": *preyLength,
		"score":      *score,
	}

	// split bait and prey lists
	baits := []string{}
	if *baitList != "" {
		baits = strings.Split(*baitList, ",")
	}
	preys := []string{}
	if *preyList != "" {
		preys = strings.Split(*preyList, ",")
	}

	// read needed columns from files
	parsedColumns, err := columnparser.ReadFile(files, columnMap, *logFile)
	if err != nil {
		os.Exit(1)
	}

	// filter rows
	filtered, err := filter.Data(parsedColumns, *primaryFilter, baits, preys, *scoreType, *logFile)
	if err != nil {
		os.Exit(1)
	}

	// check for common errors in filtered data that result from incorrect input format
	err = errorcheck.Required(filtered, *analysisType, *control, *preyLength, *logFile)
	if err != nil {
		os.Exit(1)
	}

	transform.Preys(filtered, control)
}
