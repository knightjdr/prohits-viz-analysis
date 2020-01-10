package settings

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Log shared settings", func() {
	It("should log settings", func() {
		settings := types.Settings{
			Abundance:       "abundanceColumn",
			AbundanceCap:    10.00,
			Condition:       "conditionColumn",
			Control:         "controlColumn",
			Files:           []string{"/folder/file1.txt", "file2.txt"},
			LogBase:         "2",
			MinAbundance:    5.00,
			Normalization:   "total",
			PrimaryFilter:   0.01,
			Readout:         "readoutColumn",
			ReadoutLength:   "readoutLengthColumn",
			Score:           "scoreColumn",
			ScoreType:       "lte",
			SecondaryFilter: 0.05,
			Type:            "dotplot",
		}

		expected := "Analysis type: dotplot\n\n" +
			"Files used\n" +
			"- file1.txt\n" +
			"- file2.txt\n\n" +
			"Columns used\n" +
			"- abundance: abundanceColumn\n" +
			"- condition: conditionColumn\n" +
			"- readout: readoutColumn\n" +
			"- score: scoreColumn\n" +
			"- control: controlColumn\n" +
			"- readout length: readoutLengthColumn\n\n" +
			"Readout abundance transformations\n" +
			"- control subtraction was performed\n" +
			"- readout length normalization was performed\n" +
			"- condition normalization was performed using total abundance\n" +
			"- data was log-transformed with base 2\n\n" +
			"Abundance\n" +
			"- minimum abundance required: 5\n" +
			"- abundances were capped at 10 for visualization\n\n" +
			"Scoring\n" +
			"- smaller scores are better\n" +
			"- primary filter: 0.01\n" +
			"- secondary filter: 0.05\n\n"

		var messages strings.Builder
		logSharedSettings(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log analysis type", func() {
	It("should log file names", func() {
		analysisType := "dotplot"

		expected := "Analysis type: dotplot\n\n"

		var messages strings.Builder
		logAnalysisType(&messages, analysisType)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log file names", func() {
	It("should log file names", func() {
		files := []string{"/folder/file1.txt", "file2.txt"}

		expected := "Files used\n" +
			"- file1.txt\n" +
			"- file2.txt\n\n"

		var messages strings.Builder
		logFiles(&messages, files)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log columns", func() {
	It("should log columns", func() {
		settings := types.Settings{
			Abundance: "abundanceColumn",
			Condition: "conditionColumn",
			Readout:   "readoutColumn",
			Score:     "scoreColumn",
		}

		expected := "Columns used\n" +
			"- abundance: abundanceColumn\n" +
			"- condition: conditionColumn\n" +
			"- readout: readoutColumn\n" +
			"- score: scoreColumn\n\n"

		var messages strings.Builder
		logColumns(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log optional columns", func() {
		settings := types.Settings{
			Abundance:     "abundanceColumn",
			Condition:     "conditionColumn",
			Control:       "controlColumn",
			Readout:       "readoutColumn",
			ReadoutLength: "readoutLengthColumn",
			Score:         "scoreColumn",
		}

		expected := "Columns used\n" +
			"- abundance: abundanceColumn\n" +
			"- condition: conditionColumn\n" +
			"- readout: readoutColumn\n" +
			"- score: scoreColumn\n" +
			"- control: controlColumn\n" +
			"- readout length: readoutLengthColumn\n\n"

		var messages strings.Builder
		logColumns(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log data transformations", func() {
	It("should log data transformations with total normalization", func() {
		settings := types.Settings{
			Control:       "controlColumn",
			LogBase:       "2",
			Normalization: "total",
			ReadoutLength: "readoutLengthColumn",
		}

		expected := "Readout abundance transformations\n" +
			"- control subtraction was performed\n" +
			"- readout length normalization was performed\n" +
			"- condition normalization was performed using total abundance\n" +
			"- data was log-transformed with base 2\n\n"

		var messages strings.Builder
		logTransformations(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log data transformations with readout normalization", func() {
		settings := types.Settings{
			Control:              "controlColumn",
			LogBase:              "2",
			Normalization:        "readout",
			NormalizationReadout: "readout1",
			ReadoutLength:        "readoutLengthColumn",
		}

		expected := "Readout abundance transformations\n" +
			"- control subtraction was performed\n" +
			"- readout length normalization was performed\n" +
			"- condition normalization was performed using the readout: readout1\n" +
			"- data was log-transformed with base 2\n\n"

		var messages strings.Builder
		logTransformations(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log abundance settings", func() {
	It("should log minimum abundance", func() {
		settings := types.Settings{
			MinAbundance: 5.00,
		}

		expected := "Abundance\n" +
			"- minimum abundance required: 5\n\n"

		var messages strings.Builder
		logAbundance(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log abundance cap", func() {
		settings := types.Settings{
			AbundanceCap: 10.00,
			MinAbundance: 5.00,
		}

		expected := "Abundance\n" +
			"- minimum abundance required: 5\n" +
			"- abundances were capped at 10 for visualization\n\n"

		var messages strings.Builder
		logAbundance(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log score settings", func() {
	It("should log score settings for lte", func() {
		settings := types.Settings{
			PrimaryFilter:   0.01,
			ScoreType:       "lte",
			SecondaryFilter: 0.05,
		}

		expected := "Scoring\n" +
			"- smaller scores are better\n" +
			"- primary filter: 0.01\n" +
			"- secondary filter: 0.05\n\n"

		var messages strings.Builder
		logScoring(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log score settings for gte", func() {
		settings := types.Settings{
			PrimaryFilter: 0.01,
			ScoreType:     "gte",
		}

		expected := "Scoring\n" +
			"- larger scores are better\n" +
			"- primary filter: 0.01\n\n"

		var messages strings.Builder
		logScoring(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})
