package circheatmap

import (
	"regexp"
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestParseKnownReadouts(t *testing.T) {
	// Mock fs.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/testfile1.txt",
		[]byte(
			"a\tx\thuman\tmouse\n"+
				"a\ty\tdog\tmouse\n"+
				"z\ta\tmouse\thuman\n"+
				"b\tx\thuman\tmouse\n"+
				"c\tx\tdog\tmouse\n"+
				"y\tc\thuman\tmouse\n",
		),
		0444,
	)
	afero.WriteFile(fs.Instance, "test/unreadable.txt", []byte(""), 0444)
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

	mapping := map[string]string{
		"a": "conditionA",
		"c": "conditionC",
	}

	// TEST
	result := parseKnownReadouts(mapping, "test/testfile1.txt", "human")
	expected := map[string]map[string]bool{
		"conditionA": map[string]bool{"x": true, "z": true},
		"conditionC": map[string]bool{"y": true},
	}
	assert.Equal(t, expected, result, "Processed known file should get known readouts")

	// TEST: no file supplied
	noKnownResults := map[string]map[string]bool{
		"conditionA": make(map[string]bool),
		"conditionC": make(map[string]bool),
	}
	result = parseKnownReadouts(mapping, "", "human")
	assert.Equal(t, noKnownResults, result, "No known file should return no known readouts for each condition")

	// TEST: missing file logs message (intergration with logger) and returns empty map.
	result = parseKnownReadouts(mapping, "test/missing.txt", "human")
	assert.Equal(t, noKnownResults, result, "Missing file should return no known readouts for each condition")
	logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
	wantMessage := "file does not exist"
	matched, _ := regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "message not being logged")
}
