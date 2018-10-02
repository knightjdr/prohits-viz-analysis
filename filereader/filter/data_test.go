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

	// TEST1: filter typical data slice when bait and prey lists not supplied.
	baits := make([]string, 0)
	data := []map[string]string{
		{"bait": "a", "prey": "b", "abundance": "5", "score": "0.5"},
		{"bait": "c", "prey": "d", "abundance": "10", "score": "0.1"},
		{"bait": "e", "prey": "f", "abundance": "2|1.2", "score": "0.8"},
	}
	preys := make([]string, 0)
	// Create dataset.
	parameters := typedef.Parameters{
		BaitClustering: "none",
		BaitList:       baits,
		MinAbundance:   0,
		PreyClustering: "none",
		PreyList:       preys,
		PrimaryFilter:  0.5,
		ScoreType:      "lte",
	}
	want := []map[string]interface{}{
		{"bait": "a", "prey": "b", "abundance": "5", "score": 0.5},
		{"bait": "c", "prey": "d", "abundance": "10", "score": 0.1},
	}
	filtered := Data(data, parameters)
	assert.Equal(t, want, filtered, "Data slice is not being filtered correctly")

	// TEST2: filter typical data slice by baits.
	baits = []string{"a", "c"}
	preys = make([]string, 0)
	parameters = typedef.Parameters{
		BaitClustering: "baits",
		BaitList:       baits,
		MinAbundance:   0,
		PreyClustering: "none",
		PreyList:       preys,
		PrimaryFilter:  1,
		ScoreType:      "lte",
	}
	want = []map[string]interface{}{
		{"bait": "a", "prey": "b", "abundance": "5", "score": 0.5},
		{"bait": "c", "prey": "d", "abundance": "10", "score": 0.1},
	}
	filtered = Data(data, parameters)
	assert.Equal(t, want, filtered, "Data slice is not being filtered correctly by baits")

	// TEST3: filter typical data slice by preys.
	baits = make([]string, 0)
	preys = []string{"b", "f"}
	parameters = typedef.Parameters{
		BaitClustering: "none",
		BaitList:       baits,
		MinAbundance:   0,
		PreyClustering: "preys",
		PreyList:       preys,
		PrimaryFilter:  1,
		ScoreType:      "lte",
	}
	want = []map[string]interface{}{
		{"bait": "a", "prey": "b", "abundance": "5", "score": 0.5},
		{"bait": "e", "prey": "f", "abundance": "2|1.2", "score": 0.8},
	}
	filtered = Data(data, parameters)
	assert.Equal(t, want, filtered, "Data slice is not being filtered correctly by preys")

	// TEST4: filter typical data slice by baits and preys.
	baits = []string{"a", "c"}
	preys = []string{"b", "f"}
	parameters = typedef.Parameters{
		BaitClustering: "baits",
		BaitList:       baits,
		MinAbundance:   0,
		PreyClustering: "preys",
		PreyList:       preys,
		PrimaryFilter:  1,
		ScoreType:      "lte",
	}
	want = []map[string]interface{}{
		{"bait": "a", "prey": "b", "abundance": "5", "score": 0.5},
	}
	filtered = Data(data, parameters)
	assert.Equal(t, want, filtered, "Data slice is not being filtered correctly by preys")

	// TEST5: no filtered results after bait and prey logs error and panics.
	baits = []string{"a", "c"}
	preys = []string{"f"}
	parameters = typedef.Parameters{
		BaitClustering: "baits",
		BaitList:       baits,
		MinAbundance:   0,
		PreyClustering: "preys",
		PreyList:       preys,
		PrimaryFilter:  1,
		ScoreType:      "lte",
	}
	assert.PanicsWithValue(
		t,
		"No parsed results matching bait and prey criteria",
		func() { Data(data, parameters) },
		"No parsed results should panic",
	)
	logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
	wantMessage := "No parsed results matching bait and prey criteria"
	matched, _ := regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged")
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644) // Clear log file.

	// TEST6: score step error returns an error.
	baits = make([]string, 0)
	data = []map[string]string{
		{"bait": "a", "prey": "b", "abundance": "5", "score": "x"},
	}
	preys = make([]string, 0)
	parameters = typedef.Parameters{
		BaitClustering: "baits",
		BaitList:       baits,
		MinAbundance:   0,
		PreyClustering: "preys",
		PreyList:       preys,
		PrimaryFilter:  1,
		ScoreType:      "lte",
	}
	assert.Panics(t, func() { Data(data, parameters) }, "Invalid score type should panic")

	// TEST7: no filtered results after score step returns an error and logs it.
	baits = make([]string, 0)
	data = []map[string]string{
		{"bait": "a", "prey": "b", "abundance": "5", "score": "0.5"},
		{"bait": "c", "prey": "d", "abundance": "10", "score": "0.1"},
		{"bait": "e", "prey": "f", "abundance": "2|1.2", "score": "0.8"},
	}
	preys = make([]string, 0)
	parameters = typedef.Parameters{
		BaitClustering: "none",
		BaitList:       baits,
		MinAbundance:   0,
		PreyClustering: "none",
		PreyList:       preys,
		PrimaryFilter:  1,
		ScoreType:      "gte",
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
