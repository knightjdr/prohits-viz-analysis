// Package scv creates circular heat maps for conditions
package scv

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Generate is the entry point for correlation analysis.
func Generate(analysis *types.Analysis) {
	createFolders(analysis.Settings)

	idMaps := mapIDs(analysis)

	data := make(map[string]map[string]map[string]float64, 0)
	addAbundance(data, analysis)
	addExpression("protein", data, idMaps, analysis.Settings)
	addExpression("rna", data, idMaps, analysis.Settings)
	addSpecificity(data, analysis)

	defineKnown(data, idMaps, analysis.Settings)

	createLegend(analysis.Settings)

	writeMaps(idMaps, analysis.Settings)

	/*
		specificity := calculateSpecificity(analysis)

		createLegend(analysis.Settings)
		createInteractive(plotByCondition, analysis.Settings)
		writeData(specificity, analysis.Settings)
		writeImages(plotByCondition, analysis.Settings)
	*/
}
