package filter

// Score filters a slice map based on the score key
func Score(
	data []map[string]string,
	primaryFilter float64,
	scoreType string,
) (filtered []map[string]interface{}, err error) {
	// ensure score column is numeric and convert
	filtered, err = NumericScore(data)
	if err != nil {
		return filtered, err
	}

	// get filter function
	filterFunc := FilterFunc(scoreType)

	// find unique preys passing score
	preys := make(map[string]bool, 0)
	for i, row := range filtered {
		passes := filterFunc(filtered[i]["score"].(float64), primaryFilter)
		if passes {
			preyString := row["prey"].(string)
			if _, ok := preys[preyString]; !ok { // only add preys not already present
				preys[preyString] = true
			}
		}
	}

	// remove preys not passing score
	filteredlen := len(filtered)
	for i := filteredlen - 1; i >= 0; i-- {
		preyString := filtered[i]["prey"].(string)
		if _, ok := preys[preyString]; !ok {
			filtered = append(filtered[:i], filtered[i+1:]...)
		}
	}

	return
}
