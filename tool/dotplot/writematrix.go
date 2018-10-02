package dotplot

import (
	"encoding/csv"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// WriteMatrix writes a matrix to a tsv file with condition names as columns
// and readout names as rows.
func WriteMatrix(matrix [][]float64, conditionList, readoutList []string, filename string) (err error) {
	// Create file.
	dataTransformedFile, err := fs.Instance.Create(filename)
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
	header := append([]string{""}, conditionList...)
	err = writer.Write(header)
	// Log if error and return without panic.
	logmessage.CheckError(err, false)
	if err != nil {
		return
	}

	// Write each line.
	for i, row := range matrix {
		line := append([]string{readoutList[i]}, helper.ConvertFts(row, 2)...)
		err = writer.Write(line)

		// Log if error and return without panic.
		logmessage.CheckError(err, false)
		if err != nil {
			return
		}
	}
	return
}
