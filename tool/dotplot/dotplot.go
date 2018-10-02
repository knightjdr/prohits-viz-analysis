// Package dotplot clusters conditions and readouts for visualization as a dotplot
package dotplot

import (
	"os"
	"path/filepath"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Generate is the entry point for generating dotplots and output files.
func Generate(dataset typedef.Dataset) {
	// Create subfolders. Panic if error.

	// Cytoscape folder.
	cytoscapePath := filepath.Join(".", "cytoscape")
	err := fs.Instance.MkdirAll(cytoscapePath, os.ModePerm)
	logmessage.CheckError(err, true)

	// Interactive file folder.
	interactivePath := filepath.Join(".", "interactive")
	err = fs.Instance.MkdirAll(interactivePath, os.ModePerm)
	logmessage.CheckError(err, true)

	// Other files.
	otherPath := filepath.Join(".", "other")
	err = fs.Instance.MkdirAll(otherPath, os.ModePerm)
	logmessage.CheckError(err, true)

	// Minimap folder (can delete after making interactive files).
	mapPath := filepath.Join(".", "minimap")
	err = fs.Instance.MkdirAll(mapPath, os.ModePerm)
	logmessage.CheckError(err, true)

	// PDF folder.
	if dataset.Parameters.Pdf {
		pdfPath := filepath.Join(".", "pdf")
		err = fs.Instance.MkdirAll(pdfPath, os.ModePerm)
		logmessage.CheckError(err, true)
	}

	// PNG folder.
	if dataset.Parameters.Png {
		pngPath := filepath.Join(".", "png")
		err = fs.Instance.MkdirAll(pngPath, os.ModePerm)
		logmessage.CheckError(err, true)
	}

	// SVG folder.
	svgPath := filepath.Join(".", "svg")
	err = fs.Instance.MkdirAll(svgPath, os.ModePerm)
	logmessage.CheckError(err, true)

	// Initiate clustering method.
	if dataset.Parameters.Clustering == "biclustering" {
		Biclustering(dataset)
	} else if dataset.Parameters.Clustering == "hierarchical" {
		Hierarchical(dataset)
	} else if dataset.Parameters.Clustering == "none" {
		NoCluster(dataset)
	}

	// Remove minimap folder.
	fs.Instance.RemoveAll(mapPath)
	return
}
