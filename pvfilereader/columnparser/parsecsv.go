package columnparser

import (
	"encoding/csv"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pvfilereader/fs"
	"github.com/knightjdr/prohits-viz-analysis/pvfilereader/logmessage"
)

// acceptedTypes contains input file types with delimiter
var acceptedTypes = map[string]rune{
	"text/csv":                  ',',
	"text/plain":                '\t',
	"text/tab-separated-values": '\t',
}

// Parsecsv reads specified header columns from csv formatted files to slice
func Parsecsv(
	files []string,
	filetype []string,
	columnMap map[string]string,
	logFile string,
) []map[string]string {
	data := make([]map[string]string, 0) // return array map
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
			headerMap, err := Headermap(columnMap, lines[0])
			if err != nil {
				// skip if columns missing from file
				logmessage.Write(logFile, fmt.Sprintf("%s: missing header columns", filename))
				continue
			}

			// read file into data slice
			data = append(data, Maplines(lines[1:], headerMap)...)
		}
	}
	return data
}
