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
			Condition:       "conditionColumn",
			Control:         "controlColumn",
			Files:           []string{"/folder/file1.txt", "file2.txt"},
			LogBase:         "2",
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
			"readoutColumn abundance transformations\n" +
			"- control subtraction was performed\n" +
			"- readoutColumn length normalization was performed\n" +
			"- conditionColumn normalization was performed using total abundance\n" +
			"- data was log-transformed with base 2\n\n"

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
	It("should log no data transformations", func() {
		settings := types.Settings{}

		expected := ""

		var messages strings.Builder
		logTransformations(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log data transformations with total normalization", func() {
		settings := types.Settings{
			Condition:              "Bait",
			Control:                "controlColumn",
			LogBase:                "2",
			MockConditionAbundance: true,
			Normalization:          "total",
			Readout:                "Prey",
			ReadoutLength:          "readoutLengthColumn",
		}

		expected := "Prey abundance transformations\n" +
			"- control subtraction was performed\n" +
			"- Prey length normalization was performed\n" +
			"- Bait normalization was performed using total abundance\n" +
			"- data was log-transformed with base 2\n" +
			"- abundance values were mocked for Bait(s) with missing values\n\n"

		var messages strings.Builder
		logTransformations(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log data transformations with readout normalization", func() {
		settings := types.Settings{
			Condition:              "Bait",
			Control:                "controlColumn",
			LogBase:                "2",
			MockConditionAbundance: true,
			Normalization:          "readout",
			NormalizationReadout:   "readout1",
			Readout:                "Prey",
			ReadoutLength:          "readoutLengthColumn",
		}

		expected := "Prey abundance transformations\n" +
			"- control subtraction was performed\n" +
			"- Prey length normalization was performed\n" +
			"- Bait normalization was performed using the Prey: readout1\n" +
			"- data was log-transformed with base 2\n" +
			"- abundance values were mocked for Bait(s) with missing values\n\n"

		var messages strings.Builder
		logTransformations(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})
