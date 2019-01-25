package circheatmap

import (
	"encoding/csv"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

func parseMapFile(conditions []string, filename string) map[string]string {
	mapping := make(map[string]string, len(conditions))
	for _, condition := range conditions {
		mapping[condition] = condition
	}

	if filename == "" {
		return mapping
	}

	file, err := fs.Instance.Open(filename)
	// Skip if file cannot be opened.
	logmessage.CheckError(err, false)
	if err != nil {
		return mapping
	}
	defer file.Close()

	// Read file.
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true
	lines, err := reader.ReadAll()
	// Skip if file cannot be read
	logmessage.CheckError(err, false)
	if err != nil {
		return mapping
	}

	for _, line := range lines {
		if len(line) > 1 {
			mapping[line[1]] = line[0]
			delete(mapping, line[0])
		}
	}
	return mapping
}
