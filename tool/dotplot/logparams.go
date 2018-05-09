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
func LogParams(params typedef.Parameters) {
	logSlice := make([]string, 0)

	// Analysis type.
	logSlice = append(logSlice, fmt.Sprintf("Analysis type: %s", params.AnalysisType))
	logSlice = append(logSlice, "")

	// Files.
	logSlice = append(logSlice, "Files used")
	for _, file := range params.Files {
		logSlice = append(logSlice, fmt.Sprintf("- %s", file))
	}
	logSlice = append(logSlice, "")

	// Columns used.
	logSlice = append(logSlice, "Columns used")
	logSlice = append(logSlice, fmt.Sprintf("- abundance: %s", params.Abundance))
	logSlice = append(logSlice, fmt.Sprintf("- bait: %s", params.Bait))
	logSlice = append(logSlice, fmt.Sprintf("- prey: %s", params.Prey))
	logSlice = append(logSlice, fmt.Sprintf("- score: %s", params.Score))
	if params.Control != "" {
		logSlice = append(logSlice, fmt.Sprintf("- control: %s", params.Control))
	}
	if params.PreyLength != "" {
		logSlice = append(logSlice, fmt.Sprintf("- prey length: %s", params.PreyLength))
	}
	logSlice = append(logSlice, "")

	// Transformations.
	logSlice = append(logSlice, "Prey abundance transformations")
	if params.Control != "" {
		logSlice = append(logSlice, "- control subtraction was performed")
	}
	if params.PreyLength != "" {
		logSlice = append(logSlice, "- prey length normalization was performed")
	}
	if params.Normalization == "total" {
		logSlice = append(logSlice, "- bait normalization was performed using total abundance")
	} else if params.Normalization == "prey" {
		logSlice = append(
			logSlice,
			fmt.Sprintf("- bait normalization was performed using the prey: %s", params.NormalizationPrey),
		)
	}
	if params.LogBase != "" {
		logSlice = append(logSlice, fmt.Sprintf("- data was log-transformed with base %s", params.LogBase))
	}
	logSlice = append(logSlice, "")

	// Abundance
	logSlice = append(logSlice, "Abundance")
	logSlice = append(
		logSlice,
		fmt.Sprintf("- minimum abundance required: %s", strconv.FormatFloat(params.MinimumAbundance, 'f', -1, 64)),
	)
	logSlice = append(
		logSlice,
		fmt.Sprintf("- abundances were capped at %s for visualization", strconv.FormatFloat(params.MaximumAbundance, 'f', -1, 64)),
	)
	logSlice = append(logSlice, "")

	// Scoring.
	logSlice = append(logSlice, "Scoring")
	if params.ScoreType == "gte" {
		logSlice = append(logSlice, "- larger scores are better")
	} else {
		logSlice = append(logSlice, "- smaller scores are better")
	}
	logSlice = append(
		logSlice,
		fmt.Sprintf("- primary filter: %s", strconv.FormatFloat(params.PrimaryFilter, 'f', -1, 64)),
	)
	logSlice = append(
		logSlice,
		fmt.Sprintf("- secondary filter: %s", strconv.FormatFloat(params.SecondaryFilter, 'f', -1, 64)),
	)
	logSlice = append(logSlice, "")

	// Clustering.
	logSlice = append(logSlice, "Clustering")
	if params.Clustering == "biclustering" {
		logSlice = append(logSlice, "- biclustering was performed")
	} else if params.Clustering == "hierarchical" {
		logSlice = append(logSlice, "- hierarchical clustering was performed")
		logSlice = append(logSlice, fmt.Sprintf("- distance metric: %s", params.Distance))
		logSlice = append(logSlice, fmt.Sprintf("- linkage method: %s", params.ClusteringMethod))
	} else if params.BaitClustering == "baits" && params.PreyClustering != "preys" {
		logSlice = append(logSlice, "- preys were hierarchically clustered")
		logSlice = append(logSlice, fmt.Sprintf("- distance metric: %s", params.Distance))
		logSlice = append(logSlice, fmt.Sprintf("- linkage method: %s", params.ClusteringMethod))
	} else if params.PreyClustering == "preys" && params.BaitClustering != "baits" {
		logSlice = append(logSlice, "- baits were hierarchically clustered")
		logSlice = append(logSlice, fmt.Sprintf("- distance metric: %s", params.Distance))
		logSlice = append(logSlice, fmt.Sprintf("- linkage method: %s", params.ClusteringMethod))
	} else {
		logSlice = append(logSlice, "- no clustering was performed")
	}
	logSlice = append(logSlice, "")

	// Write log to file.
	logString := strings.Join(logSlice, "\r\n")
	afero.WriteFile(fs.Instance, "log.txt", []byte(logString), 0644)
	return
}
