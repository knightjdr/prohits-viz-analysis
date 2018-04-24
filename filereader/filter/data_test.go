package filter

import (
	"regexp"
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
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
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644)

	// TEST1: filter typical data slice.
	baits := make([]string, 0)
	data := []map[string]string{
		{"bait": "a", "prey": "b", "score": "0.5"},
		{"bait": "c", "prey": "d", "score": "0.1"},
		{"bait": "e", "prey": "f", "score": "0.8"},
	}
	preys := make([]string, 0)
	want := []map[string]interface{}{
		{"bait": "a", "prey": "b", "score": 0.5},
		{"bait": "c", "prey": "d", "score": 0.1},
	}
	filtered, err := Data(data, 0.5, baits, preys, "lte", "test/logfile.txt")
	assert.Nil(t, err, "Valid input should not produce an error")
	assert.Equal(t, want, filtered, "Data slice is not being filtered correctly")

	// TEST2: filter typical data slice by baits.
	baits = []string{"a", "c"}
	preys = make([]string, 0)
	want = []map[string]interface{}{
		{"bait": "a", "prey": "b", "score": 0.5},
		{"bait": "c", "prey": "d", "score": 0.1},
	}
	filtered, err = Data(data, 1, baits, preys, "lte", "test/logfile.txt")
	assert.Equal(t, want, filtered, "Data slice is not being filtered correctly by baits")

	// TEST3: filter typical data slice by preys.
	baits = make([]string, 0)
	preys = []string{"b", "f"}
	want = []map[string]interface{}{
		{"bait": "a", "prey": "b", "score": 0.5},
		{"bait": "e", "prey": "f", "score": 0.8},
	}
	filtered, err = Data(data, 1, baits, preys, "lte", "test/logfile.txt")
	assert.Equal(t, want, filtered, "Data slice is not being filtered correctly by preys")

	// TEST4: filter typical data slice by baits and preys.
	baits = []string{"a", "c"}
	preys = []string{"b", "f"}
	want = []map[string]interface{}{
		{"bait": "a", "prey": "b", "score": 0.5},
	}
	filtered, err = Data(data, 1, baits, preys, "lte", "test/logfile.txt")
	assert.Equal(t, want, filtered, "Data slice is not being filtered correctly by preys")

	// TEST5: no filtered results after bait and prey returns an error and logs it.
	baits = []string{"a", "c"}
	preys = []string{"f"}
	filtered, err = Data(data, 1, baits, preys, "lte", "test/logfile.txt")
	assert.NotNil(t, err, "No filtered results from bait/prey step should produce an error")
	logfile, _ := afero.ReadFile(fs.Instance, "test/logfile.txt")
	wantMessage := "No parsed results matching bait and prey criteria"
	matched, _ := regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged")
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644) // empty log file

	// TEST6: score step error returns an error.
	baits = make([]string, 0)
	data = []map[string]string{
		{"bait": "a", "prey": "b", "score": "x"},
	}
	preys = make([]string, 0)
	filtered, err = Data(data, 1, baits, preys, "lte", "test/logfile.txt")
	assert.NotNil(t, err, "Score step error should produce an error")

	// TEST7: no filtered results after score step returns an error and logs it.
	baits = make([]string, 0)
	data = []map[string]string{
		{"bait": "a", "prey": "b", "score": "0.5"},
		{"bait": "c", "prey": "d", "score": "0.1"},
		{"bait": "e", "prey": "f", "score": "0.8"},
	}
	preys = make([]string, 0)
	filtered, err = Data(data, 1, baits, preys, "gte", "test/logfile.txt")
	assert.NotNil(t, err, "No filtered results from score step should produce an error")
	logfile, _ = afero.ReadFile(fs.Instance, "test/logfile.txt")
	wantMessage = "No parsed results matching filter criteria"
	matched, _ = regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "Message not being logged")
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644) // Clear log file.
}
