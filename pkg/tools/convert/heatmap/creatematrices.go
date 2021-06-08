package heatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/matrix/convert"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createMatrices(csv *[]map[string]string, scoreType string) *types.Matrices {
	conversionSettings := convert.ConversionSettings{
		CalculateRatios: true,
		Resort:          false,
		ScoreType:       scoreType,
	}

	return convert.FromTable(csv, conversionSettings)
}
