package biclustering

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

// generateMatrix filters and formats a data matrix for the biclustering
// script. It will only keep readouts found in at least two conditions with the
// minimum criteria and will normalize all values by the max value * 10
// (not sure why this normalization is done - taking it frmo HW's script).
func generateMatrix(matrices *types.Matrices, minAbundance float64) []string {
	abundance := make([][]float64, 0)
	readouts := make([]string, 0)
	singles := make([]string, 0)
	for i, row := range matrices.Abundance {
		conditionCount := 0
		for _, value := range row {
			if value >= minAbundance {
				conditionCount++
				if conditionCount > 1 {
					abundance = append(abundance, row)
					readouts = append(readouts, matrices.Readouts[i])
					break
				}
			}
		}

		if conditionCount < 2 {
			singles = append(singles, matrices.Readouts[i])
		}
	}

	abundance = normalizeMatrix(abundance)

	var buffer strings.Builder
	header := append([]string{"PROT"}, matrices.Conditions...)
	buffer.WriteString(fmt.Sprintf("%s\n", strings.Join(header, "\t")))

	for i, row := range abundance {
		buffer.WriteString(readouts[i])
		for _, value := range row {
			buffer.WriteString(fmt.Sprintf("\t%0.5f", value))
		}
		buffer.WriteString("\n")
	}
	afero.WriteFile(fs.Instance, "biclustering/matrix.txt", []byte(buffer.String()), 0644)
	return singles
}

func normalizeMatrix(matrix [][]float64) [][]float64 {
	maxValue := float64(0)
	for _, row := range matrix {
		for _, value := range row {
			if value > maxValue {
				maxValue = value
			}
		}
	}

	normalized := make([][]float64, len(matrix))
	for i, row := range matrix {
		normalized[i] = make([]float64, len(row))
		for j, value := range row {
			normalized[i][j] = (value * 10) / maxValue
		}
	}

	return normalized
}
