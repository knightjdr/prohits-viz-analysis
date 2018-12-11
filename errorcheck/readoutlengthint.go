package errorcheck

import (
	"errors"
	"strconv"
)

// ReadoutLengthInt ensures the readout length column is an integer.
func ReadoutLengthInt(data []map[string]string, readoutLength string) (err error) {
	// If readoutLength column is null, we aren't using this column.
	if readoutLength == "" {
		return
	}

	// Check if first row's readout length is an integer, if not return err.
	_, err = strconv.ParseInt(data[0]["readoutLength"], 10, 64)
	if err != nil {
		return errors.New("Readout length column must contain integer values")
	}
	return
}
