package circheatmap

import (
	"encoding/csv"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

func parseTissues(readouts map[string]bool, filename string, tissues []string) map[string]map[string]float64 {
	expression := make(map[string]map[string]float64, len(readouts))
	for readout := range readouts {
		expression[readout] = make(map[string]float64)
	}

	file, err := fs.Instance.Open(filename)
	// Skip if file cannot be opened.
	logmessage.CheckError(err, false)
	if err != nil {
		return expression
	}
	defer file.Close()

	// Read file.
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.LazyQuotes = true
	lines, err := reader.ReadAll()
	// Skip if file cannot be read
	logmessage.CheckError(err, false)
	if err != nil {
		return expression
	}

	// Create tissue hash
	tissueMap := make(map[string]bool, len(tissues))
	for _, tissue := range tissues {
		tissueMap[tissue] = true
	}

	for _, line := range lines {
		tissue := line[1]
		gene := line[0]
		expressionValue, _ := strconv.ParseFloat(line[2], 64)
		if readouts[gene] && tissueMap[tissue] {
			expression[gene][tissue] = expressionValue
		}
	}

	return expression
}
