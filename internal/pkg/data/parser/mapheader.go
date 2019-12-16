package parser

import (
	"errors"
)

func mapHeader(columnMap map[string]string, header []string, ignoreMissing bool) (headerMap map[string]int, err error) {
	columnsFound := 0
	headerMap = make(map[string]int, 0)

	for i, definedName := range columnMap {
		if definedName != "" {
			for j, columnName := range header {
				if definedName == columnName {
					columnsFound++
					headerMap[i] = j
					continue
				}
			}
		} else {
			columnsFound++
		}
	}

	if !ignoreMissing && columnsFound != len(columnMap) {
		err = errors.New("missing header column")
	}
	return headerMap, err
}
