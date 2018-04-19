package columnparser

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/knightjdr/prohits-viz-analysis/pvfilereader/logmessage"
)

// accepted input file types
var acceptedTypes = map[string]rune{
	"text/csv":                  ',',
	"text/plain":                '\t',
	"text/tab-separated-values": '\t',
}

type Columns struct {
	abundance  string
	bait       string
	control    string
	prey       string
	preyLength int
	score      float64
}

// reads specified header columns from csv formatted files to struct
func Parsecsv(
	files []string,
	filetype []string,
	columnMap map[string]string,
	logFile string,
) []map[string]string {
	data := make([]map[string]string, 0) // return array map
	fileno := len(files)
	for i := 0; i < fileno; i++ {
		// only parse a file if it's an accepted type
		if delimiter, ok := acceptedTypes[filetype[i]]; ok {
			file, err := os.Open(files[i])
			if err != nil {
				// skip if file can't be opened
				logmessage.Write(logFile, fmt.Sprintf("%s: could not be opened", files[i]))
				continue
			}

			// read file
			reader := csv.NewReader(file)
			reader.Comma = delimiter // set delimiter
			lines, _ := reader.ReadAll()
			file.Close()

			// get header map
			headerMap, err := Headermap(columnMap, lines[0])
			if err != nil {
				// skip if columns missing from file
				logmessage.Write(logFile, fmt.Sprintf("%s: missing header columns", files[i]))
				continue
			}

			// read file into data array
			data = append(data, Maplines(lines[1:], headerMap)...)
		}
	}
	return data
}
