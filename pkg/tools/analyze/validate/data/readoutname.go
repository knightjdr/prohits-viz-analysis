package data

import "errors"

func confirmReadoutsHaveNames(data []map[string]string) (err error) {
	for _, row := range data {
		readoutString := row["readout"]
		if readoutString == "" {
			err = errors.New("all readouts should have a name")
			return
		}
	}

	return
}
