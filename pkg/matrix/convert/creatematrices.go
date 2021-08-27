package convert

import (
	"math"

	matrixMath "github.com/knightjdr/prohits-viz-analysis/pkg/matrix/math"
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
		absoluteValueOfTheAbundance := matrixMath.AbsoluteValueEntries(matrices.Abundance)
		matrices.Ratio = normalize.Matrix(absoluteValueOfTheAbundance)
		if settings.RatioDimension == "area" {
			matrices.Ratio = adjustRatioForAreaCalculation(matrices.Ratio)
		}
	}
}

// The Ratio matrix should really be called the Radius matrix, since
// the value is used as a radius for drawing the dots, however the name
// "Ratio" needs to be kept for legacy reasons. If the ratio calculated
// from two abundances should be used for scaling the circle area, rather
// than the radius, the Ratio/Radius matrix needs to be adjusted.
// The new value is simply the sqrt(originalRatio / pi). However we want
// to scale everything to one, which can be achieved by simply calculating
// the sqrt(area).
func adjustRatioForAreaCalculation(matrix [][]float64) [][]float64 {
	adjusted := make([][]float64, len(matrix))

	for i, row := range matrix {
		adjusted[i] = make([]float64, len(row))
		for j, value := range row {
			adjusted[i][j] = math.Sqrt(value)
		}
	}
	return adjusted
}
