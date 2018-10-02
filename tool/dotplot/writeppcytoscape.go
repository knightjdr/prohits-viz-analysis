package dotplot

import (
	"encoding/csv"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// WritePPCytoscape writes a tabular file with readout-readout distance information for
// use with cytoscape.
func WritePPCytoscape(matrix [][]float64, readoutList []string) {
	// Create file.
	dataTransformedFile, err := fs.Instance.Create("cytoscape/readout-readout-cytoscape.txt")
	// Log if error and return without panic.
	logmessage.CheckError(err, false)
	if err != nil {
		return
	}
	defer dataTransformedFile.Close()

	// Create csv writer.
	dataTransformedWriter := csv.NewWriter(dataTransformedFile)
	dataTransformedWriter.Comma = '\t'
	defer dataTransformedWriter.Flush()

	// Create and write header.
	header := []string{"source", "target", "distance"}
	err = dataTransformedWriter.Write(header)
	// Log if error and return without panic.
	logmessage.CheckError(err, false)
	if err != nil {
		return
	}

	// Write each row in the matrix.
	numReadouts := len(matrix)
	for i, row := range matrix {
		for j := i + 1; j < numReadouts; j++ {
			// Create row to write.
			rowSlice := make([]string, 3)
			rowSlice[0] = readoutList[i]
			rowSlice[1] = readoutList[j]
			rowSlice[2] = strconv.FormatFloat(row[j], 'f', -1, 64)
			err = dataTransformedWriter.Write(rowSlice)

			// Log if error and return without panic.
			logmessage.CheckError(err, false)
			if err != nil {
				return
			}
		}
	}
	return
}
