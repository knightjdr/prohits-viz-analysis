package errorcheck

import (
	"errors"
	"fmt"
)

// MinCondition ensures the data has the minimum number of required conditions.
func MinCondition(data []map[string]interface{}, analysisType string) (err error) {
	// Get minimum condition number for the analysis type.
	var minConditions int
	switch analysisType {
	default:
		minConditions = 2
	}

	// Find unique conditions.
	conditions := make(map[string]bool, 0)
	for _, row := range data {
		conditionName := row["condition"].(string)
		if _, ok := conditions[conditionName]; !ok {
			conditions[conditionName] = true
			if len(conditions) >= minConditions {
				return err
			}
		}
	}

	// There must be less than the minimum conditions to reach here, but check anyway.
	if len(conditions) < minConditions {
		err = errors.New(fmt.Sprintf("There are not enough conditions for analysis. Min: %d", minConditions))
	}
	return err
}
