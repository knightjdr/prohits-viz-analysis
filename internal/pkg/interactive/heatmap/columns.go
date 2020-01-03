package heatmap

import (
	"encoding/json"
	"fmt"
)

func parseColumns(columns []string) string {
	jsonString, _ := json.Marshal(columns)
	return fmt.Sprintf("\"columnDB\": %s", string(jsonString))
}
