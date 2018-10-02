package filter

// Conditions filters slice map by conditions.
func Conditions(
	data []map[string]string,
	conditions []string,
) (filtered []map[string]string) {
	// Convert condition slice to map.
	conditionMap := SliceToMap(conditions)
	// Iterate over slice and keep rows with condition in readoutMap.
	datalen := len(data)
	filtered = data
	for i := datalen - 1; i >= 0; i-- {
		if _, ok := conditionMap[data[i]["condition"]]; !ok {
			filtered = append(filtered[:i], filtered[i+1:]...)
		}
	}
	return
}
