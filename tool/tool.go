// Package tool passes analysis to the correct tool
package tool

import (
	"github.com/knightjdr/prohits-viz-analysis/tool/dotplot"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Start passes a dataset to the correct tool for analysis
func Start(dataset typedef.Dataset) {
	if dataset.Parameters.AnalysisType == "dotplot" {
		dotplot.Generate(dataset)
	}
}
