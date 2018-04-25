// Package tools passes analysis to the correct tool
package tools

import (
	"github.com/knightjdr/prohits-viz-analysis/tools/dotplot"
	"github.com/knightjdr/prohits-viz-analysis/types"
)

// Start passes a dataset to the correct tool for analysis
func Start(dataset types.Dataset) {
	if dataset.Params.AnalysisType == "dotplot" {
		dotplot.Generate(dataset)
	}
}
