package settings

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Parse", func() {
	It("should parse, convert and infer settings", func() {
		csv := []map[string]string{
			{"values": "5", "params": `{
					"type": "dotplot",
					"xAxis": "Prey",
					"yAxis": "Bait",
					"filterType": 1,
					"primary": 0.01,
					"secondary": 0.05,
					"score": "BFDR",
					"abundance": "AvgSpec",
					"invert": 1
				}`,
			},
			{"values": "1"},
			{"values": "23.8"},
			{"values": "7.5"},
		}

		expected := types.Settings{
			Abundance:       "AvgSpec",
			AbundanceCap:    50,
			EdgeColor:       "blue",
			FillColor:       "blue",
			InvertColor:     true,
			MinAbundance:    0,
			PrimaryFilter:   0.01,
			Score:           "BFDR",
			ScoreType:       "gte",
			SecondaryFilter: 0.05,
			Type:            "dotplot",
			XLabel:          "Prey",
			YLabel:          "Bait",
		}
		Expect(Parse(csv)).To(Equal(expected))
	})
})

var _ = Describe("Parse settings", func() {
	It("should parse JSON settings", func() {
		csv := []map[string]string{
			{"values": "5", "params": `{
					"type": "dotplot",
					"xAxis": "Prey",
					"yAxis": "Bait",
					"filterType": 1,
					"primary": 0.01,
					"secondary": 0.05,
					"score": "BFDR",
					"abundance": "AvgSpec",
					"invert": 1
				}`,
			},
		}

		expected := &jsonSettings{
			AbundanceColumn: "AvgSpec",
			InvertColor:     1,
			PrimaryFilter:   0.01,
			ScoreColumn:     "BFDR",
			ScoreType:       1,
			SecondaryFilter: 0.05,
			Type:            "dotplot",
			XLabel:          "Prey",
			YLabel:          "Bait",
		}
		Expect(parseSettings(csv)).To(Equal(expected))
	})

	It("should parse text settings", func() {
		csv := []map[string]string{
			{"params": "dotplot"},
			{"params": "1"},
			{"params": "0.01"},
			{"params": "0.05"},
			{"params": "BFDR"},
			{"params": "AvgSpec"},
			{"params": "1"},
		}

		expected := &jsonSettings{
			AbundanceColumn: "AvgSpec",
			InvertColor:     1,
			PrimaryFilter:   0.01,
			ScoreColumn:     "BFDR",
			ScoreType:       1,
			SecondaryFilter: 0.05,
			Type:            "dotplot",
		}
		Expect(parseSettings(csv)).To(Equal(expected))
	})
})

var _ = Describe("Has JSON settings", func() {
	It("should return true when 'params' appears to be JSON", func() {
		csv := []map[string]string{
			{"values": "5", "params": `{
					"type": "dotplot",
					"xAxis": "Prey",
					"yAxis": "Bait",
					"filterType": 1,
					"primary": 0.01,
					"secondary": 0.05,
					"score": "BFDR",
					"abundance": "AvgSpec",
					"invert": 1
				}`,
			},
		}

		Expect(hasJSONSettings(csv)).To(BeTrue())
	})

	It("should return false when 'params' does not appear to be JSON", func() {
		csv := []map[string]string{
			{"values": "5", "params": "dotplot"},
		}

		Expect(hasJSONSettings(csv)).To(BeFalse())
	})
})

var _ = Describe("Parse JSON", func() {
	It("should parse settings from JSON string", func() {
		csv := []map[string]string{
			{"values": "5", "params": `{
					"type": "dotplot",
					"xAxis": "Prey",
					"yAxis": "Bait",
					"filterType": 1,
					"primary": 0.01,
					"secondary": 0.05,
					"score": "BFDR",
					"abundance": "AvgSpec",
					"invert": 1
				}`,
			},
		}
		settings := &jsonSettings{}

		expected := &jsonSettings{
			AbundanceColumn: "AvgSpec",
			InvertColor:     1,
			PrimaryFilter:   0.01,
			ScoreColumn:     "BFDR",
			ScoreType:       1,
			SecondaryFilter: 0.05,
			Type:            "dotplot",
			XLabel:          "Prey",
			YLabel:          "Bait",
		}
		parseJSON(settings, csv)
		Expect(settings).To(Equal(expected))
	})
})

var _ = Describe("Parse text", func() {
	It("should parse settings from text", func() {
		settings := &jsonSettings{}
		csv := []map[string]string{
			{"params": "dotplot"},
			{"params": "1"},
			{"params": "0.01"},
			{"params": "0.05"},
			{"params": "BFDR"},
			{"params": "AvgSpec"},
			{"params": "1"},
		}

		expected := &jsonSettings{
			AbundanceColumn: "AvgSpec",
			InvertColor:     1,
			PrimaryFilter:   0.01,
			ScoreColumn:     "BFDR",
			ScoreType:       1,
			SecondaryFilter: 0.05,
			Type:            "dotplot",
		}
		parseText(settings, csv)
		Expect(settings).To(Equal(expected))
	})
})
