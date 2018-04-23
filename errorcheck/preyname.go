package errorcheck

import "errors"

// PreyName ensures all preys have names
func PreyName(data []map[string]interface{}) error {
	var err error
	// iterate over data
	for _, row := range data {
		preyString := row["prey"].(string)
		if preyString == "" {
			return errors.New("All preys should have a name")
		}
	}
	return err
}
