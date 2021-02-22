package scv

import (
	"encoding/json"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

func addExpression(expressionType string, data map[string]map[string]map[string]float64, idMaps map[string]map[string]string, settings types.Settings) {
	tissues, filename, expressionPrefix := defineExpressionParameters(expressionType, settings)

	if len(tissues) > 0 {
		file, err := fs.Instance.Open(filename)
		log.CheckError(err, true)

		bytes, err := afero.ReadAll(file)
		log.CheckError(err, true)

		var expression map[string]map[string]float64
		json.Unmarshal(bytes, &expression)

		for condition, conditionData := range data {
			for readout := range conditionData {
				mappedReadout := idMaps["readout"][readout]
				_, readoutHasExpressionData := expression[mappedReadout]

				for _, tissue := range tissues {
					key := fmt.Sprintf("%s expression - %s", expressionPrefix, tissue)
					if readoutHasExpressionData {
						data[condition][readout][key] = expression[mappedReadout][tissue]
					} else {
						data[condition][readout][key] = 0
					}
				}
			}
		}
	}
}

func defineExpressionParameters(expressionType string, settings types.Settings) ([]string, string, string) {
	if expressionType == "protein" {
		return settings.ProteinTissues, settings.ProteinExpressionFile, "Protein"
	}
	return settings.RnaTissues, settings.RnaExpressionFile, "RNA"
}
