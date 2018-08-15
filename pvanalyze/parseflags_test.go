package main

import (
	"os"
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
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
	afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

	// TEST1: returns map and params with correct flags specied.
	os.Args = []string{
		"cmd",
		"-abundance", "abundanceColumn",
		"-analysisType", "dotplot",
		"-bait", "baitColumn",
		"-baitClustering", "baits",
		"-baitList", "bait1,bait2",
		"-biclusteringApprox=true",
		"-clustering", "hierarchical",
		"-clusteringMethod", "complete",
		"-fillColor", "blueBlack",
		"-control", "controlColumn",
		"-distance", "euclidean",
		"-fileList", "testfile.txt",
		"-invert=false",
		"-logBase", "2",
		"-maximumAbundance", "50",
		"-minimumAbundance", "0",
		"-normalization", "prey",
		"-normalizationPrey", "prey1",
		"-pdf=false",
		"-png=true",
		"-prey", "preyColumn",
		"-preyClustering", "preys",
		"-preyLength", "preyLengthColumn",
		"-preyList", "prey1,prey2",
		"-primaryFilter", "0.1",
		"-score", "scoreColumn",
		"-scoreType", "lte",
		"-secondaryFilter", "0.2",
		"-writeDistance=false",
		"-writeDotplot=false",
		"-writeHeatmap=false",
	}
	wantMap := map[string]string{
		"abundance":  "abundanceColumn",
		"bait":       "baitColumn",
		"control":    "controlColumn",
		"prey":       "preyColumn",
		"preyLength": "preyLengthColumn",
		"score":      "scoreColumn",
	}
	wantParams := typedef.Parameters{
		Abundance:          "abundanceColumn",
		AnalysisType:       "dotplot",
		Bait:               "baitColumn",
		BaitClustering:     "baits",
		BaitList:           []string{"bait1", "bait2"},
		BiclusteringApprox: true,
		Clustering:         "hierarchical",
		ClusteringMethod:   "complete",
		FillColor:          "blueBlack",
		Control:            "controlColumn",
		Distance:           "euclidean",
		Files:              []string{"testfile.txt"},
		Invert:             false,
		LogBase:            "2",
		MaximumAbundance:   50,
		MinimumAbundance:   0,
		Normalization:      "prey",
		NormalizationPrey:  "prey1",
		Pdf:                false,
		Png:                true,
		Prey:               "preyColumn",
		PreyClustering:     "preys",
		PreyLength:         "preyLengthColumn",
		PreyList:           []string{"prey1", "prey2"},
		PrimaryFilter:      0.1,
		Score:              "scoreColumn",
		ScoreType:          "lte",
		SecondaryFilter:    0.2,
		WriteDistance:      false,
		WriteDotplot:       false,
		WriteHeatmap:       false,
	}
	columnMap, params, err := ParseFlags()
	assert.Nil(t, err, "All required arguments specified should not return an error")
	assert.Equal(t, wantMap, columnMap, "Column map is not correctly formatted")
	assert.Equal(t, wantParams, params, "Parameters are not correctly formatted")
}
