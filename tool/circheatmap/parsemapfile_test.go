package circheatmap

import (
	"regexp"
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestParseMapFile(t *testing.T) {
	// Mock fs.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/testfile1.txt",
		[]byte("condition1\trealname1\ncondition2\trealname2\ncondition3\n"),
		0444,
	)
	afero.WriteFile(fs.Instance, "test/unreadable.txt", []byte(""), 0444)
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

	conditions := []string{"condition1", "condition2", "condition4"}

	// TEST
	result := parseMapFile(conditions, "test/testfile1.txt")
	expected := map[string]string{
		"realname1":  "condition1",
		"realname2":  "condition2",
		"condition4": "condition4",
	}
	assert.Equal(t, expected, result, "Processed map file should create hash with condition name mapping")

	// TEST: no file supplied
	expected = map[string]string{
		"condition1": "condition1",
		"condition2": "condition2",
		"condition4": "condition4",
	}
	result = parseMapFile(conditions, "")
	assert.Equal(t, expected, result, "No map file should conditions mapped to themselves")

	// TEST: missing file logs message (intergration with logger) and returns empty map.
	result = parseMapFile(conditions, "test/missing.txt")
	assert.Equal(t, expected, result, "Missing file should conditions mapped to themselves")
	logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
	wantMessage := "file does not exist"
	matched, _ := regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "message not being logged")
}
