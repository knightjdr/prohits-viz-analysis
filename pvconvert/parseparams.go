package main

import (
	"encoding/json"
	"math"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// jsonParams contains the specified file parameters.
type jsonParams struct {
	AbundanceColumn string  `json:"abundance,omitempty"`
	ConditionColumn string  `json:"xAxis,omitempty"`
	ImageType       string  `json:"type,omitempty"`
	InvertColor     int     `json:"invert,omitempty"`
	PrimaryFilter   float64 `json:"primary,omitempty"`
	ReadoutColumn   string  `json:"yAxis,omitempty"`
	SecondaryFilter float64 `json:"secondary,omitempty"`
	ScoreColumn     string  `json:"score"`
	ScoreType       int     `json:"filterType,omitempty"`
}

func invertColorToBool(invertColor int) bool {
	if invertColor == 1 {
		return true
	}
	return false
}

func scoreTypeToBool(scoreType int) string {
	if scoreType == 1 {
		return "gte"
	}
	return "lte"
}

// inferSettings infers missing settings from the image type and values. A dotplot image
// should have an edge color (default to blueBlack). If there are negative values the
// fill color should be a two color scale (default to redBlue) and the minAbundance should
// be floored to the nearest integer. For dotplots with only non-negative values, the
// abundanceCap is set to 50. In all other cases the cap is set to the ceiling of the max value.
func inferSettings(csv []map[string]string, imageType string) typedef.Parameters {
	settings := typedef.Parameters{}
	if imageType == "dotplot" {
		settings.EdgeColor = "blueBlack"
	}

	max := -math.MaxFloat64
	min := math.MaxFloat64
	for _, datum := range csv {
		value, _ := strconv.ParseFloat(datum["value"], 64)
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	// Set minimum and fill color.
	if min >= 0 {
		settings.FillColor = "blueBlack"
		settings.MinAbundance = float64(0)
	} else {
		settings.FillColor = "redBlue"
		settings.MinAbundance = float64(math.Floor(min))
	}

	// Set maximum
	if imageType == "dotplot" && min >= 0 {
		settings.AbundanceCap = float64(50)
	} else {
		settings.AbundanceCap = float64(math.Ceil(max))
	}
	return settings
}

// parseParams parses parameters from
func parseParams(csv []map[string]string) (imageType string, parameters typedef.Parameters) {
	// Read file parameters.
	inputParams := &jsonParams{}
	if string([]rune(csv[0]["params"])[0]) == "{" {
		err := json.Unmarshal([]byte(csv[0]["params"]), inputParams)
		logmessage.CheckError(err, false)
	} else {
		inputParams.ImageType = csv[0]["params"]
		inputParams.ScoreType, _ = strconv.Atoi(csv[1]["params"])
		inputParams.PrimaryFilter, _ = strconv.ParseFloat(csv[2]["params"], 64)
		inputParams.SecondaryFilter, _ = strconv.ParseFloat(csv[3]["params"], 64)
		inputParams.ScoreColumn = csv[4]["params"]
		inputParams.AbundanceColumn = csv[5]["params"]
		inputParams.InvertColor, _ = strconv.Atoi(csv[6]["params"])
	}
	imageType = inputParams.ImageType

	// Convert file parameters to output parameters.
	parameters.Abundance = inputParams.AbundanceColumn
	parameters.Condition = inputParams.ConditionColumn
	parameters.InvertColor = invertColorToBool(inputParams.InvertColor)
	parameters.Readout = inputParams.ReadoutColumn
	parameters.PrimaryFilter = inputParams.PrimaryFilter
	parameters.Score = inputParams.ScoreColumn
	parameters.ScoreType = scoreTypeToBool(inputParams.ScoreType)
	parameters.SecondaryFilter = inputParams.SecondaryFilter

	// Infer missing colors and abundance thresholds.
	inferred := inferSettings(csv, imageType)
	parameters.AbundanceCap = inferred.AbundanceCap
	parameters.EdgeColor = inferred.EdgeColor
	parameters.FillColor = inferred.FillColor
	parameters.MinAbundance = inferred.MinAbundance

	return
}
