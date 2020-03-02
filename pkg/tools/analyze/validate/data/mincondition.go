package data

import (
	"fmt"
)

func confirmMinimumConditions(data []map[string]string, analysisType string) (err error) {
	minimumRequiredConditions := getMinimumRequiredConditions(analysisType)

	if countUniqueConditions(data, minimumRequiredConditions) < minimumRequiredConditions {
		err = fmt.Errorf("there are not enough conditions for analysis, min: %d", minimumRequiredConditions)
	}

	return err
}

func getMinimumRequiredConditions(analysisType string) int {
	var minConditions int

	switch analysisType {
	case "circheatmap":
		minConditions = 1
	default:
		minConditions = 2
	}

	return minConditions
}

func countUniqueConditions(data []map[string]string, minimumRequiredConditions int) int {
	conditions := make(map[string]bool, 0)

	for _, row := range data {
		condition := row["condition"]
		conditions[condition] = true
		if len(conditions) >= minimumRequiredConditions {
			return minimumRequiredConditions
		}
	}

	return len(conditions)
}
