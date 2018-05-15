package dotplot

import (
	"errors"
	"reflect"
	"testing"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestBiclustFormat(t *testing.T) {
	// Mock fs. Delay setting new
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create directory.
	fs.Instance.MkdirAll("biclustering", 0755)
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

	// Data.
	dataset := Data{
		Abundance: [][]float64{
			{2, 4, 8},
			{4, 8, 16},
			{2, 8, 8},
		},
		Baits: []string{"abait", "bbait", "cbait"},
		Score: [][]float64{
			{0.01, 0, 2},
			{2, 1, 3},
			{2, 4, 5},
		},
		Preys: []string{"xprey", "yprey", "zprey"},
	}
	filename := "biclustering/matrix.txt"

	// TEST1: dataset converted and filtered to matrix with smaller scores better.
	wantAbundance := [][]float64{
		{2.5, 5, 10},
		{1.25, 5, 5},
	}
	wantBaitList := []string{"abait", "bbait", "cbait"}
	wantFile := "PROT\tabait\tbbait\tcbait\nyprey\t2.50000\t5.00000\t10.00000\nzprey\t1.25000\t5.00000\t5.00000\n"
	wantPreyList := []string{"yprey", "zprey"}
	wantSingleList := []string{"xprey"}
	wantScore := [][]float64{
		{2, 1, 3},
		{2, 4, 5},
	}
	data := BiclustFormat(dataset, 8)
	assert.Equal(t, wantAbundance, data.Abundance, "Data not converted to bait prey abundance matrix for biclustering")
	assert.Equal(t, wantBaitList, data.Baits, "Bait list not correct for biclustering")
	assert.Equal(t, wantPreyList, data.Preys, "Prey list not correct for biclustering")
	assert.Equal(t, wantScore, data.Score, "Data not converted to bait prey score matrix for biclustering")
	assert.Equal(t, wantSingleList, data.Singles, "Singleton prey list not correct for biclustering")
	tsvFile, _ := afero.ReadFile(fs.Instance, filename)
	assert.Equal(t, wantFile, string(tsvFile), "Bait prey transformed data table not output correctly")
	fs.Instance.Remove(filename)

	// Mock Create. Method is unpatched using monkey.UnpatchAll() as
	// UnpatchInstanceMethod was not working between tests.
	dummyFile, _ := fs.Instance.Create("dummy.txt")
	fakeCreate := func(*afero.MemMapFs, string) (afero.File, error) {
		return dummyFile, errors.New("File cannot be created")
	}
	monkey.PatchInstanceMethod(reflect.TypeOf(fs.Instance), "Create", fakeCreate)

	// TEST2: write error.
	wantMessage := "File cannot be created"
	assert.PanicsWithValue(t, wantMessage, func() { BiclustFormat(dataset, 8) })
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
