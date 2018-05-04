package dotplot

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// WritePPCytoscape writes a tabular file with prey-prey distance information for
// use with cytoscape.
func WritePPCytoscape(matrix [][]float64, preyList []string) {
	// Create file.
	dataTransformedFile, err := os.Create("cytoscape/prey-prey-cytoscape.txt")
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
	logmessage.CheckError(err, false)
	if err != nil {
		return
	}

	// Write each row in the matrix.
	numPreys := len(matrix)
	for i, row := range matrix {
		for j := i + 1; j < numPreys; j++ {
			// Create row to write.
			rowSlice := make([]string, 4)
			rowSlice[0] = preyList[i]
			rowSlice[1] = preyList[j]
			rowSlice[2] = strconv.FormatFloat(row[j], 'f', -1, 64)
			err = dataTransformedWriter.Write(rowSlice)
			logmessage.CheckError(err, false)
			if err != nil {
				return
			}
		}
	}
	return
}
