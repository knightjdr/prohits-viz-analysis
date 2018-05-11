// Package dotplot clusters baits and preys for visualization as a dotplot
package dotplot

import (
	"os"
	"path/filepath"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Generate is the entry point for generating dotplots and output files.
func Generate(dataset typedef.Dataset) {
	// Create subfolders. Panic if error.
	cytoscapePath := filepath.Join(".", "cytoscape")
	err := os.MkdirAll(cytoscapePath, os.ModePerm)
	logmessage.CheckError(err, true)
	interactivePath := filepath.Join(".", "interactive")
	err = os.MkdirAll(interactivePath, os.ModePerm)
	logmessage.CheckError(err, true)
	otherPath := filepath.Join(".", "other")
	err = os.MkdirAll(otherPath, os.ModePerm)
	logmessage.CheckError(err, true)
	pdfPath := filepath.Join(".", "pdf")
	err = os.MkdirAll(pdfPath, os.ModePerm)
	logmessage.CheckError(err, true)
	pngPath := filepath.Join(".", "png")
	err = os.MkdirAll(pngPath, os.ModePerm)
	logmessage.CheckError(err, true)
	svgPath := filepath.Join(".", "svg")
	err = os.MkdirAll(svgPath, os.ModePerm)
	logmessage.CheckError(err, true)

	// Initiate clustering method.
	if dataset.Params.Clustering == "hierarchical" {
		Hierarchical(dataset)
	}
	return
}
