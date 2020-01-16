// Package transform will adjust readout values to the user's requirements.
//
// Readouts can be adjusted by (in this order):
// 1) control values (must be a pipe-separated list)
// 2) readout length
// 3) normalized across conditions
// 4) log transformation
// 5) mock condition abundances when missing.
package transform

import "github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"

// Abundance is the entry point for readout abundance transformations.
func Abundance(analysis *types.Analysis) {
	controlSubtract(analysis)
	adjustByReadoutLength(analysis)
	normalize(analysis)
	logTransform(analysis)
	mockConditionAbundance(analysis)
}
