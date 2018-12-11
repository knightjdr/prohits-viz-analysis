package errorcheck

import (
	"errors"
	"strconv"
)

// ScoreFloat ensures the score column is of type float64.
func ScoreFloat(data []map[string]string) (err error) {
	// Check if first row's score is numeric, if not return err.
	_, err = strconv.ParseFloat(data[0]["score"], 64)
	if err != nil {
		err = errors.New("Score column is not numeric")
		return
	}
	return
}
