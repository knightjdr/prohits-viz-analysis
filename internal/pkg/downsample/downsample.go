package downsample

import (
	goMath "math"

	"github.com/knightjdr/prohits-viz-analysis/pkg/math"
)

type gridParameters struct {
	multipliers []float64
	scale       float64
	startIndex  int
	totalCells  int
}

func downsample(matrix [][]float64, scale float64) [][]float64 {
	downsampled := createDownsampledMatrix(matrix, scale)

	rowSubgrid := initializeSubgrid(scale)
	for rowIndex, row := range downsampled {
		columnSubgrid := initializeSubgrid(scale)
		for columnIndex := range row {
			downsampled[rowIndex][columnIndex] = averageSubgridValues(matrix, rowSubgrid, columnSubgrid)
			columnSubgrid = updateSubgrid(columnSubgrid, columnIndex+1)
		}
		rowSubgrid = updateSubgrid(rowSubgrid, rowIndex+1)
	}

	return downsampled
}

func createDownsampledMatrix(matrix [][]float64, scale float64) [][]float64 {
	columns := int(goMath.Round(float64(len(matrix[0])) / scale))
	rows := int(goMath.Round(float64(len(matrix)) / scale))

	downsampledMatrix := make([][]float64, rows)
	for i := range downsampledMatrix {
		downsampledMatrix[i] = make([]float64, columns)
	}
	return downsampledMatrix
}

func initializeSubgrid(scale float64) gridParameters {
	fullCells := int(goMath.Floor(scale))
	totalCells := int(goMath.Ceil(scale))

	multipliers := addFullCellMultipliers(fullCells)
	if totalCells > len(multipliers) {
		multipliers = append(multipliers, goMath.Mod(scale, 1.0))
	}

	return gridParameters{
		multipliers: multipliers,
		scale:       scale,
		startIndex:  0,
		totalCells:  totalCells,
	}
}

func updateSubgrid(previousGrid gridParameters, index int) gridParameters {
	scale := previousGrid.scale
	totalCells := previousGrid.totalCells

	lastMultiplierIndex := len(previousGrid.multipliers) - 1
	startMultiplier := 1.0 - previousGrid.multipliers[lastMultiplierIndex]

	fullCells := math.Floor(scale - startMultiplier)

	multipliers := []float64{startMultiplier}
	multipliers = append(multipliers, addFullCellMultipliers(fullCells)...)
	if totalCells > len(multipliers) {
		multipliers = append(multipliers, goMath.Mod(scale-float64(fullCells)-startMultiplier, 1.0))
	}

	return gridParameters{
		multipliers: multipliers,
		scale:       scale,
		startIndex:  math.Floor(scale * float64(index)),
		totalCells:  totalCells,
	}
}

func addFullCellMultipliers(length int) []float64 {
	multipliers := make([]float64, length)

	for i := 0; i < length; i++ {
		multipliers[i] = 1
	}

	return multipliers
}

func averageSubgridValues(matrix [][]float64, rowSubgrid, columnSubgrid gridParameters) float64 {
	values := make([]float64, 0)

	gridArea := float64(0)
	for i := 0; i < rowSubgrid.totalCells; i++ {
		rowIndex := rowSubgrid.startIndex + i
		rowMultiplier := rowSubgrid.multipliers[i]

		if len(matrix) > rowIndex {
			for j := 0; j < columnSubgrid.totalCells; j++ {
				columnIndex := columnSubgrid.startIndex + j
				columnMultiplier := columnSubgrid.multipliers[j]

				if len(matrix[0]) > columnIndex {
					cellArea := rowMultiplier * columnMultiplier
					gridArea += cellArea
					value := matrix[rowIndex][columnIndex] * cellArea
					values = append(values, value)
				}
			}
		}
	}

	return math.SumFloat(values) / gridArea
}
