package columnparser

import "errors"

// creates a map of specified headers to column numbers
func Headermap(columnMap map[string]string, header []string) (map[string]int, error) {
	columnsFound := 0                 // tracks header columns found
	headerMap := make(map[string]int) // return map
	// map columns to header
	for i, definedName := range columnMap {
		if definedName != "" { // ignore empty map values
			for j, columnName := range header {
				if definedName == columnName {
					columnsFound++
					headerMap[i] = j
					continue
				}
			}
		} else {
			columnsFound++ // empty map values get treated as found
		}
	}

	// check if any specified columns were not found
	var err error
	if columnsFound != len(columnMap) {
		err = errors.New("Missing header column")
	}
	return headerMap, err
}
