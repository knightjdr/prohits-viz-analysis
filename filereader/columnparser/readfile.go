// Package columnparser reads csv formatted files and returns specified columns
package columnparser

import (
	"errors"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// ReadFile will read a csv file(s) into a slice
func ReadFile(files []string, columnMap map[string]string, logFile string) ([]map[string]string, error) {
	// get mime type for each file
	filetype := make([]string, len(files))
	for i, filename := range files {
		var filetypeErr error
		filetype[i], filetypeErr = FileType(filename, logFile)
		if filetypeErr != nil {
			logmessage.Write(logFile, fmt.Sprintf("%s: could not be opened", filename))
		}
	}

	// read required header columns from files to slice map
	parsed := ParseCsv(files, filetype, columnMap, logFile)

	// if parsed slice is empty, return error
	var err error
	if len(parsed) == 0 {
		err = errors.New("No parsed results")
	}
	return parsed, err
}
