package columnparser

import (
	"encoding/csv"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// acceptedTypes contains input file types with delimiter
var acceptedTypes = map[string]rune{
	"text/csv":                  ',',
	"text/plain":                '\t',
	"text/tab-separated-values": '\t',
}

// ParseCsv reads specified header columns from csv formatted files to slice.
func ParseCsv(
	files []string,
	filetype []string,
	columnMap map[string]string,
	ignoreMissing bool,
) (parsed []map[string]string) {
	for i, filename := range files {
		// Only parse a file if it's an accepted type.
		if delimiter, ok := acceptedTypes[filetype[i]]; ok {
			file, err := fs.Instance.Open(filename)
			// Skip if file cannot be opened.
			logmessage.CheckError(err, false)
			if err != nil {
				continue
			}

			// Read file.
			reader := csv.NewReader(file)
			reader.Comma = delimiter    // Set delimiter.
			reader.FieldsPerRecord = -1 // Negative so that rows with different number of fields are still parsed.
			reader.LazyQuotes = true
			lines, err := reader.ReadAll()
			// Skip if file cannot be read
			logmessage.CheckError(err, false)
			if err != nil {
				continue
			}
			file.Close()

			// Get header map.
			headerMap, err := HeaderMap(columnMap, lines[0], ignoreMissing)
			// Skip if columns missing from file.
			logmessage.CheckError(err, false)
			if err != nil {
				continue
			}

			// read file into data slice
			parsed = append(parsed, MapLines(lines[1:], headerMap)...)
		}
	}
	return
}
