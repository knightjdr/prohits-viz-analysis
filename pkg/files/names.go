package files

import "path/filepath"

// ParseBaseNames converts a slice of file paths to a slice
// of file names.
func ParseBaseNames(files []string) []string {
	parsedFiles := make([]string, len(files))
	for i, file := range files {
		parsedFiles[i] = filepath.Base(file)
	}
	return parsedFiles
}
