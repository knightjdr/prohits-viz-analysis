package filter

// ConditionReadout filters slice map by conditions and readouts.
func ConditionReadout(
	data []map[string]string,
	conditions []string,
	readouts []string,
) (filtered []map[string]string) {
	// Convert condition slice to map.
	conditionMap := SliceToMap(conditions)

	// Convert readout slice to map.
	readoutMap := SliceToMap(readouts)

	// Iterate over slice and keep rows with condition and readouts in maps.
	datalen := len(data)
	filtered = data
	for i := datalen - 1; i >= 0; i-- {
		_, okCondition := conditionMap[data[i]["condition"]]
		_, okReadout := readoutMap[data[i]["readout"]]
		if !okCondition || !okReadout {
			filtered = append(filtered[:i], filtered[i+1:]...)
		}
	}
	return
}
