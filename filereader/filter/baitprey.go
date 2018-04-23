package filter

// BaitPrey filters slice map by baits and preys
func BaitPrey(
	data []map[string]string,
	baits []string,
	preys []string,
) []map[string]string {
	// convert bait slice to map
	baitMap := SliceToMap(baits)
	// convert prey slice to map
	preyMap := SliceToMap(preys)
	// iterate over slice and keep rows with bait and preys in maps
	datalen := len(data)
	filtered := data
	for i := datalen - 1; i >= 0; i-- {
		_, okBait := baitMap[data[i]["bait"]]
		_, okPrey := preyMap[data[i]["prey"]]
		if !okBait || !okPrey {
			filtered = append(filtered[:i], filtered[i+1:]...)
		}
	}
	return filtered
}
