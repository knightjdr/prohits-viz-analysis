package data

import "errors"

func confirmParsedData(data []map[string]string) (err error) {
	if len(data) == 0 {
		err = errors.New("no parsed results satisfying filter criteria")
	}

	return
}
