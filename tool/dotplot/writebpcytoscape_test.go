package dotplot

import (
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestWriteBPCytoscape(t *testing.T) {
	// Mock fs. Delay setting new
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create directory.
	fs.Instance.MkdirAll("cytoscape", 0755)
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

	// Data.
	params := typedef.Parameters{
		Abundance:     "AvgSpec",
		Bait:          "Bait",
		Prey:          "PreyGene",
		PrimaryFilter: 0.01,
		Score:         "BFDR",
		ScoreType:     "lte",
	}
	dataset := typedef.Dataset{
		Data: []map[string]interface{}{
			{"bait": "bait1", "prey": "prey1", "abundance": "10", "score": 0.01},
			{"bait": "bait1", "prey": "prey2", "abundance": "5.5", "score": 0.02},
			{"bait": "bait2", "prey": "prey1", "abundance": "1", "score": float64(0)},
			{"bait": "bait2", "prey": "prey3", "abundance": "75", "score": 0.01},
		},
		Params: params,
	}

	// TEST1: typical dataset output.
	want := "Bait\tPreyGene\tAvgSpec\tBFDR\nbait1\tprey1\t10.00\t0.01\nbait2\tprey1\t1.00\t0.00\nbait2\tprey3\t75.00\t0.01\n"
	WriteBPCytoscape(dataset)
	tsvFile, _ := afero.ReadFile(fs.Instance, "cytoscape/bait-prey-cytoscape.txt")
	assert.Equal(t, want, string(tsvFile), "Bait-prey cytoscape file not output correctly")
	fs.Instance.Remove("cytoscape/bait-prey-cytoscape.txt")

	// Mock Create. Method is unpatched using monkey.UnpatchAll() as
	// UnpatchInstanceMethod was not working between tests.
	dummyFile, _ := fs.Instance.Create("dummy.txt")
	fakeCreate := func(*afero.MemMapFs, string) (afero.File, error) {
		return dummyFile, errors.New("File cannot be created")
	}
	monkey.PatchInstanceMethod(reflect.TypeOf(fs.Instance), "Create", fakeCreate)

	// TEST2: write error.
	WriteBPCytoscape(dataset)
	// Ensure error is logged.
	logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
	want = "File cannot be created"
	matched, _ := regexp.MatchString(want, string(logfile))
	assert.True(t, matched, "File creation error message not being logged")
	// Ensure file has not been created.
	fileExists, _ := afero.Exists(fs.Instance, "cytoscape/bait-prey-cytoscape.txt")
	assert.False(
		t,
		fileExists,
		"Error during file creation should not generate file",
	)
	// Remove Create patch.
	monkey.UnpatchAll()
}
