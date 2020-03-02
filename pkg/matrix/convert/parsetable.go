package convert

import (
	"fmt"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/pkg/parse"
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

func parseTable(table *[]map[string]string, settings ConversionSettings) *tableData {
	data := &tableData{
		conditions:       make(map[string]int, 0),
		readouts:         make(map[string]int, 0),
		readoutCondition: make(map[readoutCondition]readoutData),
		worstScore:       float64(0),
	}

	parse := getRowParser(data, settings)

	for _, row := range *table {
		parse(row)
	}

	return data
}

func getRowParser(data *tableData, settings ConversionSettings) func(map[string]string) {
	findWorseScore := getScoreTest(settings.ScoreType)

	if settings.KeepReps {
		return func(row map[string]string) {
			abundances := parse.PipeSeparatedStringToArray(row["abundance"])
			condition := row["condition"]
			readout := row["readout"]
			score, _ := strconv.ParseFloat(row["score"], 64)

			data.worstScore = findWorseScore(score, data.worstScore)

			if _, ok := data.readouts[readout]; !ok {
				data.readouts[readout] = len(data.readouts)
			}

			noReps := len(abundances)
			for i := 0; i < noReps; i++ {
				conditionRep := fmt.Sprintf("%sR%d", condition, i+1)
				if _, ok := data.conditions[conditionRep]; !ok {
					data.conditions[conditionRep] = len(data.conditions)
				}

				data.readoutCondition[readoutCondition{readout, conditionRep}] = readoutData{abundance: abundances[i], score: score}
			}
		}
	}

	return func(row map[string]string) {
		abundance := parse.PipeSeparatedStringToMean(row["abundance"])
		condition := row["condition"]
		readout := row["readout"]
		score, _ := strconv.ParseFloat(row["score"], 64)

		data.worstScore = findWorseScore(score, data.worstScore)

		if _, ok := data.conditions[condition]; !ok {
			data.conditions[condition] = len(data.conditions)
		}

		if _, ok := data.readouts[readout]; !ok {
			data.readouts[readout] = len(data.readouts)
		}

		data.readoutCondition[readoutCondition{readout, condition}] = readoutData{abundance: abundance, score: score}
	}
}
