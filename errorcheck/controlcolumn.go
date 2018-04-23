package errorcheck

import (
	"errors"
	"strconv"
	"strings"
)

// ControlColumn ensures the control column is a pipe-separated list of numeric values
func ControlColumn(data []map[string]interface{}, control string) error {
	var err error
	// if control column is null, we aren't using this column
	if control == "" {
		return err
	}

	// split first control column
	column := strings.Split(data[0]["control"].(string), "|")
	for _, value := range column {
		_, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return errors.New("Control column is not a pipe-separated list of numbers")
		}
	}
	return err
}
