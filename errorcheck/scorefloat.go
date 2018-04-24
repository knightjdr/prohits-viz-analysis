package errorcheck

import (
	"errors"
	"reflect"
)

// ScoreFloat ensures the score column is of type float64.
func ScoreFloat(data []map[string]interface{}) (err error) {
	// Check if first row's score is numeric, if not return err.
	typeof := reflect.TypeOf(data[0]["score"])
	if typeof.String() != "float64" {
		err = errors.New("Score column is not numeric")
		return
	}
	return
}
