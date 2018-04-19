package columnparser

// maps array elements to column headers
func Maplines(lines [][]string, headerMap map[string]int) []map[string]string {
	headerlen := len(headerMap)
	lineslen := len(lines)
	mappedLines := make([]map[string]string, lineslen) // return array map
	for i := 0; i < lineslen; i++ {
		lineMap := make(map[string]string, headerlen)
		for k, v := range headerMap {
			lineMap[k] = lines[i][v]
		}
		mappedLines[i] = lineMap
	}
	return mappedLines
}
