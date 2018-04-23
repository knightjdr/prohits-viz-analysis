package errorcheck

import (
	"errors"
	"reflect"
)

// ScoreFloat ensures the score column is of type float64
func ScoreFloat(data []map[string]interface{}) error {
	var err error
	// check if first row's score is numeric, if not return err
	typeof := reflect.TypeOf(data[0]["score"])
	if typeof.String() != "float64" {
		return errors.New("Score column is not numeric")
	}
	return err
}
