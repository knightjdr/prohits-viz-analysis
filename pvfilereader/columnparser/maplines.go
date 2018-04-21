package columnparser

// Maplines maps slice elements to specified header columns
func Maplines(lines [][]string, headerMap map[string]int) []map[string]string {
	headerlen := len(headerMap)
	mappedLines := make([]map[string]string, len(lines)) // return slice map
	for i, line := range lines {
		lineMap := make(map[string]string, headerlen)
		for j, columnNum := range headerMap {
			lineMap[j] = line[columnNum]
		}
		mappedLines[i] = lineMap
	}
	return mappedLines
}
