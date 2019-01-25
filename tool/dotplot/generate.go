// Package dotplot clusters conditions and readouts for visualization as a dotplot
package dotplot

import (
	"path/filepath"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/tool/dotplot/biclustering"
	"github.com/knightjdr/prohits-viz-analysis/tool/dotplot/hierarchical"
	"github.com/knightjdr/prohits-viz-analysis/tool/dotplot/nocluster"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Generate is the entry point for generating dotplots and output files.
func Generate(dataset *typedef.Dataset) {
	// Determine folders to create.
	folders := make([]string, 0)
	folders = append(folders, []string{"cytoscape", "interactive", "minimap", "other", "svg"}...)
	if dataset.Parameters.Png {
		folders = append(folders, "png")
	}

	// Create subfolders. Will panic if error.
	helper.CreateFolders(folders)

	// Initiate clustering method.
	if dataset.Parameters.Clustering == "biclustering" {
		biclustering.Run(dataset)
	} else if dataset.Parameters.Clustering == "hierarchical" {
		hierarchical.Run(dataset)
	} else if dataset.Parameters.Clustering == "none" {
		nocluster.Run(dataset)
	}

	// Remove minimap folder.
	mapPath := filepath.Join(".", "minimap")
	fs.Instance.RemoveAll(mapPath)
	return
}
