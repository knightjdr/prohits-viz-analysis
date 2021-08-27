package correlation

import (
	"regexp"

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
		KeepReps:        analysis.Settings.UseReplicates,
		ScoreType:       analysis.Settings.ScoreType,
		SeparateRepsBy:  "row",
	}
	matrices := convert.FromTable(&analysis.Data, matrixSettings)

	corrData := correlation.Data{
		Columns:                   matrices.Conditions,
		Dimension:                 "column",
		IgnoreSourceTargetMatches: analysis.Settings.IgnoreSourceTargetMatches,
		Matrix:                    matrices.Abundance,
		Method:                    analysis.Settings.Correlation,
		Rows:                      stripReplicates(matrices.Readouts, analysis.Settings.UseReplicates),
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
		KeepReps:        analysis.Settings.UseReplicates,
		ScoreType:       analysis.Settings.ScoreType,
		SeparateRepsBy:  "column",
	}
	matrices := convert.FromTable(&analysis.Data, matrixSettings)

	corrData := correlation.Data{
		Columns:                   stripReplicates(matrices.Conditions, analysis.Settings.UseReplicates),
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

func stripReplicates(names []string, useReplicates bool) []string {
	if useReplicates {
		re := regexp.MustCompile(`_R\d$`)

		stripped := make([]string, len(names))
		for i, name := range names {
			stripped[i] = re.ReplaceAllString(name, "")
		}
		return stripped
	}
	return names
}
