package filter

// Preys filters slice map by preys.
func Preys(
	data []map[string]string,
	preys []string,
) (filtered []map[string]string) {
	// Convert prey slice to map.
	preyMap := SliceToMap(preys)

	// Iterate over slice and keep rows with prey in preyMap.
	datalen := len(data)
	filtered = data
	for i := datalen - 1; i >= 0; i-- {
		if _, ok := preyMap[data[i]["prey"]]; !ok {
			filtered = append(filtered[:i], filtered[i+1:]...)
		}
	}
	return
}
