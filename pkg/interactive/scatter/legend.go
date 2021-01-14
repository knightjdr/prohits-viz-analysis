package scatter

import (
	"encoding/json"
	"fmt"
)

func parseLegend(legend []map[string]string) string {
	jsonString, _ := json.Marshal(legend)
	return fmt.Sprintf("\"legend\": %s", string(jsonString))
}
