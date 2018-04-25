package errorcheck

import (
	"regexp"
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/types"
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
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644)

	// TEST1: filter typical data slice.
	data := []map[string]interface{}{
		{"abundance": "2", "bait": "a", "control": "1|5|2", "prey": "b", "preyLength": "10", "score": 0.5},
		{"abundance": "2|3.1", "bait": "c", "control": "2|5.1|2", "prey": "d", "preyLength": "1", "score": 0.1},
		{"abundance": "4", "bait": "e", "control": "1", "prey": "f", "preyLength": "100", "score": 0.8},
	}
	dataset := types.Dataset{
		Data: data,
		Params: types.Parameters{
			AnalysisType: "dotplot",
			Control:      "controlColumn",
			LogFile:      "test/logfile.txt",
			PreyLength:   "preyLengthColumn",
		},
	}

	err := Required(dataset)
	assert.Nil(t, err, "Valid input should not produce an error")

	// TEST2: no data returns an error.
	data = []map[string]interface{}{}
	dataset.Data = data
	err = Required(dataset)
	assert.NotNil(t, err, "No data returns an error")
	logfile, _ := afero.ReadFile(fs.Instance, "test/logfile.txt")
	wantMessage := "No data passes the required filters"
	matched, _ := regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged when there is no data")
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644) // empty log file

	// TEST3: not enough baits.
	data = []map[string]interface{}{
		{"abundance": "2", "bait": "a", "control": "1|5|2", "prey": "b", "preyLength": "10", "score": 0.5},
	}
	dataset.Data = data
	err = Required(dataset)
	assert.NotNil(t, err, "Less than required number of baits should produce error")
	logfile, _ = afero.ReadFile(fs.Instance, "test/logfile.txt")
	wantMessage = "There are not enough baits for analysis. Min: 2"
	matched, _ = regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged when there are not enough baits")
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644) // empty log file

	// TEST4: missing prey names.
	data = []map[string]interface{}{
		{"abundance": "2", "bait": "a", "control": "1|5|2", "prey": "b", "preyLength": "10", "score": 0.5},
		{"abundance": "2|3.1", "bait": "c", "control": "2|5.1|2", "prey": "", "preyLength": "1", "score": 0.1},
		{"abundance": "4", "bait": "e", "control": "1", "prey": "f", "preyLength": "100", "score": 0.8},
	}
	dataset.Data = data
	err = Required(dataset)
	assert.NotNil(t, err, "Missing prey names should produce error")
	logfile, _ = afero.ReadFile(fs.Instance, "test/logfile.txt")
	wantMessage = "All preys should have a name"
	matched, _ = regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged when there are missing prey names")
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644) // empty log file

	// TEST5: abundance column should be a pipe-separated list of numbers.
	data = []map[string]interface{}{
		{"abundance": "a", "bait": "a", "control": "5", "prey": "b", "preyLength": "10", "score": 0.5},
		{"abundance": "2|3.1", "bait": "c", "control": "2|5.1|2", "prey": "d", "preyLength": "1", "score": 0.1},
		{"abundance": "4", "bait": "e", "control": "1", "prey": "f", "preyLength": "100", "score": 0.8},
	}
	dataset.Data = data
	err = Required(dataset)
	assert.NotNil(t, err, "Non pipe-separated abundance column should produce error")
	logfile, _ = afero.ReadFile(fs.Instance, "test/logfile.txt")
	wantMessage = "Abundance column is not a pipe-separated list of numbers"
	matched, _ = regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged when the abundance column is not valid")
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644) // empty log file

	// TEST6: score column should be a float64.
	data = []map[string]interface{}{
		{"abundance": "2", "bait": "a", "control": "1|5|2", "prey": "b", "preyLength": "10", "score": "0.5"},
		{"abundance": "2|3.1", "bait": "c", "control": "2|5.1|2", "prey": "d", "preyLength": "1", "score": 0.1},
		{"abundance": "4", "bait": "e", "control": "1", "prey": "f", "preyLength": "100", "score": 0.8},
	}
	dataset.Data = data
	err = Required(dataset)
	assert.NotNil(t, err, "Incorrect score type should produce error")
	logfile, _ = afero.ReadFile(fs.Instance, "test/logfile.txt")
	wantMessage = "Score column is not numeric"
	matched, _ = regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged when the score column is not valid")
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644) // empty log file

	// TEST7: prey length column should be parsable as an integer.
	data = []map[string]interface{}{
		{"abundance": "2", "bait": "a", "control": "1|5|2", "prey": "b", "preyLength": "a", "score": 0.5},
		{"abundance": "2|3.1", "bait": "c", "control": "2|5.1|2", "prey": "d", "preyLength": "1", "score": 0.1},
		{"abundance": "4", "bait": "e", "control": "1", "prey": "f", "preyLength": "100", "score": 0.8},
	}
	dataset.Data = data
	err = Required(dataset)
	assert.NotNil(t, err, "Non-integer parsable prey length should produce error")
	logfile, _ = afero.ReadFile(fs.Instance, "test/logfile.txt")
	wantMessage = "Prey length column must contain integer values"
	matched, _ = regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged when the prey length column is not valid")
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644) // empty log file

	// TEST8: control column should be a pipe-separated list of numbers.
	data = []map[string]interface{}{
		{"abundance": "2", "bait": "a", "control": "a", "prey": "b", "preyLength": "10", "score": 0.5},
		{"abundance": "2|3.1", "bait": "c", "control": "2|5.1|2", "prey": "d", "preyLength": "1", "score": 0.1},
		{"abundance": "4", "bait": "e", "control": "1", "prey": "f", "preyLength": "100", "score": 0.8},
	}
	dataset.Data = data
	err = Required(dataset)
	assert.NotNil(t, err, "Non pipe-separated control column should produce error")
	logfile, _ = afero.ReadFile(fs.Instance, "test/logfile.txt")
	wantMessage = "Control column is not a pipe-separated list of numbers"
	matched, _ = regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged when the control column is not valid")
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644) // empty log file
}
