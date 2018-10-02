package dotplot

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// LogParams writes parameters used to a log file.
func LogParams(parameters typedef.Parameters) {
	logSlice := make([]string, 0)

	// Analysis type.
	logSlice = append(logSlice, fmt.Sprintf("Analysis type: %s", parameters.AnalysisType))
	logSlice = append(logSlice, "")

	// Files.
	logSlice = append(logSlice, "Files used")
	for _, file := range parameters.Files {
		logSlice = append(logSlice, fmt.Sprintf("- %s", file))
	}
	logSlice = append(logSlice, "")

	// Columns used.
	logSlice = append(logSlice, "Columns used")
	logSlice = append(logSlice, fmt.Sprintf("- abundance: %s", parameters.Abundance))
	logSlice = append(logSlice, fmt.Sprintf("- bait: %s", parameters.Bait))
	logSlice = append(logSlice, fmt.Sprintf("- prey: %s", parameters.Prey))
	logSlice = append(logSlice, fmt.Sprintf("- score: %s", parameters.Score))
	if parameters.Control != "" {
		logSlice = append(logSlice, fmt.Sprintf("- control: %s", parameters.Control))
	}
	if parameters.PreyLength != "" {
		logSlice = append(logSlice, fmt.Sprintf("- prey length: %s", parameters.PreyLength))
	}
	logSlice = append(logSlice, "")

	// Transformations.
	logSlice = append(logSlice, "Prey abundance transformations")
	if parameters.Control != "" {
		logSlice = append(logSlice, "- control subtraction was performed")
	}
	if parameters.PreyLength != "" {
		logSlice = append(logSlice, "- prey length normalization was performed")
	}
	if parameters.Normalization == "total" {
		logSlice = append(logSlice, "- bait normalization was performed using total abundance")
	} else if parameters.Normalization == "prey" {
		logSlice = append(
			logSlice,
			fmt.Sprintf("- bait normalization was performed using the prey: %s", parameters.NormalizationPrey),
		)
	}
	if parameters.LogBase != "" {
		logSlice = append(logSlice, fmt.Sprintf("- data was log-transformed with base %s", parameters.LogBase))
	}
	logSlice = append(logSlice, "")

	// Abundance
	logSlice = append(logSlice, "Abundance")
	logSlice = append(
		logSlice,
		fmt.Sprintf("- minimum abundance required: %s", strconv.FormatFloat(parameters.MinAbundance, 'f', -1, 64)),
	)
	logSlice = append(
		logSlice,
		fmt.Sprintf("- abundances were capped at %s for visualization", strconv.FormatFloat(parameters.AbundanceCap, 'f', -1, 64)),
	)
	logSlice = append(logSlice, "")

	// Scoring.
	logSlice = append(logSlice, "Scoring")
	if parameters.ScoreType == "gte" {
		logSlice = append(logSlice, "- larger scores are better")
	} else {
		logSlice = append(logSlice, "- smaller scores are better")
	}
	logSlice = append(
		logSlice,
		fmt.Sprintf("- primary filter: %s", strconv.FormatFloat(parameters.PrimaryFilter, 'f', -1, 64)),
	)
	logSlice = append(
		logSlice,
		fmt.Sprintf("- secondary filter: %s", strconv.FormatFloat(parameters.SecondaryFilter, 'f', -1, 64)),
	)
	logSlice = append(logSlice, "")

	// Clustering.
	logSlice = append(logSlice, "Clustering")
	if parameters.Clustering == "biclustering" {
		logSlice = append(logSlice, "- biclustering was performed")
	} else if parameters.Clustering == "hierarchical" {
		logSlice = append(logSlice, "- hierarchical clustering was performed")
		logSlice = append(logSlice, fmt.Sprintf("- distance metric: %s", parameters.Distance))
		logSlice = append(logSlice, fmt.Sprintf("- linkage method: %s", parameters.ClusteringMethod))
	} else if parameters.BaitClustering == "baits" && parameters.PreyClustering != "preys" {
		logSlice = append(logSlice, "- preys were hierarchically clustered")
		logSlice = append(logSlice, fmt.Sprintf("- distance metric: %s", parameters.Distance))
		logSlice = append(logSlice, fmt.Sprintf("- linkage method: %s", parameters.ClusteringMethod))
	} else if parameters.PreyClustering == "preys" && parameters.BaitClustering != "baits" {
		logSlice = append(logSlice, "- baits were hierarchically clustered")
		logSlice = append(logSlice, fmt.Sprintf("- distance metric: %s", parameters.Distance))
		logSlice = append(logSlice, fmt.Sprintf("- linkage method: %s", parameters.ClusteringMethod))
	} else {
		logSlice = append(logSlice, "- no clustering was performed")
	}
	logSlice = append(logSlice, "")

	// Write log to file.
	logString := strings.Join(logSlice, "\r\n")
	afero.WriteFile(fs.Instance, "log.txt", []byte(logString), 0644)
	return
}
