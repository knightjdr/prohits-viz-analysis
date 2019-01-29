package circheatmap

import (
	"testing"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/stretchr/testify/assert"
)

func TestAddKnownTest(t *testing.T) {
	// Mock parseKnownReadouts
	monkey.Patch(parseKnownReadouts, func(mapping map[string]string, filename string, species string) map[string]map[string]bool {
		return map[string]map[string]bool{
			"conditionA": map[string]bool{"x": true, "z": false},
			"conditionC": map[string]bool{"y": true},
		}
	})
	defer monkey.UnpatchAll()

	conditionMapping := map[string]string{
		"conditionA": "a",
		"conditionC": "c",
	}
	data := map[string]map[string]map[string]float64{
		"conditionA": map[string]map[string]float64{
			"x": map[string]float64{
				"abundance": 50,
			},
			"z": map[string]float64{
				"abundance": 10,
			},
		},
		"conditionC": map[string]map[string]float64{
			"y": map[string]float64{
				"abundance": 25,
			},
		},
	}

	// TEST: when known parameter is true, should return condition data where readouts have known key
	parameters := typedef.Parameters{
		Known:     true,
		KnownFile: "file",
		Species:   "Homo sapiens",
	}
	want := map[string]map[string]map[string]float64{
		"conditionA": map[string]map[string]float64{
			"x": map[string]float64{
				"abundance": 50,
				"known":     1,
			},
			"z": map[string]float64{
				"abundance": 10,
				"known":     0,
			},
		},
		"conditionC": map[string]map[string]float64{
			"y": map[string]float64{
				"abundance": 25,
				"known":     1,
			},
		},
	}
	assert.Equal(t, want, addKnown(data, conditionMapping, parameters), "Known property should be added to readouts")

	// TEST: when known parameter is false, should return input data
	parameters = typedef.Parameters{
		Known:     false,
		KnownFile: "file",
		Species:   "Homo sapiens",
	}
	assert.Equal(t, data, addKnown(data, conditionMapping, parameters), "Input data should be returned as is")
}
