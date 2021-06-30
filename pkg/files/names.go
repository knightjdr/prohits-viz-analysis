package files

import (
	"path/filepath"
	"strings"
)

// ParseBaseNameWithoutExtension extracts a file name without the extension or path.
func ParseBaseNameWithoutExtension(file string) string {
	basename := filepath.Base(file)
	parsedFile := strings.TrimSuffix(basename, filepath.Ext(basename))
	return parsedFile
}

// ParseBaseNames converts a slice of file paths to a slice
// of file names.
func ParseBaseNames(files []string) []string {
	parsedFiles := make([]string, len(files))
	for i, file := range files {
		parsedFiles[i] = filepath.Base(file)
	}
	return parsedFiles
}
