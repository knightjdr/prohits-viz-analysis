package heatmap

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Settings", func() {
	It("should parse settings", func() {
		settings := map[string]interface{}{
			"abundanceCap":    50,
			"edgeColor":       "red",
			"fillColor":       "blue",
			"imageType":       "dotplot",
			"invertColor":     true,
			"minAbundance":    10,
			"primaryFilter":   0.01,
			"secondaryFilter": 0.05,
		}

		expectedString, _ := json.Marshal(settings)
		expected := fmt.Sprintf("\"settings\": %s", expectedString)
		Expect(parseSettings(settings)).To(Equal(expected))
	})
})
