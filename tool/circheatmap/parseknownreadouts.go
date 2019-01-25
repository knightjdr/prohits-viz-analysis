package circheatmap

import (
	"encoding/csv"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

func parseKnownReadouts(mapping map[string]string, filename string, species string) map[string]map[string]bool {
	known := make(map[string]map[string]bool, len(mapping))
	for _, condition := range mapping {
		known[condition] = make(map[string]bool)
	}

	file, err := fs.Instance.Open(filename)
	// Skip if file cannot be opened.
	logmessage.CheckError(err, false)
	if err != nil {
		return known
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
		return known
	}

	for _, line := range lines {
		source := line[0]
		sourceSpecies := line[2]
		target := line[1]
		targetSpecies := line[3]
		if sourceSpecies == species || targetSpecies == species {
			if mapping[source] != "" {
				condition := mapping[source]
				known[condition][target] = true
			} else if mapping[target] != "" {
				condition := mapping[target]
				known[condition][source] = true
			}
		}
	}

	return known
}
