package heatmap

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Columns", func() {
	It("should parse settings for heatmap and return string", func() {
		columns := []string{"bait1", "bait2", "bait3"}

		expectedString, _ := json.Marshal(columns)
		expected := fmt.Sprintf("\"columnDB\": %s", expectedString)
		Expect(parseColumns(columns)).To(Equal(expected))
	})
})
