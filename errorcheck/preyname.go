package errorcheck

import "errors"

// PreyName ensures all preys have names.
func PreyName(data []map[string]interface{}) (err error) {
	// iterate over data.
	for _, row := range data {
		preyString := row["prey"].(string)
		if preyString == "" {
			err = errors.New("All preys should have a name")
			return
		}
	}
	return
}
