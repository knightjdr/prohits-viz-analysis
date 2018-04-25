package main

import (
	"os"
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/types"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestParseFlagsCorrect(t *testing.T) {
	// Argument unmocking.
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Mock filesystem.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(fs.Instance, "test/logfile.txt", []byte(""), 0644)

	// TEST1: returns map and params with correct flags specied.
	os.Args = oldArgs
	os.Args = []string{
		"cmd",
		"-abundance", "abundanceColumn",
		"-analysisType", "dotplot",
		"-bait", "baitColumn",
		"-baitList", "bait1,bait2",
		"-control", "controlColumn",
		"-fileList", "testfile.txt",
		"-logBase", "2",
		"-logFile", "test/logfile.txt",
		"-normalization", "prey",
		"-normalizationPrey", "prey1",
		"-prey", "preyColumn",
		"-preyLength", "preyLengthColumn",
		"-preyList", "prey1,prey2",
		"-primaryFilter", "0.1",
		"-score", "scoreColumn",
		"-scoreType", "lte",
		"-secondaryFilter", "0.2",
	}
	wantMap := map[string]string{
		"abundance":  "abundanceColumn",
		"bait":       "baitColumn",
		"control":    "controlColumn",
		"prey":       "preyColumn",
		"preyLength": "preyLengthColumn",
		"score":      "scoreColumn",
	}
	wantParams := types.Parameters{
		Abundance:         "abundanceColumn",
		AnalysisType:      "dotplot",
		Bait:              "baitColumn",
		BaitList:          []string{"bait1", "bait2"},
		Control:           "controlColumn",
		Files:             []string{"testfile.txt"},
		LogBase:           "2",
		LogFile:           "test/logfile.txt",
		Normalization:     "prey",
		NormalizationPrey: "prey1",
		Prey:              "preyColumn",
		PreyLength:        "preyLengthColumn",
		PreyList:          []string{"prey1", "prey2"},
		PrimaryFilter:     0.1,
		Score:             "scoreColumn",
		ScoreType:         "lte",
		SecondaryFilter:   0.2,
	}
	columnMap, params, err := ParseFlags()
	assert.Nil(t, err, "All required arguments specified should not return an error")
	assert.Equal(t, wantMap, columnMap, "Column map is not correctly formatted")
	assert.Equal(t, wantParams, params, "Parameters are not correctly formatted")
}
