package heatmap

import (
	"encoding/json"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Settings", func() {
	It("should parse settings for heatmap and return string", func() {
		settings := types.Settings{
			AbundanceCap:  50,
			FillColor:     "blue",
			InvertColor:   true,
			MinAbundance:  10,
			PrimaryFilter: 0.01,
		}

		expectedSettings := map[string]map[string]map[string]interface{}{
			"main": map[string]map[string]interface{}{
				"current": map[string]interface{}{
					"abundanceCap":  50,
					"fillColor":     "blue",
					"imageType":     "heatmap",
					"invertColor":   true,
					"minAbundance":  10,
					"primaryFilter": 0.01,
				},
			},
		}
		expectedString, _ := json.Marshal(expectedSettings)
		expected := fmt.Sprintf("\"settings\": %s", expectedString)
		Expect(parseSettings("heatmap", settings)).To(Equal(expected))
	})

	It("should parse settings for dotplot", func() {
		settings := types.Settings{
			AbundanceCap:    50,
			EdgeColor:       "red",
			FillColor:       "blue",
			InvertColor:     true,
			MinAbundance:    10,
			PrimaryFilter:   0.01,
			SecondaryFilter: 0.05,
		}

		expectedSettings := map[string]map[string]map[string]interface{}{
			"main": map[string]map[string]interface{}{
				"current": map[string]interface{}{
					"abundanceCap":    50,
					"edgeColor":       "red",
					"fillColor":       "blue",
					"imageType":       "dotplot",
					"invertColor":     true,
					"minAbundance":    10,
					"primaryFilter":   0.01,
					"secondaryFilter": 0.05,
				},
			},
		}
		expectedString, _ := json.Marshal(expectedSettings)
		expected := fmt.Sprintf("\"settings\": %s", expectedString)
		Expect(parseSettings("dotplot", settings)).To(Equal(expected))
	})
})
