package dotplot

import (
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestWritePPCytoscape(t *testing.T) {
	// Mock fs. Delay setting new
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create directory.
	fs.Instance.MkdirAll("cytoscape", 0755)
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

	// Data.
	readoutList := []string{"readout1", "readout2", "readout3"}
	matrix := [][]float64{
		{0, 10, 74.2},
		{10, 0, 90.12},
		{74.2, 90.12, 0},
	}

	// TEST1: readout table output.
	want := "source\ttarget\tdistance\nreadout1\treadout2\t10\nreadout1\treadout3\t74.2\nreadout2\treadout3\t90.12\n"
	WritePPCytoscape(matrix, readoutList)
	tsvFile, _ := afero.ReadFile(fs.Instance, "cytoscape/readout-readout-cytoscape.txt")
	assert.Equal(t, want, string(tsvFile), "Readout distance table not output correctly")
	fs.Instance.Remove("cytoscape/readout-readout-cytoscape.txt")

	// Mock Create.
	dummyFile, _ := fs.Instance.Create("dummy.txt")
	fakeCreate := func(*afero.MemMapFs, string) (afero.File, error) {
		return dummyFile, errors.New("File cannot be created")
	}
	monkey.PatchInstanceMethod(reflect.TypeOf(fs.Instance), "Create", fakeCreate)

	// TEST2: write error.
	WritePPCytoscape(matrix, readoutList)
	// Ensure error is logged.
	logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
	want = "File cannot be created"
	matched, _ := regexp.MatchString(want, string(logfile))
	assert.True(t, matched, "File creation error message not being logged")
	// Ensure file has not been created.
	fileExists, _ := afero.Exists(fs.Instance, "cytoscape/readout-readout-cytoscape.txt")
	assert.False(
		t,
		fileExists,
		"Error during file creation should not generate file",
	)
	// Remove Create patch.
	monkey.UnpatchAll()
}
