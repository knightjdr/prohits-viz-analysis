package main

import (
	"encoding/json"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// Parameters contains the image parameters.
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

type parameters struct {
	abundanceColumn string
	conditionColumn string
	imageType       string
	invertColor     bool
	primaryFilter   float64
	readoutColumn   string
	secondaryFilter float64
	scoreColumn     string
	scoreType       string
}

func convertInvertColor(invertColor int) bool {
	if invertColor == 1 {
		return true
	}
	return false
}

func convertScoreType(scoreType int) string {
	if scoreType == 1 {
		return "gte"
	}
	return "lte"
}

// parseParams parses parameters from
func parseParams(csv []map[string]string) parameters {
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

	// Convert file parameters to output parameters.
	params := parameters{
		abundanceColumn: inputParams.AbundanceColumn,
		conditionColumn: inputParams.ConditionColumn,
		imageType:       inputParams.ImageType,
		invertColor:     convertInvertColor(inputParams.InvertColor),
		primaryFilter:   inputParams.PrimaryFilter,
		readoutColumn:   inputParams.ReadoutColumn,
		secondaryFilter: inputParams.SecondaryFilter,
		scoreColumn:     inputParams.ScoreColumn,
		scoreType:       convertScoreType(inputParams.ScoreType),
	}

	return params
}
