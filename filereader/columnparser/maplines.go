package columnparser

// MapLines maps slice elements to specified header columns.
func MapLines(lines [][]string, headerMap map[string]int) (mappedLines []map[string]string) {
	headerlen := len(headerMap)
	mappedLines = make([]map[string]string, len(lines))
	for i, line := range lines {
		lineMap := make(map[string]string, headerlen)
		for j, columnNum := range headerMap {
			lineMap[j] = line[columnNum]
		}
		mappedLines[i] = lineMap
	}
	return mappedLines
}
