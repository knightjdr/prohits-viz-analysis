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

	// TEST1: returns map and parameters with correct flags specied.
	os.Args = []string{
		"cmd",
		"-abundance", "abundanceColumn",
		"-abundanceCap", "50",
		"-analysisType", "dotplot",
		"-condition", "conditionColumn",
		"-conditionClustering", "conditions",
		"-conditionList", "condition1,condition2",
		"-biclusteringApprox=true",
		"-clustering", "hierarchical",
		"-clusteringMethod", "complete",
		"-fillColor", "blueBlack",
		"-control", "controlColumn",
		"-distance", "euclidean",
		"-fileList", "testfile.txt",
		"-invertColor=false",
		"-logBase", "2",
		"-minAbundance", "0",
		"-normalization", "readout",
		"-normalizationReadout", "readout1",
		"-pdf=false",
		"-png=true",
		"-readout", "readoutColumn",
		"-readoutClustering", "readouts",
		"-readoutLength", "readoutLengthColumn",
		"-readoutList", "readout1,readout2",
		"-primaryFilter", "0.1",
		"-score", "scoreColumn",
		"-scoreType", "lte",
		"-secondaryFilter", "0.2",
		"-writeDistance=false",
		"-writeDotplot=false",
		"-writeHeatmap=false",
	}
	wantMap := map[string]string{
		"abundance":     "abundanceColumn",
		"condition":     "conditionColumn",
		"control":       "controlColumn",
		"readout":       "readoutColumn",
		"readoutLength": "readoutLengthColumn",
		"score":         "scoreColumn",
	}
	wantParams := typedef.Parameters{
		Abundance:            "abundanceColumn",
		AbundanceCap:         50,
		AnalysisType:         "dotplot",
		Condition:            "conditionColumn",
		ConditionClustering:  "conditions",
		ConditionList:        []string{"condition1", "condition2"},
		BiclusteringApprox:   true,
		Clustering:           "hierarchical",
		ClusteringMethod:     "complete",
		FillColor:            "blueBlack",
		Control:              "controlColumn",
		Distance:             "euclidean",
		Files:                []string{"testfile.txt"},
		InvertColor:          false,
		LogBase:              "2",
		MinAbundance:         0,
		Normalization:        "readout",
		NormalizationReadout: "readout1",
		Pdf:                  false,
		Png:                  true,
		Readout:              "readoutColumn",
		ReadoutClustering:    "readouts",
		ReadoutLength:        "readoutLengthColumn",
		ReadoutList:          []string{"readout1", "readout2"},
		PrimaryFilter:        0.1,
		Score:                "scoreColumn",
		ScoreType:            "lte",
		SecondaryFilter:      0.2,
		WriteDistance:        false,
		WriteDotplot:         false,
		WriteHeatmap:         false,
	}
	columnMap, parameters, err := ParseFlags()
	assert.Nil(t, err, "All required arguments specified should not return an error")
	assert.Equal(t, wantMap, columnMap, "Column map is not correctly formatted")
	assert.Equal(t, wantParams, parameters, "Parameters are not correctly formatted")
}
