package main

import (
	"errors"
	"flag"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// ParseFlags parses the command line arguments
func ParseFlags() (columnMap map[string]string, params typedef.Parameters, err error) {
	// Command line flag arguments.
	abundance := flag.String("abundance", "", "Abundance column")
	analysisType := flag.String("analysisType", "", "Analysis type, one of: baitbait, correlation, dotplot or specificity")
	bait := flag.String("bait", "", "Bait column")
	baitClustering := flag.String("baitClustering", "", "Should baits be clustered: one of baits or none")
	baitList := flag.String("baitList", "", "List of baits")
	biclusteringApprox := flag.Bool("biclusteringApprox", false, "Perform approximate biclustering")
	clustering := flag.String("clustering", "", "Clustering type")
	clusteringMethod := flag.String("clusteringMethod", "", "Clustering type")
	edgeColor := flag.String("edgeColor", "", "Edge color")
	fillColor := flag.String("fillColor", "", "Fill color")
	control := flag.String("control", "", "Control column")
	distance := flag.String("distance", "", "Distance metric")
	fileList := flag.String("fileList", "", "Input files as comma-separated list")
	invert := flag.Bool("invert", false, "Invert fill color")
	logBase := flag.String("logBase", "", "Base for log transformation")
	maximumAbundance := flag.Float64("maximumAbundance", 50, "Maximum abundance")
	minimumAbundance := flag.Float64("minimumAbundance", 0, "Minimum abundance")
	normalization := flag.String("normalization", "", "Normalization type")
	normalizationPrey := flag.String("normalizationPrey", "", "Prey to use for normalization")
	pdf := flag.Bool("pdf", false, "Generate pdf files")
	png := flag.Bool("png", false, "Generate png files")
	prey := flag.String("prey", "", "Prey column")
	preyClustering := flag.String("preyClustering", "", "Should preys be clustered: one of none or preys")
	preyLength := flag.String("preyLength", "", "Prey length column")
	preyList := flag.String("preyList", "", "List of preys")
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
	if *bait == "" {
		logmessage.Write("No bait column specified")
		argsError = true
	}
	if *fileList == "" {
		logmessage.Write("No input file specified")
		argsError = true
	}
	if *prey == "" {
		logmessage.Write("No prey column specified")
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
		"abundance":  *abundance,
		"bait":       *bait,
		"control":    *control,
		"prey":       *prey,
		"preyLength": *preyLength,
		"score":      *score,
	}

	// Split bait and prey lists.
	baits := []string{}
	if *baitList != "" {
		baits = strings.Split(*baitList, ",")
	}
	preys := []string{}
	if *preyList != "" {
		preys = strings.Split(*preyList, ",")
	}

	// Create parameter struct.
	params = typedef.Parameters{
		Abundance:          *abundance,
		AnalysisType:       *analysisType,
		Bait:               *bait,
		BaitClustering:     *baitClustering,
		BaitList:           baits,
		BiclusteringApprox: *biclusteringApprox,
		Clustering:         *clustering,
		ClusteringMethod:   *clusteringMethod,
		EdgeColor:          *edgeColor,
		FillColor:          *fillColor,
		Control:            *control,
		Distance:           *distance,
		Files:              files,
		Invert:             *invert,
		LogBase:            *logBase,
		MaximumAbundance:   *maximumAbundance,
		MinimumAbundance:   *minimumAbundance,
		Normalization:      *normalization,
		NormalizationPrey:  *normalizationPrey,
		Pdf:                *pdf,
		Png:                *png,
		Prey:               *prey,
		PreyClustering:     *preyClustering,
		PreyLength:         *preyLength,
		PreyList:           preys,
		PrimaryFilter:      *primaryFilter,
		Score:              *score,
		ScoreType:          *scoreType,
		SecondaryFilter:    *secondaryFilter,
		WriteDistance:      *writeDistance,
		WriteDotplot:       *writeDotplot,
		WriteHeatmap:       *writeHeatmap,
	}

	return
}
