package parser

func mapLines(lines [][]string, headerMap map[string]int) []map[string]string {
	headerLength := len(headerMap)
	mappedLines := make([]map[string]string, len(lines))

	for i, line := range lines {
		lineMap := make(map[string]string, headerLength)
		for j, columnNum := range headerMap {
			lineMap[j] = line[columnNum]
		}
		mappedLines[i] = lineMap
	}

	return mappedLines
}
