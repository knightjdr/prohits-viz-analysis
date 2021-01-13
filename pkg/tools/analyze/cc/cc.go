// Package cc create a scatter plot between two conditions.
package cc

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Generate is the entry point for correlation analysis.
func Generate(analysis *types.Analysis) {
	createFolders(analysis.Settings)

	fmt.Print(analysis.Data)

	/*
		createCorrelationImages(corrConditions, corrReadouts, analysis.Settings)
		writeTrees(corrConditions, corrReadouts, analysis.Settings) */
}
