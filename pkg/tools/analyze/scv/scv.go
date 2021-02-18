// Package scv creates circular heat maps for conditions
package scv

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Generate is the entry point for correlation analysis.
func Generate(analysis *types.Analysis) {
	createFolders(analysis.Settings)

	mapIDs(analysis)

	/* specificity := calculateSpecificity(analysis)
	plotByCondition := formatDataForPlot(specificity, analysis.Settings)

	createLegend(analysis.Settings)
	createInteractive(plotByCondition, analysis.Settings)
	writeData(specificity, analysis.Settings)
	writeImages(plotByCondition, analysis.Settings) */
}
