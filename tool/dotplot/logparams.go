package dotplot

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/helper"
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
	fileNames := helper.Filename(parameters.Files)
	for _, file := range fileNames {
		logSlice = append(logSlice, fmt.Sprintf("- %s", file))
	}
	logSlice = append(logSlice, "")

	// Columns used.
	logSlice = append(logSlice, "Columns used")
	logSlice = append(logSlice, fmt.Sprintf("- abundance: %s", parameters.Abundance))
	logSlice = append(logSlice, fmt.Sprintf("- condition: %s", parameters.Condition))
	logSlice = append(logSlice, fmt.Sprintf("- readout: %s", parameters.Readout))
	logSlice = append(logSlice, fmt.Sprintf("- score: %s", parameters.Score))
	if parameters.Control != "" {
		logSlice = append(logSlice, fmt.Sprintf("- control: %s", parameters.Control))
	}
	if parameters.ReadoutLength != "" {
		logSlice = append(logSlice, fmt.Sprintf("- readout length: %s", parameters.ReadoutLength))
	}
	logSlice = append(logSlice, "")

	// Transformations.
	logSlice = append(logSlice, "Readout abundance transformations")
	if parameters.Control != "" {
		logSlice = append(logSlice, "- control subtraction was performed")
	}
	if parameters.ReadoutLength != "" {
		logSlice = append(logSlice, "- readout length normalization was performed")
	}
	if parameters.Normalization == "total" {
		logSlice = append(logSlice, "- condition normalization was performed using total abundance")
	} else if parameters.Normalization == "readout" {
		logSlice = append(
			logSlice,
			fmt.Sprintf("- condition normalization was performed using the readout: %s", parameters.NormalizationReadout),
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
		if parameters.BiclusteringApprox {
			logSlice = append(logSlice, "- approximate biclustering was performed")
		} else {
			logSlice = append(logSlice, "- biclustering was performed")
		}
	} else if parameters.Clustering == "hierarchical" {
		logSlice = append(logSlice, "- hierarchical clustering was performed")
		logSlice = append(logSlice, fmt.Sprintf("- distance metric: %s", parameters.Distance))
		logSlice = append(logSlice, fmt.Sprintf("- linkage method: %s", parameters.ClusteringMethod))
	} else if parameters.ConditionClustering == "none" && parameters.ReadoutClustering == "readouts" {
		logSlice = append(logSlice, "- readouts were hierarchically clustered")
		logSlice = append(logSlice, fmt.Sprintf("- distance metric: %s", parameters.Distance))
		logSlice = append(logSlice, fmt.Sprintf("- linkage method: %s", parameters.ClusteringMethod))
	} else if parameters.ReadoutClustering == "none" && parameters.ConditionClustering == "conditions" {
		logSlice = append(logSlice, "- conditions were hierarchically clustered")
		logSlice = append(logSlice, fmt.Sprintf("- distance metric: %s", parameters.Distance))
		logSlice = append(logSlice, fmt.Sprintf("- linkage method: %s", parameters.ClusteringMethod))
	} else {
		logSlice = append(logSlice, "- no clustering was performed")
	}
	if parameters.Clustering != "biclustering" {
		if parameters.ClusteringOptimize {
			logSlice = append(logSlice, "- leaf clusters were optimized")
		} else {
			logSlice = append(logSlice, "- leaf clusters were not optimized")
		}
	}

	logSlice = append(logSlice, "")

	// Write log to file.
	logString := strings.Join(logSlice, "\r\n")
	afero.WriteFile(fs.Instance, "log.txt", []byte(logString), 0644)
	return
}
