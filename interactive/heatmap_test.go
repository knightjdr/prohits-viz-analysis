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
	want := "{\n\t\"columns\": {\n" +
		"\t\t\"names\": [\n\t\t\t\"col1\",\n\t\t\t\"col2\",\n\t\t\t\"col3\"\n\t\t],\n\t\t\"ref\": null\n\t},\n" +
		"\t\"parameters\": {\n\t\t\"param1\": 2,\n\t\t\"param2\": \"a\",\n\t\t\"param3\": [\n\t\t\t\"a\",\n\t\t\t\"b\",\n\t\t\t\"c\"\n\t\t]\n\t},\n" +
		"\t\"settings\": {\n\t\t\"current\": {\n\t\t\t\"param1\": 1,\n\t\t\t\"param2\": \"b\",\n\t\t\t\"param3\": [\n\t\t\t\t\"x\",\n\t\t\t\t\"y\",\n\t\t\t\t\"z\"\n\t\t\t]\n\t\t}\n\t},\n" +
		"\t\"rows\": {\n\t\t\"list\": [\n\t\t\t{\n\t\t\t\t\"data\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"ratio\": 0.2,\n\t\t\t\t\t\t\"score\": 0.01,\n\t\t\t\t\t\t\"value\": 1\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"ratio\": 0.5,\n\t\t\t\t\t\t\"score\": 0.05,\n\t\t\t\t\t\t\"value\": 2\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"ratio\": 1,\n\t\t\t\t\t\t\"score\": 0.08,\n\t\t\t\t\t\t\"value\": 3\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"name\": \"row1\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"data\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"ratio\": 0.7,\n\t\t\t\t\t\t\"score\": 1,\n\t\t\t\t\t\t\"value\": 4\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"ratio\": 0.8,\n\t\t\t\t\t\t\"score\": 0.07,\n\t\t\t\t\t\t\"value\": 5\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"ratio\": 1,\n\t\t\t\t\t\t\"score\": 0.5,\n\t\t\t\t\t\t\"value\": 6\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"name\": \"row2\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"data\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"ratio\": 1,\n\t\t\t\t\t\t\"score\": 0.2,\n\t\t\t\t\t\t\"value\": 7\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"ratio\": 0.2,\n\t\t\t\t\t\t\"score\": 0.7,\n\t\t\t\t\t\t\"value\": 8\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"ratio\": 0.5,\n\t\t\t\t\t\t\"score\": 0.01,\n\t\t\t\t\t\t\"value\": 9\n\t\t\t\t\t}\n\t\t\t\t],\n" +
		"\t\t\t\t\"name\": \"row3\"\n\t\t\t}\n\t\t]\n\t},\n" +
		"\t\"minimap\": {\n\t\t\"image\": \"pngImage\"\n\t}\n}"
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
