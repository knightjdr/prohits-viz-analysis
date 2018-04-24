package errorcheck

import (
	"errors"
	"fmt"
)

// MinBait ensures the data has the minimum number of required baits.
func MinBait(data []map[string]interface{}, analysisType string) (err error) {
	// Get minimum bait number for the analysis type.
	var minBaits int
	switch analysisType {
	default:
		minBaits = 2
	}

	// Find unique baits.
	baits := make(map[string]bool, 0)
	for _, row := range data {
		baitName := row["bait"].(string)
		if _, ok := baits[baitName]; !ok {
			baits[baitName] = true
			if len(baits) >= minBaits {
				return err
			}
		}
	}

	// There must be less than the minimum baits to reach here, but check anyway.
	if len(baits) < minBaits {
		err = errors.New(fmt.Sprintf("There are not enough baits for analysis. Min: %d", minBaits))
	}
	return err
}
