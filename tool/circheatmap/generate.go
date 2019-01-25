// Package circheatmap draws circular heatmaps for all conditions in a file
package circheatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/tool/dotplot/file"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Generate is the entry point for generating circular heatmaps and output files.
func Generate(dataset *typedef.Dataset) {
	// Determine folders to create.
	folders := make([]string, 0)
	folders = append(folders, []string{"interactive", "svg"}...)
	if dataset.Parameters.Png {
		folders = append(folders, "png")
	}

	// Create subfolders. Will panic if error.
	helper.CreateFolders(folders)

	// Write log.
	file.LogParams(dataset.Parameters)

	// Determine what metrics are being read from data file.
	readoutMetrics := metrics(dataset.Parameters)

	// Parse data
	conditionNames, readoutNames, conditionData := parseConditions(dataset.FileData, dataset.Parameters, readoutMetrics)

	// Set how real gene names should be mapped to condition names.
	mapping := parseMapFile(conditionNames, dataset.Parameters.ConditionMap)

	// Add known metric if requested.
	if dataset.Parameters.Known {
		// Get known readout data
		known := parseKnownReadouts(mapping, dataset.Parameters.KnownFile, dataset.Parameters.Species)

		// Add known status to readouts
		for condition, readouts := range conditionData {
			for readout := range readouts {
				if known[condition][readout] {
					conditionData[condition][readout]["known"] = 1
				} else {
					conditionData[condition][readout]["known"] = 0
				}
			}
		}
	}

	// Add tissue expression if requested.
	if len(dataset.Parameters.Tissues) > 0 {
		// Add tissue names to readout metrics
		for _, tissue := range dataset.Parameters.Tissues {
			readoutMetrics[tissue] = tissue
		}

		// Get expression data for readouts
		expression := parseTissues(readoutNames, dataset.Parameters.TissueFile, dataset.Parameters.Tissues)

		// Add expression data to condition data.
		for condition, readouts := range conditionData {
			for readout := range readouts {
				for _, tissue := range dataset.Parameters.Tissues {
					if expression[readout][tissue] > 0 {
						conditionData[condition][readout][tissue] = expression[readout][tissue]
					} else {
						conditionData[condition][readout][tissue] = 0
					}
				}
			}
		}
	}

	return
}
