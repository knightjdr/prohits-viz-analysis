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

func TestWriteMatrix(t *testing.T) {
	// Mock fs. Delay setting new
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create directory.
	fs.Instance.MkdirAll("other", 0755)
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

	// Data.
	conditionList := []string{"condition1", "condition2", "condition3"}
	filename := "other/test.txt"
	matrix := [][]float64{
		{0, 10, 74.2},
		{5, 7.2, 90.12},
		{8.3, 2, 1.4},
	}
	readoutList := []string{"readout1", "readout2", "readout3"}

	// TEST1: condition readout transformed data.
	want := "\tcondition1\tcondition2\tcondition3\nreadout1\t0.00\t10.00\t74.20\nreadout2\t5.00\t7.20\t90.12\nreadout3\t8.30\t2.00\t1.40\n"
	WriteMatrix(matrix, conditionList, readoutList, filename)
	tsvFile, _ := afero.ReadFile(fs.Instance, filename)
	assert.Equal(t, want, string(tsvFile), "Condition readout transformed data table not output correctly")
	fs.Instance.Remove(filename)

	// Mock Create. Method is unpatched using monkey.UnpatchAll() as
	// UnpatchInstanceMethod was not working between tests.
	dummyFile, _ := fs.Instance.Create("dummy.txt")
	fakeCreate := func(*afero.MemMapFs, string) (afero.File, error) {
		return dummyFile, errors.New("File cannot be created")
	}
	monkey.PatchInstanceMethod(reflect.TypeOf(fs.Instance), "Create", fakeCreate)

	// TEST2: write error.
	WriteMatrix(matrix, conditionList, readoutList, filename)
	// Ensure error is logged.
	logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
	want = "File cannot be created"
	matched, _ := regexp.MatchString(want, string(logfile))
	assert.True(t, matched, "File creation error message not being logged")
	// Ensure file has not been created.
	fileExists, _ := afero.Exists(fs.Instance, filename)
	assert.False(
		t,
		fileExists,
		"Error during file creation should not generate file",
	)
	// Remove Create patch.
	monkey.UnpatchAll()
}
