package data

import (
	"fmt"
	"strconv"
)

func confirmReadLengthIsInt(data []map[string]string, readoutLength string) (err error) {
	if readoutLength == "" {
		return
	}

	for _, row := range data {
		value := row["readoutLength"]
		_, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("readout length column must contain integer values, offending value: %s", value)
		}
	}

	return
}
