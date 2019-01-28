// Package tool passes analysis to the correct tool
package tool

import (
	"github.com/knightjdr/prohits-viz-analysis/tool/circheatmap"
	"github.com/knightjdr/prohits-viz-analysis/tool/dotplot"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Start passes a dataset to the correct tool for analysis
func Start(dataset *typedef.Dataset) {
	// Write log.
	logParams(dataset.Parameters)

	if dataset.Parameters.AnalysisType == "dotplot" {
		dotplot.Generate(dataset)
	} else if dataset.Parameters.AnalysisType == "circheatmap" {
		circheatmap.Generate(dataset)
	}
}
