package transform

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Control subtraction", func() {
	It("should substract control average from abundance", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"abundance": "10", "control": "2|3|10"},
				{"abundance": "1", "control": "2|3|10"},
				{"abundance": "10|5", "control": "2|3|10"},
				{"abundance": "10|5|2.5", "control": "2"},
			},
			Settings: types.Settings{
				Control: "ctrl",
			},
		}

		expected := []map[string]string{
			{"abundance": "5", "control": "2|3|10"},
			{"abundance": "0", "control": "2|3|10"},
			{"abundance": "5|0", "control": "2|3|10"},
			{"abundance": "8|3|0.5", "control": "2"},
		}

		controlSubtract(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})

	It("should not modify data when control subtract not requested", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"abundance": "10", "control": "2|3|10"},
				{"abundance": "1", "control": "2|3|10"},
				{"abundance": "10|5", "control": "2|3|10"},
				{"abundance": "10|5|2.5", "control": "2"},
			},
		}

		expected := []map[string]string{
			{"abundance": "10", "control": "2|3|10"},
			{"abundance": "1", "control": "2|3|10"},
			{"abundance": "10|5", "control": "2|3|10"},
			{"abundance": "10|5|2.5", "control": "2"},
		}

		controlSubtract(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})
})
