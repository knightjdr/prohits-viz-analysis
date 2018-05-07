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
	preyList := []string{"prey1", "prey2", "prey3"}
	matrix := [][]float64{
		{0, 10, 74.2},
		{10, 0, 90.12},
		{74.2, 90.12, 0},
	}

	// TEST1: prey table output.
	want := "source\ttarget\tdistance\nprey1\tprey2\t10\nprey1\tprey3\t74.2\nprey2\tprey3\t90.12\n"
	WritePPCytoscape(matrix, preyList)
	tsvFile, _ := afero.ReadFile(fs.Instance, "cytoscape/prey-prey-cytoscape.txt")
	assert.Equal(t, want, string(tsvFile), "Prey distance table not output correctly")
	fs.Instance.Remove("cytoscape/prey-prey-cytoscape.txt")

	// Mock Create.
	dummyFile, _ := fs.Instance.Create("dummy.txt")
	fakeCreate := func(*afero.MemMapFs, string) (afero.File, error) {
		return dummyFile, errors.New("File cannot be created")
	}
	monkey.PatchInstanceMethod(reflect.TypeOf(fs.Instance), "Create", fakeCreate)

	// TEST2: write error.
	WritePPCytoscape(matrix, preyList)
	// Ensure error is logged.
	logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
	want = "File cannot be created"
	matched, _ := regexp.MatchString(want, string(logfile))
	assert.True(t, matched, "File creation error message not being logged")
	// Ensure file has not been created.
	fileExists, _ := afero.Exists(fs.Instance, "cytoscape/prey-prey-cytoscape.txt")
	assert.False(
		t,
		fileExists,
		"Error during file creation should not generate file",
	)
	// Remove Create patch.
	monkey.UnpatchAll()
}
