// Package hierarchical clusters data for dot plots.
package hierarchical

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/matrix/convert"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

// Cluster data.
func Cluster(analysis *types.Analysis) {
	matrixSettings := convert.ConversionSettings{
		CalculateRatios: true,
		ScoreType:       analysis.Settings.ScoreType,
	}
	matrices := convert.FromTable(&analysis.Data, matrixSettings)

	clusteredData := cluster(matrices, analysis.Settings)
	sortedMatrices := sortMatrices(matrices, clusteredData)

	createDotplot(sortedMatrices, clusteredData, analysis.Settings)
	createHeatmap(sortedMatrices, clusteredData, analysis.Settings)
	writeDistance(sortedMatrices, clusteredData, analysis.Settings)
	createPNGs(sortedMatrices, clusteredData, analysis.Settings)
	createCytoscape(analysis.Data, sortedMatrices, analysis.Settings)
	writeMatrices(sortedMatrices)
	writeTrees(clusteredData, analysis.Settings)
}
