package circheatmap

import (
	"testing"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestAddExpression(t *testing.T) {
	// Mock parseTissues
	monkey.Patch(parseTissues, func(readouts map[string]bool, filename string, tissues []string) map[string]map[string]float64 {
		return map[string]map[string]float64{
			"readout1": map[string]float64{"HEK 293": 10.4, "HeLa": 6.7},
			"readout3": map[string]float64{"HEK 293": 5, "HeLa": 8.1},
		}
	})
	defer monkey.UnpatchAll()

	data := map[string]map[string]map[string]float64{
		"conditionA": map[string]map[string]float64{
			"readout1": map[string]float64{
				"abundance": 50,
			},
			"readout3": map[string]float64{
				"abundance": 10,
			},
		},
		"conditionC": map[string]map[string]float64{
			"readout1": map[string]float64{
				"abundance": 25,
			},
		},
	}
	metrics := map[string]string{
		"abundance": "Abd",
	}
	readoutNames := map[string]bool{
		"readout1": true,
		"readout3": false,
	}

	// TEST: when known parameter is true, should return condition data where readouts have known key
	parameters := typedef.Parameters{
		Tissues:    []string{"HEK 293", "HeLa"},
		TissueFile: "file",
	}
	wantData := map[string]map[string]map[string]float64{
		"conditionA": map[string]map[string]float64{
			"readout1": map[string]float64{
				"abundance": 50,
				"HEK 293":   10.4,
				"HeLa":      6.7,
			},
			"readout3": map[string]float64{
				"abundance": 10,
				"HEK 293":   5,
				"HeLa":      8.1,
			},
		},
		"conditionC": map[string]map[string]float64{
			"readout1": map[string]float64{
				"abundance": 25,
				"HEK 293":   10.4,
				"HeLa":      6.7,
			},
		},
	}
	wantMetrics := map[string]string{
		"abundance": "Abd",
		"HEK 293":   "RNA expression HEK 293",
		"HeLa":      "RNA expression HeLa",
	}
	resultData, resultMetrics := addExpression(data, readoutNames, metrics, parameters)
	assert.Equal(t, wantData, resultData, "Expression data should be added to readouts")
	assert.Equal(t, wantMetrics, resultMetrics, "Tissue fields should be added to metrics")

	// TEST: when known parameter is false, should return input data
	parameters = typedef.Parameters{
		Tissues:    []string{},
		TissueFile: "file",
	}
	resultData, resultMetrics = addExpression(data, readoutNames, metrics, parameters)
	assert.Equal(t, data, resultData, "Expression data should not be added to readouts")
	assert.Equal(t, metrics, resultMetrics, "Tissue fields should not be added to metrics")
}
