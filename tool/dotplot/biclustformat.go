package dotplot

import (
	"encoding/csv"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// BiclustData holds information about the input data table.
type BiclustData struct {
	Abundance, Score              [][]float64
	Conditions, Readouts, Singles []string
}

// BiclustFormat filters and formats a data matrix for the biclustering
// script. It will only keep readouts found in at least two conditions with the
// minimum criteria and will normalize all values by the max value * 10
// (not sure why this normalization is done - taking it frmo HW's script).
// It returns a 2D matrix sorted by condition and readout alphabetically, as well
// as the list of readouts kept and readouts omitted (all sorted alphabetically).
func BiclustFormat(data typedef.Matrices, min float64) (filteredData BiclustData) {
	// Condition list doesn't change.
	filteredData.Conditions = data.Conditions

	// Iterate over matrix and find readouts that pass minimum abundance cutoff for
	// at least two conditions. Keep these readout rows in filtered matrix.
	keepReadouts := make([]int, 0)
	numCols := len(data.Abundance[0])
	singleReadouts := make([]int, 0)
	for i, row := range data.Abundance {
		conditionCount := 0
		newrow := make([]float64, numCols)
		copy(newrow, row)
		for _, value := range row {
			if value >= min {
				conditionCount++
				if conditionCount > 1 {
					keepReadouts = append(keepReadouts, i)
					filteredData.Abundance = append(filteredData.Abundance, newrow)
					filteredData.Score = append(filteredData.Score, data.Score[i])
					filteredData.Readouts = append(filteredData.Readouts, data.Readouts[i])
					break
				}
			}
		}

		// Add readout to singleton list if it doesn't meet requirement.
		if conditionCount < 2 {
			singleReadouts = append(singleReadouts, i)
			filteredData.Singles = append(filteredData.Singles, data.Readouts[i])
		}
	}

	// Find maximum abundance value in matrix.
	maxValue := float64(0)
	for _, row := range filteredData.Abundance {
		for _, value := range row {
			if value > maxValue {
				maxValue = value
			}
		}
	}

	// Normalize abundance matrix.
	for i, row := range filteredData.Abundance {
		for j, value := range row {
			filteredData.Abundance[i][j] = (value * 10) / maxValue
		}
	}

	// Output filtered file for biclustering.

	// Create file.
	dataFile, err := fs.Instance.Create("biclustering/matrix.txt")
	// Log if error and panic.
	logmessage.CheckError(err, true)
	defer dataFile.Close()

	// Create csv writer.
	writer := csv.NewWriter(dataFile)
	writer.Comma = '\t'
	defer writer.Flush()

	// Create and write header.
	header := append([]string{"PROT"}, filteredData.Conditions...)
	err = writer.Write(header)
	// Log if error and panic.
	logmessage.CheckError(err, true)

	// Write each line.
	for i, row := range filteredData.Abundance {
		line := append([]string{filteredData.Readouts[i]}, helper.ConvertFts(row, 5)...)
		err = writer.Write(line)

		// Log if error and return without panic.
		logmessage.CheckError(err, true)
	}
	return
}
