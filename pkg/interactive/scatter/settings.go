package scatter

import (
	"encoding/json"
	"fmt"
)

func parseSettings(settings map[string]interface{}) string {
	jsonString, _ := json.Marshal(settings)
	return fmt.Sprintf("\"settings\": %s", string(jsonString))
}
