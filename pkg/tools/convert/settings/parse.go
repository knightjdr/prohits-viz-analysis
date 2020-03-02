// Package settings parsed and/or infers settings for files to convert.
package settings

import (
	"encoding/json"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Parse "params" from file.
func Parse(csv []map[string]string) types.Settings {
	fileSettings := parseSettings(csv)
	conversionSettings := convert(fileSettings)
	addInferredSettings(csv, &conversionSettings)
	return conversionSettings
}

func parseSettings(csv []map[string]string) *jsonSettings {
	fileSettings := &jsonSettings{}
	if hasJSONSettings(csv) {
		parseJSON(fileSettings, csv)
	} else {
		parseText(fileSettings, csv)
	}

	return fileSettings
}

func hasJSONSettings(csv []map[string]string) bool {
	if string([]rune(csv[0]["params"])[0]) == "{" {
		return true
	}
	return false
}

func parseJSON(settings *jsonSettings, csv []map[string]string) {
	str := csv[0]["params"]
	err := json.Unmarshal([]byte(str), settings)
	log.CheckError(err, false)
}

func parseText(settings *jsonSettings, csv []map[string]string) {
	settings.Type = csv[0]["params"]
	settings.ScoreType, _ = strconv.Atoi(csv[1]["params"])
	settings.PrimaryFilter, _ = strconv.ParseFloat(csv[2]["params"], 64)
	settings.SecondaryFilter, _ = strconv.ParseFloat(csv[3]["params"], 64)
	settings.ScoreColumn = csv[4]["params"]
	settings.AbundanceColumn = csv[5]["params"]
	settings.InvertColor, _ = strconv.Atoi(csv[6]["params"])
}
