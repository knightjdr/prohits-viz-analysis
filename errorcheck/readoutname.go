package errorcheck

import "errors"

// ReadoutName ensures all readouts have names.
func ReadoutName(data []map[string]interface{}) (err error) {
	// iterate over data.
	for _, row := range data {
		readoutString := row["readout"].(string)
		if readoutString == "" {
			err = errors.New("All readouts should have a name")
			return
		}
	}
	return
}
