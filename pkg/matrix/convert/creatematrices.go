package convert

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/normalize"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createMatrices(matrices *types.Matrices, data *tableData, settings ConversionSettings) {
	addAbundanceAndScoreMatrices(matrices, data)
	addRatioMatrix(matrices, settings)
}

func addAbundanceAndScoreMatrices(matrices *types.Matrices, data *tableData) {
	matrices.Abundance = make([][]float64, len(matrices.Readouts))
	matrices.Score = make([][]float64, len(matrices.Readouts))

	for i, readout := range matrices.Readouts {
		matrices.Abundance[i] = make([]float64, len(matrices.Conditions))
		matrices.Score[i] = make([]float64, len(matrices.Conditions))
		for j, condition := range matrices.Conditions {
			if value, ok := data.readoutCondition[readoutCondition{readout, condition}]; ok {
				matrices.Abundance[i][j] = value.abundance
				matrices.Score[i][j] = value.score
			} else {
				matrices.Abundance[i][j] = 0
				matrices.Score[i][j] = data.worstScore
			}
		}
	}
}

func addRatioMatrix(matrices *types.Matrices, settings ConversionSettings) {
	if settings.CalculateRatios {
		matrices.Ratio = normalize.Matrix(matrices.Abundance)
	}
}
