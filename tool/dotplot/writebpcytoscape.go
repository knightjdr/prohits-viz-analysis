package dotplot

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/types"
)

// WriteBPCytoscape writes a tabular file with bait and prey information for
// use with cytoscape.
func WriteBPCytoscape(dataset types.Dataset) {
	// Create file.
	dataTransformedFile, err := os.Create("cytoscape/bait-prey-cytoscape.txt")
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
	header := []string{dataset.Params.Bait, dataset.Params.Prey, dataset.Params.Abundance, dataset.Params.Score}
	err = writer.Write(header)
	logmessage.CheckError(err, false)
	if err != nil {
		return
	}

	// Write each row in the dataset if the score passed the cutoff.
	cutoff := dataset.Params.PrimaryFilter
	filterFunc := helper.FilterFunc(dataset.Params.ScoreType)
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
			rowSlice[0] = row["bait"].(string)
			rowSlice[1] = row["prey"].(string)
			rowSlice[2] = strconv.FormatFloat(abundanceSum, 'f', 2, 64)
			rowSlice[3] = strconv.FormatFloat(row["score"].(float64), 'f', 2, 64)
			err = writer.Write(rowSlice)
			logmessage.CheckError(err, false)
			if err != nil {
				return
			}
		}
	}
	return
}