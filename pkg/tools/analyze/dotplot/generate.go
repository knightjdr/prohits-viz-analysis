// Package dotplot clusters conditions and readouts for visualization as a dotplot.
package dotplot

import (
	"path/filepath"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/dotplot/biclustering"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/dotplot/hierarchical"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/dotplot/nocluster"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Generate is the entry point for generating dotplots.
func Generate(analysis *types.Analysis) {
	createFolders(analysis.Settings)

	switch analysis.Settings.Clustering {
	case "biclustering":
		biclustering.Cluster(analysis)
	case "none":
		nocluster.Process(analysis)
	default:
		hierarchical.Cluster(analysis)
	}

	fs.Instance.RemoveAll(filepath.Join(".", "minimap"))
}
