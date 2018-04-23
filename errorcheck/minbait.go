package errorcheck

import (
	"errors"
	"fmt"
)

// MinBait ensures the data has the minimum number of required baits
func MinBait(data []map[string]interface{}, analysisType string) error {
	var err error
	// get minimum bait number for the analysis type
	var minBaits int
	switch analysisType {
	default:
		minBaits = 2
	}

	// find unique baits
	baits := make(map[string]bool, 0)
	for _, row := range data {
		baitString := row["bait"].(string)
		if _, ok := baits[baitString]; !ok {
			baits[baitString] = true
			if len(baits) >= minBaits {
				return err
			}
		}
	}

	// there must be less than the minimum baits to reach here, but check anyway
	if len(baits) < minBaits {
		err = errors.New(fmt.Sprintf("There are not enough baits for analysis. Min: %d", minBaits))
	}
	return err
}
