package columnparser

import (
	"encoding/csv"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// acceptedTypes contains input file types with delimiter
var acceptedTypes = map[string]rune{
	"text/csv":                  ',',
	"text/plain":                '\t',
	"text/tab-separated-values": '\t',
}

// ParseCsv reads specified header columns from csv formatted files to slice
func ParseCsv(
	files []string,
	filetype []string,
	columnMap map[string]string,
	logFile string,
) (parsed []map[string]string) {
	for i, filename := range files {
		// only parse a file if it's an accepted type
		if delimiter, ok := acceptedTypes[filetype[i]]; ok {
			file, err := fs.Instance.Open(filename)
			if err != nil {
				// skip if file cannot be opened
				logmessage.Write(logFile, fmt.Sprintf("%s: could not be opened", filename))
				continue
			}

			// read file
			reader := csv.NewReader(file)
			reader.Comma = delimiter // set delimiter
			lines, err := reader.ReadAll()
			if err != nil {
				// skip if file cannot be read
				logmessage.Write(logFile, fmt.Sprintf("%s: could not be read", filename))
				continue
			}
			file.Close()

			// get header map
			headerMap, err := HeaderMap(columnMap, lines[0])
			if err != nil {
				// skip if columns missing from file
				logmessage.Write(logFile, fmt.Sprintf("%s: %s", filename, err.Error()))
				continue
			}

			// read file into data slice
			parsed = append(parsed, MapLines(lines[1:], headerMap)...)
		}
	}
	return
}
