// Package biclustering clusters data using nestedcluster by H. Choi.
package biclustering

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/analyze/dotplot/hierarchical"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/matrix/convert"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

// Cluster data.
func Cluster(analysis *types.Analysis) {
	matrixSettings := convert.ConversionSettings{
		ScoreType: analysis.Settings.ScoreType,
	}
	matrices := convert.FromTable(&analysis.Data, matrixSettings)

	setParameters(len(matrices.Conditions), analysis.Settings)
	orderedData := order(matrices, analysis.Settings.MinAbundance)
	sortedMatrices := sortMatrices(matrices, orderedData)

	hierarchical.CreateDotplot(sortedMatrices, orderedData, analysis.Settings)
	hierarchical.CreateHeatmap(sortedMatrices, orderedData, analysis.Settings)
	hierarchical.WriteDistance(sortedMatrices, orderedData, analysis.Settings)
	hierarchical.CreatePNGs(sortedMatrices, orderedData, analysis.Settings)
	hierarchical.CreateCytoscape(analysis.Data, sortedMatrices, analysis.Settings)
	hierarchical.WriteMatrices(sortedMatrices)
	hierarchical.WriteTrees(orderedData, analysis.Settings)
}
