package dotplot

import (
	"regexp"
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestLogParams(t *testing.T) {
	// Mock filesystem.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// TEST1: test typical run with hierarchical clustering.
	parameters := typedef.Parameters{
		Abundance:            "abundanceColumn",
		AnalysisType:         "dotplot",
		Condition:            "conditionColumn",
		ConditionClustering:  "",
		Clustering:           "hierarchical",
		ClusteringMethod:     "ward",
		FillColor:            "blueBlack",
		Control:              "controlColumn",
		Distance:             "canberra",
		Files:                []string{"file1.txt", "file2.txt"},
		LogBase:              "2",
		AbundanceCap:         50,
		MinAbundance:         0,
		Normalization:        "total",
		NormalizationReadout: "",
		Readout:              "readoutsColumn",
		ReadoutClustering:    "",
		ReadoutLength:        "readoutLengthColumn",
		PrimaryFilter:        0.01,
		Score:                "scoreColumn",
		ScoreType:            "lte",
		SecondaryFilter:      0.05,
	}
	LogParams(parameters)
	logfile, _ := afero.ReadFile(fs.Instance, "log.txt")
	want := "Analysis type: dotplot\r\n" +
		"\r\n" +
		"Files used\r\n" +
		"- file1.txt\r\n" +
		"- file2.txt\r\n" +
		"\r\n" +
		"Columns used\r\n" +
		"- abundance: abundanceColumn\r\n" +
		"- condition: conditionColumn\r\n" +
		"- readout: readoutsColumn\r\n" +
		"- score: scoreColumn\r\n" +
		"- control: controlColumn\r\n" +
		"- readout length: readoutLengthColumn\r\n" +
		"\r\n" +
		"Readout abundance transformations\r\n" +
		"- control subtraction was performed\r\n" +
		"- readout length normalization was performed\r\n" +
		"- condition normalization was performed using total abundance\r\n" +
		"- data was log-transformed with base 2\r\n" +
		"\r\n" +
		"Abundance\r\n" +
		"- minimum abundance required: 0\r\n" +
		"- abundances were capped at 50 for visualization\r\n" +
		"\r\n" +
		"Scoring\r\n" +
		"- smaller scores are better\r\n" +
		"- primary filter: 0.01\r\n" +
		"- secondary filter: 0.05\r\n" +
		"\r\n" +
		"Clustering\r\n" +
		"- hierarchical clustering was performed\r\n" +
		"- distance metric: canberra\r\n" +
		"- linkage method: ward\r\n"
	assert.Equal(t, want, string(logfile), "Logfile not correct")

	// TEST2: log normalization to a specific readout.
	parameters = typedef.Parameters{
		Abundance:            "abundanceColumn",
		AnalysisType:         "dotplot",
		Condition:            "conditionColumn",
		ConditionClustering:  "",
		Clustering:           "hierarchical",
		ClusteringMethod:     "ward",
		FillColor:            "blueBlack",
		Control:              "controlColumn",
		Distance:             "canberra",
		Files:                []string{"file1.txt", "file2.txt"},
		LogBase:              "2",
		AbundanceCap:         50,
		MinAbundance:         0,
		Normalization:        "readout",
		NormalizationReadout: "readoutX",
		Readout:              "readoutsColumn",
		ReadoutClustering:    "",
		ReadoutLength:        "readoutLengthColumn",
		PrimaryFilter:        0.01,
		Score:                "scoreColumn",
		ScoreType:            "lte",
		SecondaryFilter:      0.05,
	}
	LogParams(parameters)
	logfile, _ = afero.ReadFile(fs.Instance, "log.txt")
	want = "- condition normalization was performed using the readout: readoutX"
	matched, _ := regexp.MatchString(want, string(logfile))
	assert.True(t, matched, "Condition normalization to a readout not recorded in log")

	// TEST3: log higher scores are better.
	parameters = typedef.Parameters{
		Abundance:            "abundanceColumn",
		AnalysisType:         "dotplot",
		Condition:            "conditionColumn",
		ConditionClustering:  "",
		Clustering:           "hierarchical",
		ClusteringMethod:     "ward",
		FillColor:            "blueBlack",
		Control:              "controlColumn",
		Distance:             "canberra",
		Files:                []string{"file1.txt", "file2.txt"},
		LogBase:              "2",
		AbundanceCap:         50,
		MinAbundance:         0,
		Normalization:        "readout",
		NormalizationReadout: "readoutX",
		Readout:              "readoutsColumn",
		ReadoutClustering:    "",
		ReadoutLength:        "readoutLengthColumn",
		PrimaryFilter:        0.01,
		Score:                "scoreColumn",
		ScoreType:            "gte",
		SecondaryFilter:      0.05,
	}
	LogParams(parameters)
	logfile, _ = afero.ReadFile(fs.Instance, "log.txt")
	want = "- larger scores are better"
	matched, _ = regexp.MatchString(want, string(logfile))
	assert.True(t, matched, "Larger scores are better not recorded in log")

	// TEST3: log biclustering.
	parameters = typedef.Parameters{
		Abundance:            "abundanceColumn",
		AnalysisType:         "dotplot",
		Condition:            "conditionColumn",
		ConditionClustering:  "",
		Clustering:           "biclustering",
		ClusteringMethod:     "ward",
		FillColor:            "blueBlack",
		Control:              "controlColumn",
		Distance:             "canberra",
		Files:                []string{"file1.txt", "file2.txt"},
		LogBase:              "2",
		AbundanceCap:         50,
		MinAbundance:         0,
		Normalization:        "readout",
		NormalizationReadout: "readoutX",
		Readout:              "readoutsColumn",
		ReadoutClustering:    "",
		ReadoutLength:        "readoutLengthColumn",
		PrimaryFilter:        0.01,
		Score:                "scoreColumn",
		ScoreType:            "gte",
		SecondaryFilter:      0.05,
	}
	LogParams(parameters)
	logfile, _ = afero.ReadFile(fs.Instance, "log.txt")
	want = "- biclustering was performed"
	matched, _ = regexp.MatchString(want, string(logfile))
	assert.True(t, matched, "Biclustering not recorded in log")

	// TEST4: condition and readout clustering.
	parameters = typedef.Parameters{
		Abundance:            "abundanceColumn",
		AnalysisType:         "dotplot",
		Condition:            "conditionColumn",
		ConditionClustering:  "conditions",
		Clustering:           "none",
		ClusteringMethod:     "ward",
		FillColor:            "blueBlack",
		Control:              "controlColumn",
		Distance:             "canberra",
		Files:                []string{"file1.txt", "file2.txt"},
		LogBase:              "2",
		AbundanceCap:         50,
		MinAbundance:         0,
		Normalization:        "readout",
		NormalizationReadout: "readoutX",
		Readout:              "readoutsColumn",
		ReadoutClustering:    "readouts",
		ReadoutLength:        "readoutLengthColumn",
		PrimaryFilter:        0.01,
		Score:                "scoreColumn",
		ScoreType:            "gte",
		SecondaryFilter:      0.05,
	}
	LogParams(parameters)
	logfile, _ = afero.ReadFile(fs.Instance, "log.txt")
	want = "- no clustering was performed"
	matched, _ = regexp.MatchString(want, string(logfile))
	assert.True(t, matched, "No clustering not recorded in log")

	// TEST4: condition clustering alone.
	parameters = typedef.Parameters{
		Abundance:            "abundanceColumn",
		AnalysisType:         "dotplot",
		Condition:            "conditionColumn",
		ConditionClustering:  "none",
		Clustering:           "none",
		ClusteringMethod:     "ward",
		FillColor:            "blueBlack",
		Control:              "controlColumn",
		Distance:             "canberra",
		Files:                []string{"file1.txt", "file2.txt"},
		LogBase:              "2",
		AbundanceCap:         50,
		MinAbundance:         0,
		Normalization:        "readout",
		NormalizationReadout: "readoutX",
		Readout:              "readoutsColumn",
		ReadoutClustering:    "readouts",
		ReadoutLength:        "readoutLengthColumn",
		PrimaryFilter:        0.01,
		Score:                "scoreColumn",
		ScoreType:            "gte",
		SecondaryFilter:      0.05,
	}
	LogParams(parameters)
	logfile, _ = afero.ReadFile(fs.Instance, "log.txt")
	want = "- conditions were hierarchically clustered"
	matched, _ = regexp.MatchString(want, string(logfile))
	assert.True(t, matched, "Exclusive condition clustering not recorded in log")

	// TEST5: readout clustering alone.
	parameters = typedef.Parameters{
		Abundance:            "abundanceColumn",
		AnalysisType:         "dotplot",
		Condition:            "conditionColumn",
		ConditionClustering:  "conditions",
		Clustering:           "none",
		ClusteringMethod:     "ward",
		FillColor:            "blueBlack",
		Control:              "controlColumn",
		Distance:             "canberra",
		Files:                []string{"file1.txt", "file2.txt"},
		LogBase:              "2",
		AbundanceCap:         50,
		MinAbundance:         0,
		Normalization:        "readout",
		NormalizationReadout: "readoutX",
		Readout:              "readoutsColumn",
		ReadoutClustering:    "none",
		ReadoutLength:        "readoutLengthColumn",
		PrimaryFilter:        0.01,
		Score:                "scoreColumn",
		ScoreType:            "gte",
		SecondaryFilter:      0.05,
	}
	LogParams(parameters)
	logfile, _ = afero.ReadFile(fs.Instance, "log.txt")
	want = "- readouts were hierarchically clustered"
	matched, _ = regexp.MatchString(want, string(logfile))
	assert.True(t, matched, "Exclusive readout clustering not recorded in log")
}
