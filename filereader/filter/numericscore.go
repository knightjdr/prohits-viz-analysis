package filter

import (
	"errors"
	"strconv"
)

// NumericScore checks if score column is numeric and convert its values from strings.
func NumericScore(data []map[string]string) (converted []map[string]interface{}, err error) {
	converted = make([]map[string]interface{}, len(data))
	// check if first row's score is numeric, if not return err
	_, err = strconv.ParseFloat(data[0]["score"], 64)
	if err != nil {
		err = errors.New("Score column is not numeric")
		return
	}

	// iterate data slice and convert scores to numeric
	for i, row := range data {
		newRow := make(map[string]interface{})
		// iterate over keys in map. Keep non-score keys as is; convert otherwise
		for key, value := range row {
			if key != "score" {
				newRow[key] = value
			} else {
				var scoreErr error
				newRow["score"], scoreErr = strconv.ParseFloat(row["score"], 64)
				if scoreErr != nil { // invalid score columns get set to zero
					newRow["score"] = 0
				}
			}
		}
		converted[i] = newRow
	}
	return
}
