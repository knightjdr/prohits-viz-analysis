package interactive

import (
	"encoding/json"
	"errors"
	"regexp"
	"testing"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestHeatmap(t *testing.T) {
	// Mock filesystem.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Starting with some hypothetical data for a dotplot, generate a data matrix
	// with formatted rows for use in the heatmap.
	abundance := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	columns := []string{"col1", "col2", "col3"}
	parameters := map[string]interface{}{
		"param1": 2.0,
		"param2": "a",
		"param3": []string{"a", "b", "c"},
	}
	ratios := [][]float64{
		{0.2, 0.5, 1},
		{0.7, 0.8, 1},
		{1, 0.2, 0.5},
	}
	rows := []string{"row1", "row2", "row3"}
	scores := [][]float64{
		{0.01, 0.05, 0.08},
		{1, 0.07, 0.5},
		{0.2, 0.7, 0.01},
	}
	settings := map[string]interface{}{
		"param1": 1.0,
		"param2": "b",
		"param3": []string{"x", "y", "z"},
	}
	uri := "pngImage"
	numCols := len(columns)
	numRows := len(rows)
	data := make([]map[string]interface{}, numRows)
	for i, row := range abundance {
		rowslice := make([]map[string]float64, numCols)
		for j, value := range row {
			rowslice[j] = map[string]float64{
				"ratio": ratios[i][j],
				"score": scores[i][j],
				"value": value,
			}
		}
		data[i] = map[string]interface{}{
			"name": rows[i],
			"data": rowslice,
		}
	}

	// TEST1: generate JSON.
	want := "{\"columns\":[\"col1\",\"col2\",\"col3\"]," +
		"\"parameters\":{\"param1\":2,\"param2\":\"a\",\"param3\":[\"a\",\"b\",\"c\"]}," +
		"\"settings\":{\"param1\":1,\"param2\":\"b\",\"param3\":[\"x\",\"y\",\"z\"]}," +
		"\"rows\":[" +
		"{\"data\":[" +
		"{\"ratio\":0.2,\"score\":0.01,\"value\":1}," +
		"{\"ratio\":0.5,\"score\":0.05,\"value\":2}," +
		"{\"ratio\":1,\"score\":0.08,\"value\":3}]," +
		"\"name\":\"row1\"}," +
		"{\"data\":[" +
		"{\"ratio\":0.7,\"score\":1,\"value\":4}," +
		"{\"ratio\":0.8,\"score\":0.07,\"value\":5}," +
		"{\"ratio\":1,\"score\":0.5,\"value\":6}]," +
		"\"name\":\"row2\"}," +
		"{\"data\":[" +
		"{\"ratio\":1,\"score\":0.2,\"value\":7}," +
		"{\"ratio\":0.2,\"score\":0.7,\"value\":8}," +
		"{\"ratio\":0.5,\"score\":0.01,\"value\":9}]," +
		"\"name\":\"row3\"}]," +
		"\"minimap\":\"pngImage\"}"
	assert.Equal(
		t,
		want,
		Heatmap(data, columns, parameters, settings, uri),
		"Heatmap json is not correct",
	)

	// Mock error.
	fakeMarshall := func(interface{}) ([]uint8, error) {
		return []uint8{}, errors.New("Error creating json")
	}
	marshallPatch := monkey.Patch(json.Marshal, fakeMarshall)
	defer marshallPatch.Unpatch()

	// TEST2: error.
	Heatmap(data, columns, parameters, settings, uri)
	logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
	wantMessage := "Error creating json"
	matched, _ := regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "JSON error not being logged")
}
