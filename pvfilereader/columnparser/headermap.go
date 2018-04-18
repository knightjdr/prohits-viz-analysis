package columnparser

import "errors"

// creates a map of specified headers to column numbers
func Headermap(columns []string, header []string) (map[string]int, error) {
	columnsFound := 0 // tracks header columns found
	columnLen := len(columns)
	headerLen := len(header)
	headerMap := make(map[string]int) // return map
	// map columns to header
	for i := 0; i < columnLen; i++ {
		for j := 0; j < headerLen; j++ {
			if columns[i] == header[j] {
				columnsFound++
				headerMap[columns[i]] = j
				continue
			}
		}
	}
	// check if any specified columns were not found
	var err error
	if columnsFound != columnLen {
		err = errors.New("Missing header column")
	}
	return headerMap, err
}
