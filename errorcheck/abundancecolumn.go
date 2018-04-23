package errorcheck

import (
	"errors"
	"strconv"
	"strings"
)

// AbundanceColumn ensures the abundance column is a pipe-separated list of numeric values
func AbundanceColumn(data []map[string]interface{}) error {
	var err error
	// split first abundance column
	column := strings.Split(data[0]["abundance"].(string), "|")
	for _, value := range column {
		_, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return errors.New("Abundance column is not a pipe-separated list of numbers")
		}
	}
	return err
}
