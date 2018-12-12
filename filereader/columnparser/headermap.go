package columnparser

import (
	"errors"
)

// HeaderMap creates a map of specified headers to column numbers.
func HeaderMap(columnMap map[string]string, header []string, ignoreMissing bool) (headerMap map[string]int, err error) {
	columnsFound := 0 // Tracks header columns found.
	headerMap = make(map[string]int)

	// Map columns to header.
	for i, definedName := range columnMap {
		if definedName != "" { // Ignore empty map values.
			for j, columnName := range header {
				if definedName == columnName {
					columnsFound++
					headerMap[i] = j
					continue
				}
			}
		} else {
			columnsFound++ // Empty map values get treated as found.
		}
	}

	// Check if any specified columns were not found.
	if !ignoreMissing && columnsFound != len(columnMap) {
		err = errors.New("Missing header column")
	}
	return headerMap, err
}
