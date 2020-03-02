package correlation

import (
	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/pkg/correlation"
	"github.com/knightjdr/prohits-viz-analysis/pkg/data/filter"
	"github.com/knightjdr/prohits-viz-analysis/pkg/matrix/convert"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

type correlationData struct {
	dendrogram   hclust.Dendrogram
	labels       []string
	matrix       [][]float64
	sortedLabels []string
	tree         hclust.TreeLayout
}

func correlateConditions(analysis *types.Analysis) *correlationData {
	conditionAnalysis := &types.Analysis{
		Data: analysis.Data,
		Settings: types.Settings{
			MinAbundance:  analysis.Settings.ConditionAbundanceFilter,
			PrimaryFilter: analysis.Settings.ConditionScoreFilter,
			ScoreType:     analysis.Settings.ScoreType,
		},
	}
	filter.Process(conditionAnalysis)

	matrixSettings := convert.ConversionSettings{
		CalculateRatios: false,
		ScoreType:       analysis.Settings.ScoreType,
	}
	matrices := convert.FromTable(&analysis.Data, matrixSettings)

	corrData := correlation.Data{
		Columns:                   matrices.Conditions,
		Dimension:                 "column",
		IgnoreSourceTargetMatches: analysis.Settings.IgnoreSourceTargetMatches,
		Matrix:                    matrices.Abundance,
		Method:                    analysis.Settings.Correlation,
		Rows:                      matrices.Readouts,
	}
	correlationMatrix := corrData.Correlate()

	return &correlationData{
		labels: matrices.Conditions,
		matrix: correlationMatrix,
	}
}

func correlateReadouts(analysis *types.Analysis) *correlationData {
	readoutAnalysis := &types.Analysis{
		Data: analysis.Data,
		Settings: types.Settings{
			MinAbundance:  analysis.Settings.ReadoutAbundanceFilter,
			PrimaryFilter: analysis.Settings.ReadoutScoreFilter,
			ScoreType:     analysis.Settings.ScoreType,
		},
	}
	filter.Process(readoutAnalysis)

	matrixSettings := convert.ConversionSettings{
		CalculateRatios: false,
		KeepReps:        true,
		ScoreType:       analysis.Settings.ScoreType,
	}
	matrices := convert.FromTable(&analysis.Data, matrixSettings)

	corrData := correlation.Data{
		Columns:                   matrices.Readouts,
		Dimension:                 "row",
		IgnoreSourceTargetMatches: analysis.Settings.IgnoreSourceTargetMatches,
		Matrix:                    matrices.Abundance,
		Method:                    analysis.Settings.Correlation,
		Rows:                      matrices.Readouts,
	}
	correlationMatrix := corrData.Correlate()

	return &correlationData{
		labels: matrices.Readouts,
		matrix: correlationMatrix,
	}
}
