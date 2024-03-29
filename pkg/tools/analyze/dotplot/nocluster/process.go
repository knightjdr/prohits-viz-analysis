// Package nocluster generates dot plots based on requested condition and readout ordering.
package nocluster

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/matrix/convert"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/dotplot/hierarchical"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Process data.
func Process(analysis *types.Analysis) {
	matrixSettings := convert.ConversionSettings{
		CalculateRatios: true,
		RatioDimension:  analysis.Settings.RatioDimension,
		ScoreType:       analysis.Settings.ScoreType,
	}
	matrices := convert.FromTable(&analysis.Data, matrixSettings)

	analysis.Settings = hierarchical.AdjustSettings(analysis.Settings, matrices.Abundance)
	orderedData := order(matrices, analysis)
	sortedMatrices := sortMatrices(matrices, orderedData)

	hierarchical.CreateDotplot(sortedMatrices, orderedData, analysis.Settings)
	hierarchical.CreateHeatmap(sortedMatrices, orderedData, analysis.Settings)
	hierarchical.WriteDistance(sortedMatrices, orderedData, analysis.Settings)
	hierarchical.CreatePNGs(sortedMatrices, orderedData, analysis.Settings)
	hierarchical.CreateCytoscape(analysis.Data, sortedMatrices, analysis.Settings)
	hierarchical.WriteMatrices(sortedMatrices)
	hierarchical.WriteTrees(orderedData, analysis.Settings)
}
