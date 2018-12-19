package file

import (
	"encoding/csv"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// WriteBBCytoscape writes a tabular file with condition-condition distance information for
// use with cytoscape.
func WriteBBCytoscape(matrix [][]float64, conditionList []string) {
	// Create file.
	dataTransformedFile, err := fs.Instance.Create("cytoscape/condition-condition-cytoscape.txt")
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
	// Log error and return but don't panic.
	logmessage.CheckError(err, false)
	if err != nil {
		return
	}

	// Write each row in the matrix.
	numConditions := len(matrix)
	for i, row := range matrix {
		for j := i + 1; j < numConditions; j++ {
			// Create row to write.
			rowSlice := make([]string, 3)
			rowSlice[0] = conditionList[i]
			rowSlice[1] = conditionList[j]
			rowSlice[2] = strconv.FormatFloat(row[j], 'f', -1, 64)
			err = dataTransformedWriter.Write(rowSlice)

			// Log error and return but don't panic.
			logmessage.CheckError(err, false)
			if err != nil {
				return
			}
		}
	}
	return
}
