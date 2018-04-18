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

// reads specified header columns from csv formatted files to struct
func Parsecsv(files []string, filetype []string, columns []string, logFile string) map[string][]map[string]string {
	data := make(map[string][]map[string]string) // return array map
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
			lines, err := reader.ReadAll()
			if err != nil {
				// skip if file can't be read
				logmessage.Write(logFile, fmt.Sprintf("%s: could not read file", files[i]))
				continue
			}
			file.Close()

			// get header map
			header, err := Headermap(columns, lines[0])
			if err != nil {
				// skip if columns missing from file
				logmessage.Write(logFile, fmt.Sprintf("%s: missing headers", files[i]))
				continue
			}

			// read file into data array
			fmt.Println(header)
		}
	}
	return data
}
