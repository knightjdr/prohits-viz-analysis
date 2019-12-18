package convert

import (
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/parse"
)

type readoutCondition struct {
	readout, condition string
}
type readoutData struct {
	abundance, score float64
}

type tableData struct {
	conditions       map[string]int
	readoutCondition map[readoutCondition]readoutData
	readouts         map[string]int
	worstScore       float64
}

func parseTable(table *[]map[string]string, scoreType string) *tableData {
	data := &tableData{
		conditions:       make(map[string]int, 0),
		readouts:         make(map[string]int, 0),
		readoutCondition: make(map[readoutCondition]readoutData),
		worstScore:       float64(0),
	}

	findWorseScore := getScoreTest(scoreType)
	for _, row := range *table {
		abundance := parse.PipeSeparatedStringToMean(row["abundance"])
		condition := row["condition"]
		readout := row["readout"]
		score, _ := strconv.ParseFloat(row["score"], 64)

		if _, ok := data.conditions[condition]; !ok {
			data.conditions[condition] = len(data.conditions)
		}

		if _, ok := data.readouts[readout]; !ok {
			data.readouts[readout] = len(data.readouts)
		}

		data.readoutCondition[readoutCondition{readout, condition}] = readoutData{abundance: abundance, score: score}
		data.worstScore = findWorseScore(score, data.worstScore)
	}

	return data
}
