package scatter

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Legend scatter", func() {
	It("should parse legend and return string", func() {
		legend := []map[string]string{
			{"color": "#ff0000", "text": "point1"},
			{"color": "#00ff00", "text": "point2"},
		}

		expectedString, _ := json.Marshal(legend)
		expected := fmt.Sprintf("\"legend\": %s", expectedString)
		Expect(parseLegend(legend)).To(Equal(expected))
	})
})
