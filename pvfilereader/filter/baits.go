package filter

// Filters slice map by baits
func Baits(
	data []map[string]string,
	baits []string,
) []map[string]string {
	// convert bait slice to map
	baitMap := Slicetomap(baits)
	// iterate over slice and keep rows with bait in preyMap
	datalen := len(data)
	filtered := data
	for i := datalen - 1; i >= 0; i-- {
		if _, ok := baitMap[data[i]["bait"]]; !ok {
			filtered = append(filtered[:i], filtered[i+1:]...)
		}
	}
	return filtered
}
