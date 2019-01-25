package circheatmap

import (
	"regexp"
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestParseTissues(t *testing.T) {
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
			"readout1\tHEK 293\t10.4\n"+
				"readout1\tHeLa\t6.7\n"+
				"readout1\tOther\t0.5\n"+
				"readout2\tHEK 293\t50\n"+
				"readout3\tHEK 293\t5\n"+
				"readout3\tHeLa\t8.1\n",
		),
		0444,
	)
	afero.WriteFile(fs.Instance, "test/unreadable.txt", []byte(""), 0444)
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

	readouts := map[string]bool{
		"readout1": true,
		"readout3": true,
	}

	// TEST
	result := parseTissues(readouts, "test/testfile1.txt", []string{"HEK 293", "HeLa"})
	expected := map[string]map[string]float64{
		"readout1": map[string]float64{"HEK 293": 10.4, "HeLa": 6.7},
		"readout3": map[string]float64{"HEK 293": 5, "HeLa": 8.1},
	}
	assert.Equal(t, expected, result, "Processed tissue file should get expression data for readouts")

	// TEST: no file supplied
	noExpressionResults := map[string]map[string]float64{
		"readout1": make(map[string]float64),
		"readout3": make(map[string]float64),
	}
	result = parseTissues(readouts, "", []string{})
	assert.Equal(t, noExpressionResults, result, "No tissue file should return no expression data for each readout")

	// TEST: missing file logs message (intergration with logger) and returns empty map.
	result = parseTissues(readouts, "test/missing.txt", []string{})
	assert.Equal(t, noExpressionResults, result, "Missing file should return no expression data for each readout")
	logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
	wantMessage := "file does not exist"
	matched, _ := regexp.MatchString(wantMessage, string(logfile))
	assert.True(t, matched, "message not being logged")
}
