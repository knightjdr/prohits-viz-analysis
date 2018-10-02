package filter

// Readouts filters slice map by readouts.
func Readouts(
	data []map[string]string,
	readouts []string,
) (filtered []map[string]string) {
	// Convert readout slice to map.
	readoutMap := SliceToMap(readouts)

	// Iterate over slice and keep rows with readout in readoutMap.
	datalen := len(data)
	filtered = data
	for i := datalen - 1; i >= 0; i-- {
		if _, ok := readoutMap[data[i]["readout"]]; !ok {
			filtered = append(filtered[:i], filtered[i+1:]...)
		}
	}
	return
}
