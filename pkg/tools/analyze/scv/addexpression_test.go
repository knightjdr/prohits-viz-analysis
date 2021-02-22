package scv

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Add expression data for SCV", func() {
	It("should add protein expression", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		expressionData := map[string]map[string]float64{
			"1": {
				"HEK 293": 5,
				"HeLa":    4.5,
			},
			"2": {
				"HEK 293": 3.5,
				"HeLa":    3.7,
			},
		}
		expressionFile, _ := json.Marshal(expressionData)

		fs.Instance.MkdirAll("helper-files", 0755)
		afero.WriteFile(fs.Instance, "helper-files/protein-expression.json", []byte(expressionFile), 0444)

		data := map[string]map[string]map[string]float64{
			"conditionA": {
				"readoutA": {
					"Abundance":  10,
					"FoldChange": 5,
				},
				"readoutC": {
					"Abundance":  20,
					"FoldChange": 10,
				},
			},
			"conditionB": {
				"readoutA": {
					"Abundance":  10,
					"FoldChange": 7,
				},
				"readoutB": {
					"Abundance":  20,
					"FoldChange": 10,
				},
			},
		}
		expressionType := "protein"
		idMaps := map[string]map[string]string{
			"condition": {
				"conditionA": "1",
				"conditionB": "2",
			},
			"readout": {
				"readoutA": "1",
				"readoutB": "2",
				"readoutC": "3",
			},
		}
		settings := types.Settings{
			ProteinExpressionFile: "helper-files/protein-expression.json",
			ProteinTissues:        []string{"HEK 293", "HeLa"},
		}

		expected := map[string]map[string]map[string]float64{
			"conditionA": {
				"readoutA": {
					"Abundance":                    10,
					"FoldChange":                   5,
					"Protein expression - HEK 293": 5,
					"Protein expression - HeLa":    4.5,
				},
				"readoutC": {
					"Abundance":                    20,
					"FoldChange":                   10,
					"Protein expression - HEK 293": 0,
					"Protein expression - HeLa":    0,
				},
			},
			"conditionB": {
				"readoutA": {
					"Abundance":                    10,
					"FoldChange":                   7,
					"Protein expression - HEK 293": 5,
					"Protein expression - HeLa":    4.5,
				},
				"readoutB": {
					"Abundance":                    20,
					"FoldChange":                   10,
					"Protein expression - HEK 293": 3.5,
					"Protein expression - HeLa":    3.7,
				},
			},
		}

		addExpression(expressionType, data, idMaps, settings)
		Expect(data).To(Equal(expected))
	})

	It("should add RNA expression", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		expressionData := map[string]map[string]float64{
			"1": {
				"HEK 293": 5,
				"HeLa":    4.5,
			},
			"2": {
				"HEK 293": 3.5,
				"HeLa":    3.7,
			},
		}
		expressionFile, _ := json.Marshal(expressionData)

		fs.Instance.MkdirAll("helper-files", 0755)
		afero.WriteFile(fs.Instance, "helper-files/rna-expression.json", []byte(expressionFile), 0444)

		data := map[string]map[string]map[string]float64{
			"conditionA": {
				"readoutA": {
					"Abundance":  10,
					"FoldChange": 5,
				},
				"readoutC": {
					"Abundance":  20,
					"FoldChange": 10,
				},
			},
			"conditionB": {
				"readoutA": {
					"Abundance":  10,
					"FoldChange": 7,
				},
				"readoutB": {
					"Abundance":  20,
					"FoldChange": 10,
				},
			},
		}
		expressionType := "rna"
		idMaps := map[string]map[string]string{
			"condition": {
				"conditionA": "1",
				"conditionB": "2",
			},
			"readout": {
				"readoutA": "1",
				"readoutB": "2",
				"readoutC": "3",
			},
		}
		settings := types.Settings{
			RnaExpressionFile: "helper-files/rna-expression.json",
			RnaTissues:        []string{"HEK 293", "HeLa"},
		}

		expected := map[string]map[string]map[string]float64{
			"conditionA": {
				"readoutA": {
					"Abundance":                10,
					"FoldChange":               5,
					"RNA expression - HEK 293": 5,
					"RNA expression - HeLa":    4.5,
				},
				"readoutC": {
					"Abundance":                20,
					"FoldChange":               10,
					"RNA expression - HEK 293": 0,
					"RNA expression - HeLa":    0,
				},
			},
			"conditionB": {
				"readoutA": {
					"Abundance":                10,
					"FoldChange":               7,
					"RNA expression - HEK 293": 5,
					"RNA expression - HeLa":    4.5,
				},
				"readoutB": {
					"Abundance":                20,
					"FoldChange":               10,
					"RNA expression - HEK 293": 3.5,
					"RNA expression - HeLa":    3.7,
				},
			},
		}

		addExpression(expressionType, data, idMaps, settings)
		Expect(data).To(Equal(expected))
	})
})
