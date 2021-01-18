// Package specificity creates scatter plots for condition showing readout specificty.
package specificity

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Generate is the entry point for correlation analysis.
func Generate(analysis *types.Analysis) {
	createFolders(analysis.Settings)

	calculateSpecificity(analysis)

	createLegend(analysis.Settings)
}
