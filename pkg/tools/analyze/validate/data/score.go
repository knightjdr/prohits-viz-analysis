package data

import (
	"fmt"
	"strconv"
	"strings"
)

func confirmScoreIsFloat(data []map[string]string) (err error) {
	for _, row := range data {
		value := row["score"]
		_, err = strconv.ParseFloat(value, 64)
		if strings.EqualFold(value, "nan") || err != nil {
			return fmt.Errorf("score column must contain numeric values, offending value: %s", value)
		}
	}

	return
}
