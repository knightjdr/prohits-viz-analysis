package nocluster

func checkRequestedList(fileData []map[string]string, column string, requestedList []string) []string {
	namesFromFile := make(map[string]bool)
	for _, datum := range fileData {
		name := datum[column]
		if !namesFromFile[name] {
			namesFromFile[name] = true
		}
	}

	keep := make([]string, 0)
	for _, item := range requestedList {
		if namesFromFile[item] {
			keep = append(keep, item)
		}
	}

	return keep
}
