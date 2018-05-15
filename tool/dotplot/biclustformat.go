package dotplot

import (
	"encoding/csv"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// BiclustData holds information about the input data table.
type BiclustData struct {
	Abundance, Score      [][]float64
	Baits, Preys, Singles []string
}

// BiclustFormat filters and formats a data matrix for the biclustering
// script. It will only keep preys found in at least two baits with the
// minimum criteria and will normalize all values by the max value * 10
// (not sure why this normalization is done - taking it frmo HW's script).
// It returns a 2D matrix sorted by bait and prey alphabetically, as well
// as the list of preys kept and preys omitted (all sorted alphabetically).
func BiclustFormat(data Data, min float64) (filteredData BiclustData) {
	// Bait list doesn't change.
	filteredData.Baits = data.Baits

	// Iterate over matrix and find preys that pass minimum abundance cutoff for
	// at least two baits. Keep these prey rows in filtered matrix.
	keepPreys := make([]int, 0)
	numCols := len(data.Abundance[0])
	singlePreys := make([]int, 0)
	for i, row := range data.Abundance {
		baitCount := 0
		newrow := make([]float64, numCols)
		copy(newrow, row)
		for _, value := range row {
			if value >= min {
				baitCount++
				if baitCount > 1 {
					keepPreys = append(keepPreys, i)
					filteredData.Abundance = append(filteredData.Abundance, newrow)
					filteredData.Score = append(filteredData.Score, data.Score[i])
					filteredData.Preys = append(filteredData.Preys, data.Preys[i])
					break
				}
			}
		}

		// Add prey to singleton list if it doesn't meet requirement.
		if baitCount < 2 {
			singlePreys = append(singlePreys, i)
			filteredData.Singles = append(filteredData.Singles, data.Preys[i])
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
	header := append([]string{"PROT"}, filteredData.Baits...)
	err = writer.Write(header)
	// Log if error and panic.
	logmessage.CheckError(err, true)

	// Write each line.
	for i, row := range filteredData.Abundance {
		line := append([]string{filteredData.Preys[i]}, helper.ConvertFts(row, 5)...)
		err = writer.Write(line)

		// Log if error and return without panic.
		logmessage.CheckError(err, true)
	}
	return
}
