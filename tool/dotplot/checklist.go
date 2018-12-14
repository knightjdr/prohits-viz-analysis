package dotplot

// checkList will return at list of column or row names. Only those names in both the
// input file and the input user list will be returned. The names will be ordered based
// on the input list.
func checkList(fileData []map[string]string, column string, inputList []string) []string {
	fileList := make(map[string]bool)
	for _, datum := range fileData {
		name := datum[column]
		if !fileList[name] {
			fileList[name] = true
		}
	}
	var keep []string
	for _, item := range inputList {
		if fileList[item] {
			keep = append(keep, item)
		}
	}
	return keep
}
