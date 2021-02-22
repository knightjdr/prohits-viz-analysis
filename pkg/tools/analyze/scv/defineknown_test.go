package scv

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Define knownness for SCV", func() {
	It("should add abundance columns to data", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		interactionData := map[string][]string{
			"1": {"3"},
		}
		interactionFile, _ := json.Marshal(interactionData)

		fs.Instance.MkdirAll("helper-files", 0755)
		afero.WriteFile(fs.Instance, "helper-files/interactions.json", []byte(interactionFile), 0444)

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
			Known:     "interaction",
			KnownFile: "helper-files/interactions.json",
		}

		expected := map[string]map[string]bool{
			"conditionA": {
				"readoutA": false,
				"readoutC": true,
			},
			"conditionB": {
				"readoutA": false,
				"readoutB": false,
			},
		}

		Expect(defineKnown(data, idMaps, settings)).To(Equal(expected))
	})
})
