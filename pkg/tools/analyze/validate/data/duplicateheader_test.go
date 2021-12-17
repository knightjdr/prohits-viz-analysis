package data

import (
	"errors"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validate header", func() {
	It("should return nil when there are no duplicate headers", func() {
		data := []map[string]string{
			{"condition": "conditionA", "readout": "readoutA", "abundance": "1", "score": "0.05"},
			{"condition": "conditionB", "readout": "readoutB", "abundance": "2", "score": "0.01"},
			{"condition": "conditionA", "readout": "readoutC", "abundance": "3", "score": "0.05"},
			{"condition": "conditionB", "readout": "readoutC", "abundance": "4", "score": "0.01"},
			{"condition": "conditionC", "readout": "readoutA", "abundance": "5", "score": "0.02"},
			{"condition": "conditionC", "readout": "readoutB", "abundance": "6", "score": "0"},
		}
		settings := types.Settings{
			Abundance: "AvgSpec",
			Condition: "Bait",
			Readout:   "PreyGene",
			Score:     "FDR",
		}

		Expect(checkForDuplicateHeader(data, settings)).To(BeNil())
	})

	It("should return an error when the data matches a duplicated header", func() {
		data := []map[string]string{
			{"condition": "conditionA", "readout": "readoutA", "abundance": "1", "score": "0.05"},
			{"condition": "conditionB", "readout": "readoutB", "abundance": "2", "score": "0.01"},
			{"condition": "conditionA", "readout": "readoutC", "abundance": "3", "score": "0.05"},
			{"condition": "conditionB", "readout": "readoutC", "abundance": "4", "score": "0.01"},
			{"condition": "Bait", "readout": "PreyGene", "abundance": "AvgSpec", "score": "FDR"},
			{"condition": "conditionC", "readout": "readoutA", "abundance": "5", "score": "0.02"},
		}
		settings := types.Settings{
			Abundance: "AvgSpec",
			Condition: "Bait",
			Readout:   "PreyGene",
			Score:     "FDR",
		}

		expected := errors.New(
			"the file should only contain a single header row - duplicates detected; " +
				"this can happen when you manually merge multiple files and do not remove the " +
				"header rows from the additional files",
		)
		Expect(checkForDuplicateHeader(data, settings)).To(Equal(expected))
	})
})
