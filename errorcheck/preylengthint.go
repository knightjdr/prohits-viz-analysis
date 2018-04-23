package errorcheck

import (
	"errors"
	"strconv"
)

// PreyLengthInt ensures the prey length column is an integer
func PreyLengthInt(data []map[string]interface{}, preyLength string) error {
	var err error
	// if preyLength column is null, we aren't using this column
	if preyLength == "" {
		return err
	}

	// check if first row's prey length is an integer, if not return err
	_, err = strconv.ParseInt(data[0]["preyLength"].(string), 10, 64)
	if err != nil {
		return errors.New("Prey length column must contain integer values")
	}
	return err
}
