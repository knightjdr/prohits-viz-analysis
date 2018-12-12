// Package columnparser reads csv formatted files and returns specified columns
package columnparser

import (
	"errors"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// ReadFile will read a csv file(s) into a slice.
func ReadFile(files []string, columnMap map[string]string, ignoreMissing bool) (parsed []map[string]string) {
	// Get mime type for each file.
	filetype := make([]string, len(files))
	for i, filename := range files {
		var filetypeErr error
		filetype[i], filetypeErr = FileType(filename)
		// If a filetype cannot be opened, log it but don't panic. Will just skip.
		logmessage.CheckError(filetypeErr, false)
	}

	// Read required header columns from files to slice map.
	parsed = ParseCsv(files, filetype, columnMap, ignoreMissing)

	// If parsed slice is empty, log and panic.
	if len(parsed) == 0 {
		logmessage.CheckError(errors.New("No parsed results"), true)
	}
	return
}
