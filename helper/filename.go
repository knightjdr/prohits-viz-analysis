package helper

import "path/filepath"

// Filename converts a slice of file path to a slice
// of file names.
func Filename(files []string) []string {
	parsedFiles := make([]string, len(files))
	for i, file := range files {
		parsedFiles[i] = filepath.Base(file)
	}
	return parsedFiles
}
