// Package dotplot clusters baits and preys for visualization as a dotplot
package dotplot

import "github.com/knightjdr/prohits-viz-analysis/types"

// Generate is the entry point for generating dotplots.
func Generate(dataset types.Dataset) {
	// generate bait-prey table
	matrix, baitList, preyList := BaitPreyMatrix(dataset.Data)

	// Generate prey distance matrix.
}
