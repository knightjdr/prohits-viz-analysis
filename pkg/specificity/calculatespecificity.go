// Package specificity calculates specificity scores.
package specificity

import (
	"math"
	"strconv"
	"strings"

	customMath "github.com/knightjdr/prohits-viz-analysis/pkg/math"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"gonum.org/v1/gonum/stat"
)

// Calculate scores for readouts between conditions.
func Calculate(analysis *types.Analysis) map[string]map[string]map[string]float64 {
	abundanceByReadout, noConditions := getAbundanceByReadout(analysis.Data)
	specificity := calculateSpecificityByMetric(abundanceByReadout, noConditions, analysis.Settings.SpecificityMetric)

	return reshapeSpecifictyByCondition(specificity)
}

func getAbundanceByReadout(data []map[string]string) (map[string]map[string]map[string]float64, int) {
	abundanceByReadout := make(map[string]map[string]map[string]float64)
	conditions := make(map[string]bool)
	for _, datum := range data {
		abundance := datum["abundance"]
		condition := datum["condition"]
		readout := datum["readout"]
		score, _ := strconv.ParseFloat(datum["score"], 64)

		conditions[condition] = true

		if _, ok := abundanceByReadout[readout]; !ok {
			abundanceByReadout[readout] = make(map[string]map[string]float64)
		}

		abundances := make([]float64, 0)
		reproducibility := 0.0
		for _, strValue := range strings.Split(abundance, "|") {
			value, _ := strconv.ParseFloat(strValue, 64)
			abundances = append(abundances, value)
			if value != 0 {
				reproducibility++
			}
		}

		abundanceByReadout[readout][condition] = map[string]float64{
			"abundance":       stat.Mean(abundances, nil),
			"reproducibility": reproducibility,
			"score":           score,
		}
	}

	return abundanceByReadout, len(conditions)
}

func calculateSpecificityByMetric(abundanceByReadout map[string]map[string]map[string]float64, noConditions int, metric string) map[string]map[string]map[string]float64 {
	defineSpecificity := getSpecificityMetric(metric, noConditions)
	specificity := make(map[string]map[string]map[string]float64, len(abundanceByReadout))

	for readout, conditionData := range abundanceByReadout {
		specificity[readout] = make(map[string]map[string]float64, len(conditionData))
		for condition := range conditionData {
			specificity[readout][condition] = defineSpecificity(condition, conditionData)
		}
	}
	return specificity
}

func getSpecificityMetric(metric string, noConditions int) func(condition string, abundanceByCondition map[string]map[string]float64) map[string]float64 {
	if metric == "zscore" {
		return func(condition string, abundanceByCondition map[string]map[string]float64) map[string]float64 {
			values := make([]float64, noConditions)
			i := 0
			for _, datum := range abundanceByCondition {
				values[i] = datum["abundance"]
				i++
			}
			mean, sd := stat.MeanStdDev(values, nil)
			specificity := 0.0
			if sd != 0 {
				specificity = customMath.Round((abundanceByCondition[condition]["abundance"]-mean)/sd, 0.01)
			}
			return map[string]float64{
				"abundance":   abundanceByCondition[condition]["abundance"],
				"score":       abundanceByCondition[condition]["score"],
				"specificity": specificity,
			}
		}
	}
	if metric == "sscore" {
		return func(condition string, abundanceByCondition map[string]map[string]float64) map[string]float64 {
			freq := float64(noConditions) / float64(len(abundanceByCondition))
			adjustedAbundance := freq * math.Abs(abundanceByCondition[condition]["abundance"])
			return map[string]float64{
				"abundance":   abundanceByCondition[condition]["abundance"],
				"score":       abundanceByCondition[condition]["score"],
				"specificity": customMath.Round(math.Sqrt(adjustedAbundance), 0.01),
			}
		}
	}
	if metric == "dscore" {
		return func(condition string, abundanceByCondition map[string]map[string]float64) map[string]float64 {
			freq := float64(noConditions) / float64(len(abundanceByCondition))
			multiplier := math.Pow(freq, abundanceByCondition[condition]["reproducibility"])
			adjustedAbundance := multiplier * math.Abs(abundanceByCondition[condition]["abundance"])
			return map[string]float64{
				"abundance":   abundanceByCondition[condition]["abundance"],
				"score":       abundanceByCondition[condition]["score"],
				"specificity": customMath.Round(math.Sqrt(adjustedAbundance), 0.01),
			}
		}
	}
	if metric == "wdscore" {
		return func(condition string, abundanceByCondition map[string]map[string]float64) map[string]float64 {
			values := make([]float64, noConditions)
			i := 0
			for _, datum := range abundanceByCondition {
				values[i] = math.Abs(datum["abundance"])
				i++
			}
			mean, sd := stat.MeanStdDev(values, nil)
			omega := sd / mean
			if omega < 1 {
				omega = 1
			}

			freq := float64(noConditions) / float64(len(abundanceByCondition))
			weightedFrequency := freq * omega
			multiplier := math.Pow(weightedFrequency, abundanceByCondition[condition]["reproducibility"])
			adjustedAbundance := multiplier * math.Abs(abundanceByCondition[condition]["abundance"])
			return map[string]float64{
				"abundance":   abundanceByCondition[condition]["abundance"],
				"score":       abundanceByCondition[condition]["score"],
				"specificity": customMath.Round(math.Sqrt(adjustedAbundance), 0.01),
			}
		}
	}
	return func(condition string, abundanceByCondition map[string]map[string]float64) map[string]float64 {
		values := make([]float64, noConditions-1)
		i := 0
		for key, datum := range abundanceByCondition {
			if key != condition {
				values[i] = math.Abs(datum["abundance"])
				i++
			}
		}
		mean := stat.Mean(values, nil)
		specificity := float64(0)
		if math.Abs(abundanceByCondition[condition]["abundance"]) > 0 {
			specificity = math.Abs(abundanceByCondition[condition]["abundance"] / mean)
		}
		return map[string]float64{
			"abundance":   abundanceByCondition[condition]["abundance"],
			"score":       abundanceByCondition[condition]["score"],
			"specificity": customMath.Round(specificity, 0.01),
		}
	}
}

func reshapeSpecifictyByCondition(specificity map[string]map[string]map[string]float64) map[string]map[string]map[string]float64 {
	specificityByCondition := make(map[string]map[string]map[string]float64)

	for readout, readoutData := range specificity {
		for condition, conditionData := range readoutData {
			if _, ok := specificityByCondition[condition]; !ok {
				specificityByCondition[condition] = make(map[string]map[string]float64)
			}
			specificityByCondition[condition][readout] = conditionData
		}
	}

	return specificityByCondition
}
