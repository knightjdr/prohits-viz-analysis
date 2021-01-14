package scatter

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Settings", func() {
	It("should parse settings", func() {
		settings := map[string]interface{}{
			"abundanceCap": 50,
			"imageType":    "scatter",
		}

		expectedString, _ := json.Marshal(settings)
		expected := fmt.Sprintf("\"settings\": %s", expectedString)
		Expect(parseSettings(settings)).To(Equal(expected))
	})
})
