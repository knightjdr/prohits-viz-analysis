package errorcheck

import (
	"regexp"
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestRequired(t *testing.T) {
	// Mock filesystem.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

	// TEST1: filter typical data slice.
	data := []map[string]interface{}{
		{"abundance": "2", "condition": "a", "control": "1|5|2", "readout": "b", "readoutLength": "10", "score": 0.5},
		{"abundance": "2|3.1", "condition": "c", "control": "2|5.1|2", "readout": "d", "readoutLength": "1", "score": 0.1},
		{"abundance": "4", "condition": "e", "control": "1", "readout": "f", "readoutLength": "100", "score": 0.8},
	}
	dataset := typedef.Dataset{
		Data: data,
		Parameters: typedef.Parameters{
			AnalysisType:  "dotplot",
			Control:       "controlColumn",
			ReadoutLength: "readoutLengthColumn",
		},
	}

	err := Required(dataset)
	assert.Nil(t, err, "Valid input should not produce an error")

	// TEST2: no data panics and logs message.
	data = []map[string]interface{}{}
	dataset.Data = data
	assert.PanicsWithValue(
		t,
		"No data passes the required filters",
		func() { Required(dataset) },
		"No data should panic",
	)
	logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
	wantMessage := "No data passes the required filters"
	matched, _ := regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged when there is no data")
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644) // Clear log file.

	// TEST3: not enough conditions.
	data = []map[string]interface{}{
		{"abundance": "2", "condition": "a", "control": "1|5|2", "readout": "b", "readoutLength": "10", "score": 0.5},
	}
	dataset.Data = data
	err = Required(dataset)
	assert.NotNil(t, err, "Less than required number of conditions should produce error")
	logfile, _ = afero.ReadFile(fs.Instance, "error.txt")
	wantMessage = "There are not enough conditions for analysis. Min: 2"
	matched, _ = regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged when there are not enough conditions")
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644) // empty log file

	// TEST4: missing readout names.
	data = []map[string]interface{}{
		{"abundance": "2", "condition": "a", "control": "1|5|2", "readout": "b", "readoutLength": "10", "score": 0.5},
		{"abundance": "2|3.1", "condition": "c", "control": "2|5.1|2", "readout": "", "readoutLength": "1", "score": 0.1},
		{"abundance": "4", "condition": "e", "control": "1", "readout": "f", "readoutLength": "100", "score": 0.8},
	}
	dataset.Data = data
	err = Required(dataset)
	assert.NotNil(t, err, "Missing readout names should produce error")
	logfile, _ = afero.ReadFile(fs.Instance, "error.txt")
	wantMessage = "All readouts should have a name"
	matched, _ = regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged when there are missing readout names")
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644) // empty log file

	// TEST5: abundance column should be a pipe-separated list of numbers.
	data = []map[string]interface{}{
		{"abundance": "a", "condition": "a", "control": "5", "readout": "b", "readoutLength": "10", "score": 0.5},
		{"abundance": "2|3.1", "condition": "c", "control": "2|5.1|2", "readout": "d", "readoutLength": "1", "score": 0.1},
		{"abundance": "4", "condition": "e", "control": "1", "readout": "f", "readoutLength": "100", "score": 0.8},
	}
	dataset.Data = data
	err = Required(dataset)
	assert.NotNil(t, err, "Non pipe-separated abundance column should produce error")
	logfile, _ = afero.ReadFile(fs.Instance, "error.txt")
	wantMessage = "Abundance column is not a pipe-separated list of numbers"
	matched, _ = regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged when the abundance column is not valid")
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644) // empty log file

	// TEST6: score column should be a float64.
	data = []map[string]interface{}{
		{"abundance": "2", "condition": "a", "control": "1|5|2", "readout": "b", "readoutLength": "10", "score": "0.5"},
		{"abundance": "2|3.1", "condition": "c", "control": "2|5.1|2", "readout": "d", "readoutLength": "1", "score": 0.1},
		{"abundance": "4", "condition": "e", "control": "1", "readout": "f", "readoutLength": "100", "score": 0.8},
	}
	dataset.Data = data
	err = Required(dataset)
	assert.NotNil(t, err, "Incorrect score type should produce error")
	logfile, _ = afero.ReadFile(fs.Instance, "error.txt")
	wantMessage = "Score column is not numeric"
	matched, _ = regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged when the score column is not valid")
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644) // empty log file

	// TEST7: readout length column should be parsable as an integer.
	data = []map[string]interface{}{
		{"abundance": "2", "condition": "a", "control": "1|5|2", "readout": "b", "readoutLength": "a", "score": 0.5},
		{"abundance": "2|3.1", "condition": "c", "control": "2|5.1|2", "readout": "d", "readoutLength": "1", "score": 0.1},
		{"abundance": "4", "condition": "e", "control": "1", "readout": "f", "readoutLength": "100", "score": 0.8},
	}
	dataset.Data = data
	err = Required(dataset)
	assert.NotNil(t, err, "Non-integer parsable readout length should produce error")
	logfile, _ = afero.ReadFile(fs.Instance, "error.txt")
	wantMessage = "Readout length column must contain integer values"
	matched, _ = regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged when the readout length column is not valid")
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644) // empty log file

	// TEST8: control column should be a pipe-separated list of numbers.
	data = []map[string]interface{}{
		{"abundance": "2", "condition": "a", "control": "a", "readout": "b", "readoutLength": "10", "score": 0.5},
		{"abundance": "2|3.1", "condition": "c", "control": "2|5.1|2", "readout": "d", "readoutLength": "1", "score": 0.1},
		{"abundance": "4", "condition": "e", "control": "1", "readout": "f", "readoutLength": "100", "score": 0.8},
	}
	dataset.Data = data
	err = Required(dataset)
	assert.NotNil(t, err, "Non pipe-separated control column should produce error")
	logfile, _ = afero.ReadFile(fs.Instance, "error.txt")
	wantMessage = "Control column is not a pipe-separated list of numbers"
	matched, _ = regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged when the control column is not valid")
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644) // empty log file
}
