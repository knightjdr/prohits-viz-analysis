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

func TestWriteRatios(t *testing.T) {
	// Mock fs. Delay setting new
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create directory.
	fs.Instance.MkdirAll("other", 0755)
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

	// Data.
	baitList := []string{"bait1", "bait2", "bait3"}
	matrix := [][]float64{
		{0, 10, 74.2},
		{5, 7.2, 90.12},
		{8.3, 2, 1.4},
	}
	preyList := []string{"prey1", "prey2", "prey3"}

	// TEST1: bait prey transformed ratios.
	want := "\tbait1\tbait2\tbait3\nprey1\t0.00\t0.13\t1.00\nprey2\t0.06\t0.08\t1.00\nprey3\t1.00\t0.24\t0.17\n"
	WriteRatios(matrix, baitList, preyList)
	tsvFile, _ := afero.ReadFile(fs.Instance, "other/data-transformed-ratios.txt")
	assert.Equal(t, want, string(tsvFile), "Bait prey transformed ratios table not output correctly")
	fs.Instance.Remove("other/data-transformed-ratios.txt")

	// Mock Create. Method is unpatched using monkey.UnpatchAll() as
	// UnpatchInstanceMethod was not working between tests.
	dummyFile, _ := fs.Instance.Create("dummy.txt")
	fakeCreate := func(*afero.MemMapFs, string) (afero.File, error) {
		return dummyFile, errors.New("File cannot be created")
	}
	monkey.PatchInstanceMethod(reflect.TypeOf(fs.Instance), "Create", fakeCreate)

	// TEST2: write error.
	WriteRatios(matrix, baitList, preyList)
	// Ensure error is logged.
	logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
	want = "File cannot be created"
	matched, _ := regexp.MatchString(want, string(logfile))
	assert.True(t, matched, "File creation error message not being logged")
	// Ensure file has not been created.
	fileExists, _ := afero.Exists(fs.Instance, "other/data-transformed-ratios.txt")
	assert.False(
		t,
		fileExists,
		"Error during file creation should not generate file",
	)
	// Remove Create patch.
	monkey.UnpatchAll()
}
