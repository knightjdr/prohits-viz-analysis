// Package specificity creates scatter plots for condition showing readout specificty.
package specificity

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/specificity"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Generate is the entry point for correlation analysis.
func Generate(analysis *types.Analysis) {
	createFolders(analysis.Settings)

	specificity := specificity.Calculate(analysis)
	plotByCondition := formatDataForPlot(specificity, analysis.Settings)

	createLegend(analysis.Settings)
	createInteractive(plotByCondition, analysis.Settings)
	writeData(specificity, analysis.Settings)
	writeImages(plotByCondition, analysis.Settings)
}
