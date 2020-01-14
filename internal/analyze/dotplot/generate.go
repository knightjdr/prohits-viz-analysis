// Package dotplot clusters conditions and readouts for visualization as a dotplot.
package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/analyze/dotplot/hierarchical"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

// Generate is the entry point for generating dotplots.
func Generate(analysis *types.Analysis) {
	createFolders(analysis.Settings.Png)

	switch analysis.Settings.Clustering {
	case "biclustering":
	case "none":
	default:
		hierarchical.Cluster(analysis)
	}

	// fs.Instance.Remove(filepath.Join(".", "minimap"))
}
