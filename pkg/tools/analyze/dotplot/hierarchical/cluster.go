// Package hierarchical clusters data for dot plots.
package hierarchical

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/matrix/convert"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
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

	CreateDotplot(sortedMatrices, clusteredData, analysis.Settings)
	CreateHeatmap(sortedMatrices, clusteredData, analysis.Settings)
	WriteDistance(sortedMatrices, clusteredData, analysis.Settings)
	CreatePNGs(sortedMatrices, clusteredData, analysis.Settings)
	CreateCytoscape(analysis.Data, sortedMatrices, analysis.Settings)
	WriteMatrices(sortedMatrices)
	WriteTrees(clusteredData, analysis.Settings)
}
