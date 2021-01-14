// Package cc create a scatter plot between two conditions.
package cc

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Generate is the entry point for correlation analysis.
func Generate(analysis *types.Analysis) {
	createFolders(analysis.Settings)

	scatterData := definePoints(analysis)

	createLegend(analysis.Settings)
	createInteractive(scatterData, analysis.Settings)
	writeData(scatterData, analysis.Settings)
}
