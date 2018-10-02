package filter

import (
	"regexp"
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestData(t *testing.T) {
	// Mock filesystem.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

	// TEST1: filter typical data slice when condition and readout lists not supplied.
	conditions := make([]string, 0)
	data := []map[string]string{
		{"condition": "a", "readout": "b", "abundance": "5", "score": "0.5"},
		{"condition": "c", "readout": "d", "abundance": "10", "score": "0.1"},
		{"condition": "e", "readout": "f", "abundance": "2|1.2", "score": "0.8"},
	}
	readouts := make([]string, 0)
	// Create dataset.
	parameters := typedef.Parameters{
		ConditionClustering: "none",
		ConditionList:       conditions,
		MinAbundance:        0,
		ReadoutClustering:   "none",
		ReadoutList:         readouts,
		PrimaryFilter:       0.5,
		ScoreType:           "lte",
	}
	want := []map[string]interface{}{
		{"condition": "a", "readout": "b", "abundance": "5", "score": 0.5},
		{"condition": "c", "readout": "d", "abundance": "10", "score": 0.1},
	}
	filtered := Data(data, parameters)
	assert.Equal(t, want, filtered, "Data slice is not being filtered correctly")

	// TEST2: filter typical data slice by conditions.
	conditions = []string{"a", "c"}
	readouts = make([]string, 0)
	parameters = typedef.Parameters{
		ConditionClustering: "conditions",
		ConditionList:       conditions,
		MinAbundance:        0,
		ReadoutClustering:   "none",
		ReadoutList:         readouts,
		PrimaryFilter:       1,
		ScoreType:           "lte",
	}
	want = []map[string]interface{}{
		{"condition": "a", "readout": "b", "abundance": "5", "score": 0.5},
		{"condition": "c", "readout": "d", "abundance": "10", "score": 0.1},
	}
	filtered = Data(data, parameters)
	assert.Equal(t, want, filtered, "Data slice is not being filtered correctly by conditions")

	// TEST3: filter typical data slice by readouts.
	conditions = make([]string, 0)
	readouts = []string{"b", "f"}
	parameters = typedef.Parameters{
		ConditionClustering: "none",
		ConditionList:       conditions,
		MinAbundance:        0,
		ReadoutClustering:   "readouts",
		ReadoutList:         readouts,
		PrimaryFilter:       1,
		ScoreType:           "lte",
	}
	want = []map[string]interface{}{
		{"condition": "a", "readout": "b", "abundance": "5", "score": 0.5},
		{"condition": "e", "readout": "f", "abundance": "2|1.2", "score": 0.8},
	}
	filtered = Data(data, parameters)
	assert.Equal(t, want, filtered, "Data slice is not being filtered correctly by readouts")

	// TEST4: filter typical data slice by conditions and readouts.
	conditions = []string{"a", "c"}
	readouts = []string{"b", "f"}
	parameters = typedef.Parameters{
		ConditionClustering: "conditions",
		ConditionList:       conditions,
		MinAbundance:        0,
		ReadoutClustering:   "readouts",
		ReadoutList:         readouts,
		PrimaryFilter:       1,
		ScoreType:           "lte",
	}
	want = []map[string]interface{}{
		{"condition": "a", "readout": "b", "abundance": "5", "score": 0.5},
	}
	filtered = Data(data, parameters)
	assert.Equal(t, want, filtered, "Data slice is not being filtered correctly by readouts")

	// TEST5: no filtered results after condition and readout logs error and panics.
	conditions = []string{"a", "c"}
	readouts = []string{"f"}
	parameters = typedef.Parameters{
		ConditionClustering: "conditions",
		ConditionList:       conditions,
		MinAbundance:        0,
		ReadoutClustering:   "readouts",
		ReadoutList:         readouts,
		PrimaryFilter:       1,
		ScoreType:           "lte",
	}
	assert.PanicsWithValue(
		t,
		"No parsed results matching condition and readout criteria",
		func() { Data(data, parameters) },
		"No parsed results should panic",
	)
	logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
	wantMessage := "No parsed results matching condition and readout criteria"
	matched, _ := regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged")
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644) // Clear log file.

	// TEST6: score step error returns an error.
	conditions = make([]string, 0)
	data = []map[string]string{
		{"condition": "a", "readout": "b", "abundance": "5", "score": "x"},
	}
	readouts = make([]string, 0)
	parameters = typedef.Parameters{
		ConditionClustering: "conditions",
		ConditionList:       conditions,
		MinAbundance:        0,
		ReadoutClustering:   "readouts",
		ReadoutList:         readouts,
		PrimaryFilter:       1,
		ScoreType:           "lte",
	}
	assert.Panics(t, func() { Data(data, parameters) }, "Invalid score type should panic")

	// TEST7: no filtered results after score step returns an error and logs it.
	conditions = make([]string, 0)
	data = []map[string]string{
		{"condition": "a", "readout": "b", "abundance": "5", "score": "0.5"},
		{"condition": "c", "readout": "d", "abundance": "10", "score": "0.1"},
		{"condition": "e", "readout": "f", "abundance": "2|1.2", "score": "0.8"},
	}
	readouts = make([]string, 0)
	parameters = typedef.Parameters{
		ConditionClustering: "none",
		ConditionList:       conditions,
		MinAbundance:        0,
		ReadoutClustering:   "none",
		ReadoutList:         readouts,
		PrimaryFilter:       1,
		ScoreType:           "gte",
	}
	assert.PanicsWithValue(
		t,
		"No parsed results matching filter criteria",
		func() { Data(data, parameters) },
		"No parsed results matching filter criteria should panic",
	)
	logfile, _ = afero.ReadFile(fs.Instance, "error.txt")
	wantMessage = "No parsed results matching filter criteria"
	matched, _ = regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged")
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644) // Clear log file.
}
