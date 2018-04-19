package columnparser

import "errors"

// creates a map of specified headers to column numbers
func Headermap(columnMap map[string]string, header []string) (map[string]int, error) {
	columnsFound := 0 // tracks header columns found
	headerLen := len(header)
	headerMap := make(map[string]int) // return map
	// map columns to header
	for k, v := range columnMap {
		if v != "" { // ignore empty map values
			for i := 0; i < headerLen; i++ {
				if v == header[i] {
					columnsFound++
					headerMap[k] = i
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
