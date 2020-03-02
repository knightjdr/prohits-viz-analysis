package parser

import (
	"encoding/csv"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
)

var acceptedTypes = map[string]rune{
	"text/csv":                  ',',
	"text/plain":                '\t',
	"text/tab-separated-values": '\t',
}

func parseCSV(files []string, mimeTypes []string, columnMap map[string]string, ignoreMissing bool) (parsedData []map[string]string) {
	for i, filename := range files {
		if delimiter, ok := acceptedTypes[mimeTypes[i]]; ok {
			file, err := fs.Instance.Open(filename)
			if shouldSkipFile(err) {
				continue
			}

			reader := csv.NewReader(createReader(file))
			reader.Comma = delimiter    // Set delimiter.
			reader.FieldsPerRecord = -1 // Negative so that rows with different number of fields are still parsed.
			reader.LazyQuotes = true
			lines, err := reader.ReadAll()
			if shouldSkipFile(err) {
				continue
			}
			file.Close()

			headerMap, err := mapHeader(columnMap, lines[0], ignoreMissing)
			if shouldSkipFile(err) {
				continue
			}

			// read file into data slice
			parsedData = append(parsedData, mapLines(lines[1:], headerMap)...)
		}
	}
	return
}

func shouldSkipFile(err error) bool {
	log.CheckError(err, false)
	if err != nil {

		return true
	}
	return false
}
