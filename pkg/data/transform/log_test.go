package transform

import (
	"math"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Log transform data", func() {
	It("should log transform abundance values", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"abundance": "2"},
				{"abundance": "1"},
				{"abundance": "8|4"},
				{"abundance": "16|1|4"},
			},
			Settings: types.Settings{
				LogBase: "2",
				Type:    "dotplot",
			},
		}

		expected := []map[string]string{
			{"abundance": "1"},
			{"abundance": "0"},
			{"abundance": "3|2"},
			{"abundance": "4|0|2"},
		}

		logTransform(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})

	It("should not log transform abundance values when base is not recognized", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"abundance": "2"},
				{"abundance": "1"},
				{"abundance": "8|4"},
				{"abundance": "16|1|4"},
			},
			Settings: types.Settings{
				Type: "dotplot",
			},
		}

		expected := []map[string]string{
			{"abundance": "2"},
			{"abundance": "1"},
			{"abundance": "8|4"},
			{"abundance": "16|1|4"},
		}

		logTransform(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})
})

var _ = Describe("Log transformation permitted", func() {
	It("should transform when base and tool are valid", func() {
		tests := []map[string]string{
			{"logBase": "2", "tool": "dotplot"},
			{"logBase": "e", "tool": "dotplot"},
			{"logBase": "10", "tool": "dotplot"},
			{"logBase": "2", "tool": "correlation"},
			{"logBase": "e", "tool": "correlation"},
			{"logBase": "10", "tool": "correlation"},
		}
		for _, test := range tests {
			Expect(shouldLogTransform(test["logBase"], test["tool"])).To(BeTrue())
		}
	})

	It("should not transform when base or tool are invalid", func() {
		tests := []map[string]string{
			{"logBase": "1", "tool": "dotplot"},
			{"logBase": "2", "tool": "specificity"},
			{"logBase": "e", "tool": "specificity"},
			{"logBase": "10", "tool": "specificity"},
		}
		for _, test := range tests {
			Expect(shouldLogTransform(test["logBase"], test["tool"])).To(BeFalse())
		}
	})
})

var _ = Describe("Log transformation function", func() {
	It("should transform to base 2", func() {
		logFunc := getLogFunction("2")
		tests := []map[string]float64{
			{"input": 0, "expected": 0},
			{"input": 1, "expected": 0},
			{"input": 2, "expected": 1},
			{"input": 4, "expected": 2},
		}
		for _, test := range tests {
			Expect(logFunc(test["input"])).To(Equal(test["expected"]))
		}
	})

	It("should transform to base 10", func() {
		logFunc := getLogFunction("10")
		tests := []map[string]float64{
			{"input": 0, "expected": 0},
			{"input": 1, "expected": 0},
			{"input": 10, "expected": 1},
			{"input": 100, "expected": 2},
		}
		for _, test := range tests {
			Expect(logFunc(test["input"])).To(Equal(test["expected"]))
		}
	})

	It("should transform to base e", func() {
		logFunc := getLogFunction("e")
		tests := []map[string]float64{
			{"input": 0, "expected": 0},
			{"input": 1, "expected": 0},
			{"input": math.E, "expected": 1},
			{"input": math.SqrtE, "expected": 0.5},
		}
		for _, test := range tests {
			Expect(logFunc(test["input"])).To(Equal(test["expected"]))
		}
	})
})

var _ = Describe("Adjust abundance by log", func() {
	It("should log transform abundance values", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"abundance": "2"},
				{"abundance": "1"},
				{"abundance": "8|4"},
				{"abundance": "16|1|4"},
			},
		}
		logFunc := getLogFunction("2")

		expected := []map[string]string{
			{"abundance": "1"},
			{"abundance": "0"},
			{"abundance": "3|2"},
			{"abundance": "4|0|2"},
		}

		adjustAbundanceByLog(analysis, logFunc)
		Expect(analysis.Data).To(Equal(expected))
	})
})
