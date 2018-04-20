// Package columnparser reads csv formatted files and returns specified columns
package columnparser

import "errors"

// Readfile will read a csv file(s) into an array and filter based on criteria
func Readfile(files []string, columnMap map[string]string, logFile string) ([]map[string]string, error) {
	// get mime type for each file
	filetype := make([]string, len(files))
	for i, filename := range files {
		filetype[i], _ = Filetype(filename, logFile)
	}

	// read required header columns from files to array map
	parsed := Parsecsv(files, filetype, columnMap, logFile)

	// if parsed array is empty, return error
	var err error
	if len(parsed) == 0 {
		err = errors.New("No parsed results")
	}
	return parsed, err
}
