package dotplot

import (
	"encoding/csv"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// NormalizeSlice normalizes the values in a slice to one. The greatest value
// in the slice will be one and the other values will be divided by this.
func NormalizeSlice(unnormalized []float64) (normalized []float64) {
	// Find max.
	max := float64(0)
	for _, value := range unnormalized {
		if value > max {
			max = value
		}
	}

	// Normalize input slice.
	normalized = make([]float64, len(unnormalized))
	for i, value := range unnormalized {
		normalized[i] = value / max
	}
	return
}

// WriteRatios writes a matrix to a tsv file with bait names as columns
// and prey names as rows. The matrix values are normalized by row to 1.
func WriteRatios(matrix [][]float64, baitList, preyList []string) {
	// Create file.
	dataTransformedFile, err := fs.Instance.Create("other/data-transformed-ratios.txt")
	// Log if error and return without panic.
	logmessage.CheckError(err, false)
	if err != nil {
		return
	}
	defer dataTransformedFile.Close()

	// Create csv writer.
	writer := csv.NewWriter(dataTransformedFile)
	writer.Comma = '\t'
	defer writer.Flush()

	// Create and write header.
	header := append([]string{""}, baitList...)
	err = writer.Write(header)
	// Log if error and return without panic.
	logmessage.CheckError(err, false)
	if err != nil {
		return
	}

	// Write each line.
	for i, row := range matrix {
		normRow := NormalizeSlice(row)
		line := append([]string{preyList[i]}, helper.ConvertFts(normRow, 2)...)
		err = writer.Write(line)

		// Log if error and return without panic.
		logmessage.CheckError(err, false)
		if err != nil {
			return
		}
	}
	return
}
