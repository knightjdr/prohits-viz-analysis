package main

import (
	"errors"
	"flag"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// ParseFlags parses the command line arguments
func ParseFlags() (columnMap map[string]string, parameters typedef.Parameters, err error) {
	// Command line flag arguments.
	abundance := flag.String("abundance", "", "Abundance column")
	abundanceCap := flag.Float64("abundanceCap", 50, "Maximum abundance")
	analysisType := flag.String("analysisType", "", "Analysis type, one of: conditioncondition, correlation, dotplot or specificity")
	condition := flag.String("condition", "", "Condition column")
	conditionClustering := flag.String("conditionClustering", "", "Should conditions be clustered: one of conditions or none")
	conditionList := flag.String("conditionList", "", "List of conditions")
	biclusteringApprox := flag.Bool("biclusteringApprox", false, "Perform approximate biclustering")
	clustering := flag.String("clustering", "", "Clustering type")
	clusteringMethod := flag.String("clusteringMethod", "", "Clustering type")
	control := flag.String("control", "", "Control column")
	distance := flag.String("distance", "", "Distance metric")
	edgeColor := flag.String("edgeColor", "", "Edge color")
	fileList := flag.String("fileList", "", "Input files as comma-separated list")
	fillColor := flag.String("fillColor", "", "Fill color")
	invertColor := flag.Bool("invertColor", false, "Invert fill color")
	logBase := flag.String("logBase", "", "Base for log transformation")
	minAbundance := flag.Float64("minAbundance", 0, "Minimum abundance")
	normalization := flag.String("normalization", "", "Normalization type")
	normalizationReadout := flag.String("normalizationReadout", "", "Readout to use for normalization")
	pdf := flag.Bool("pdf", false, "Generate pdf files")
	png := flag.Bool("png", false, "Generate png files")
	readout := flag.String("readout", "", "Readout column")
	readoutClustering := flag.String("readoutClustering", "", "Should readouts be clustered: one of none or readouts")
	readoutLength := flag.String("readoutLength", "", "Readout length column")
	readoutList := flag.String("readoutList", "", "List of readouts")
	primaryFilter := flag.Float64("primaryFilter", 0, "Primary filter")
	score := flag.String("score", "", "Score column")
	scoreType := flag.String("scoreType", "lte", "Score type")
	secondaryFilter := flag.Float64("secondaryFilter", 0, "Secondary filter")
	writeDistance := flag.Bool("writeDistance", false, "Generate SVG files for distance matrix")
	writeDotplot := flag.Bool("writeDotplot", false, "Generate SVG file for dotplot")
	writeHeatmap := flag.Bool("writeHeatmap", false, "Generate SVG file for heatmap")
	flag.Parse()

	// Exit if required args are missing.
	argsError := false
	if *abundance == "" {
		logmessage.Write("No abundance column specified")
		argsError = true
	}
	if *analysisType == "" {
		logmessage.Write("No analysis type specified")
		argsError = true
	}
	if *condition == "" {
		logmessage.Write("No condition column specified")
		argsError = true
	}
	if *fileList == "" {
		logmessage.Write("No input file specified")
		argsError = true
	}
	if *readout == "" {
		logmessage.Write("No readout column specified")
		argsError = true
	}
	if *score == "" {
		logmessage.Write("No score column specified")
		argsError = true
	}
	if argsError == true {
		err = errors.New("Missing required argument(s)")
		return
	}

	// Split filesArg to array of files.
	files := strings.Split(*fileList, ",")

	// Create mapping of column names to keys.
	columnMap = map[string]string{
		"abundance":     *abundance,
		"condition":     *condition,
		"control":       *control,
		"readout":       *readout,
		"readoutLength": *readoutLength,
		"score":         *score,
	}

	// Split condition and readout lists.
	conditions := []string{}
	if *conditionList != "" {
		conditions = strings.Split(*conditionList, ",")
	}
	readouts := []string{}
	if *readoutList != "" {
		readouts = strings.Split(*readoutList, ",")
	}

	// Create parameter struct.
	parameters = typedef.Parameters{
		Abundance:            *abundance,
		AbundanceCap:         *abundanceCap,
		AnalysisType:         *analysisType,
		BiclusteringApprox:   *biclusteringApprox,
		Clustering:           *clustering,
		ClusteringMethod:     *clusteringMethod,
		Condition:            *condition,
		ConditionClustering:  *conditionClustering,
		ConditionList:        conditions,
		EdgeColor:            *edgeColor,
		Control:              *control,
		Distance:             *distance,
		FillColor:            *fillColor,
		Files:                files,
		InvertColor:          *invertColor,
		LogBase:              *logBase,
		MinAbundance:         *minAbundance,
		Normalization:        *normalization,
		NormalizationReadout: *normalizationReadout,
		Pdf:                  *pdf,
		Png:                  *png,
		Readout:              *readout,
		ReadoutClustering:    *readoutClustering,
		ReadoutLength:        *readoutLength,
		ReadoutList:          readouts,
		PrimaryFilter:        *primaryFilter,
		Score:                *score,
		ScoreType:            *scoreType,
		SecondaryFilter:      *secondaryFilter,
		WriteDistance:        *writeDistance,
		WriteDotplot:         *writeDotplot,
		WriteHeatmap:         *writeHeatmap,
	}

	return
}
