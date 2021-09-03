// Package scv creates circular heat maps for conditions
package scv

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Generate is the entry point for correlation analysis.
func Generate(analysis *types.Analysis) {
	createFolders(analysis.Settings)

	idMaps := mapIDs(analysis)

	data := make(map[string]map[string]map[string]float64)
	addAbundance(data, analysis)
	addExpression("protein", data, idMaps, analysis.Settings)
	addExpression("rna", data, idMaps, analysis.Settings)
	addSpecificity(data, analysis)

	known := defineKnown(data, idMaps, analysis.Settings)

	legend := createLegend(data, analysis.Settings)

	plots := createPlotData(data, known, legend, analysis.Settings)

	createInteractive(plots, legend, analysis.Settings)
	writeData(plots, legend, analysis.Settings)
	writeMaps(idMaps, analysis.Settings)
	writeImages(plots, legend, analysis.Settings)
}
