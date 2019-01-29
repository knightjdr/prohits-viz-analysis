// Package circheatmap draws circular heatmaps for all conditions in a file
package circheatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/tool/circheatmap/file"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Generate is the entry point for generating circular heatmaps and output files.
func Generate(dataset *typedef.Dataset) {
	// Create subfolders. Will panic if error.
	folders := make([]string, 0)
	folders = append(folders, []string{"interactive", "svg"}...)
	if dataset.Parameters.Png {
		folders = append(folders, "png")
	}
	helper.CreateFolders(folders)

	// Determine what metrics are being read from data file.
	metrics := readoutMetrics(dataset.Parameters)

	// Parse data
	conditionNames, readoutNames, conditionData := parseConditions(dataset.FileData, dataset.Parameters, metrics)

	// Set how real gene names should be mapped to condition names.
	mapping := parseMapFile(conditionNames, dataset.Parameters.ConditionMap)

	// Add "known" property to readouts
	conditionData = addKnown(conditionData, mapping, dataset.Parameters)

	// Add expression data to readouts
	conditionData, metrics = addExpression(conditionData, readoutNames, metrics, dataset.Parameters)

	// Get order for metrics and create segment settings
	metricOrder, settings := segmentSettings(metrics, dataset.Parameters)

	// Format condition data as plots for svg and interactive file.
	plots := make([]typedef.CircHeatmapPlot, len(conditionNames))
	for index, condition := range conditionNames {
		plots[index] = formatCondition(condition, conditionData[condition], dataset.Parameters.Known, metricOrder, metrics)
	}

	// Create interactive file
	file.Interactive(plots, dataset.Parameters, settings)

	return
}
