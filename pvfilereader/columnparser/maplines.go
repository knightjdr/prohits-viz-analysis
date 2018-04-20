package columnparser

// maps array elements to column headers
func Maplines(lines [][]string, headerMap map[string]int) []map[string]string {
	headerlen := len(headerMap)
	mappedLines := make([]map[string]string, len(lines)) // return array map
	for i, line := range lines {
		lineMap := make(map[string]string, headerlen)
		for j, columnNum := range headerMap {
			lineMap[j] = line[columnNum]
		}
		mappedLines[i] = lineMap
	}
	return mappedLines
}
