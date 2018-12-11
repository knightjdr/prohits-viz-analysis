package errorcheck

import (
	"errors"
)

// MinData ensures there is data after parsing and filtering.
func MinData(data []map[string]string) (err error) {
	if len(data) <= 0 {
		err = errors.New("No data passes the required filters")
	}
	return
}
