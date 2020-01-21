// Package correlation calculates the correlation between conditions and readouts.
package correlation

import (
	"path/filepath"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

// Generate is the entry point for correlation analysis.
func Generate(analysis *types.Analysis) {
	createFolders(analysis.Settings)

	corrConditions := correlateConditions(analysis)
	cluster(corrConditions, analysis.Settings)
	corrReadouts := correlateReadouts(analysis)
	cluster(corrReadouts, analysis.Settings)

	createCorrelationImages(corrConditions, corrReadouts, analysis.Settings)
	createCytoscapeFiles(corrConditions, corrReadouts, analysis.Settings)
	writeTrees(corrConditions, corrReadouts, analysis.Settings)
	createBaitPreyImages(analysis, corrConditions, corrReadouts)

	fs.Instance.Remove(filepath.Join(".", "minimap"))
}
