package dotplot

import (
	"encoding/csv"
	"strconv"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// WriteBPCytoscape writes a tabular file with condition and readout information for
// use with cytoscape.
func WriteBPCytoscape(dataset typedef.Dataset) {
	// Create file.
	dataTransformedFile, err := fs.Instance.Create("cytoscape/condition-readout-cytoscape.txt")
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
	header := []string{dataset.Parameters.Condition, dataset.Parameters.Readout, dataset.Parameters.Abundance, dataset.Parameters.Score}
	err = writer.Write(header)
	// Log if error and return without panic.
	logmessage.CheckError(err, false)
	if err != nil {
		return
	}

	// Write each row in the dataset if the score passed the cutoff.
	cutoff := dataset.Parameters.PrimaryFilter
	filterFunc := helper.FilterFunc(dataset.Parameters.ScoreType)
	for _, row := range dataset.Data {
		if filterFunc(row["score"].(float64), cutoff) {
			// Abundance could be a pipe separated list. Split and sum to accomodate.
			abundance := strings.Split(row["abundance"].(string), "|")
			abundanceSum := float64(0)
			for _, value := range abundance {
				abdFloat, _ := strconv.ParseFloat(value, 64)
				abundanceSum += abdFloat
			}

			// Create row to write.
			rowSlice := make([]string, 4)
			rowSlice[0] = row["condition"].(string)
			rowSlice[1] = row["readout"].(string)
			rowSlice[2] = strconv.FormatFloat(abundanceSum, 'f', 2, 64)
			rowSlice[3] = strconv.FormatFloat(row["score"].(float64), 'f', 2, 64)
			err = writer.Write(rowSlice)

			// Log if error and return without panic.
			logmessage.CheckError(err, false)
			if err != nil {
				return
			}
		}
	}
	return
}
