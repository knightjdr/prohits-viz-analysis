package transform

import (
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

// adjustByReadoutLength adjusts readout abundance based on the length of a readout relative to
// all others in the data set. It takes the median length of all unique readouts, calculates the median
// and then divides the median by each readout's length to get the multiplier for that readout.
func adjustByReadoutLength(analysis *types.Analysis) {
	if analysis.Settings.ReadoutLength == "" {
		return
	}

	readoutLengths := getLengthOfUniqueReadouts(analysis.Data)
	medianLength := calculateMedian(readoutLengths)
	readoutMultipliers := calculateMultipliers(readoutLengths, medianLength)
	adjustAbundanceByMultiplier(analysis, "readout", readoutMultipliers)
}

func getLengthOfUniqueReadouts(data []map[string]string) map[string]float64 {
	readoutLengths := make(map[string]float64, 0)

	for _, row := range data {
		readout := row["readout"]
		length, err := strconv.ParseFloat(row["readoutLength"], 64)
		if _, ok := readoutLengths[readout]; !ok && err == nil {
			readoutLengths[readout] = length
		}
	}

	return readoutLengths
}
