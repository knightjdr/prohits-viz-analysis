// Package filter filters data based on conditions, readouts and score.
package filter

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

// Process filters first by condition and readout, then score.
func Process(analysis *types.Analysis) {
	filterByConditionsAndReadouts(analysis)
	filterByAbundanceAndScore(analysis)
}
