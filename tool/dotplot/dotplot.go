// Package dotplot clusters baits and preys for visualization as a dotplot
package dotplot

import (
	"os"
	"path/filepath"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/types"
)

// Generate is the entry point for generating dotplots and output files.
func Generate(dataset types.Dataset) {
	// Create subfolders. Panic if error.
	cytoscapePath := filepath.Join(".", "cytoscape")
	err := os.MkdirAll(cytoscapePath, os.ModePerm)
	logmessage.CheckError(err, true)
	otherPath := filepath.Join(".", "other")
	err = os.MkdirAll(otherPath, os.ModePerm)
	logmessage.CheckError(err, true)

	// Initiate clustering method.
	if dataset.Params.Clustering == "hierarchical" {
		Hierarchical(dataset)
	}
	return
}