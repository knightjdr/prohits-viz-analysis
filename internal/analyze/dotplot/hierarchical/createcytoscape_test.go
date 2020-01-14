package hierarchical

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Create cytoscape files", func() {
	It("should create files", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("cytoscape", 0755)

		data := &sortedData{
			conditionDist: [][]float64{
				{0, 10, 74.2},
				{10, 0, 90.12},
				{74.2, 90.12, 0},
			},
			matrices: &types.Matrices{
				Conditions: []string{"condition1", "condition2", "condition3"},
				Readouts:   []string{"readout1", "readout2", "readout3"},
			},
			readoutDist: [][]float64{
				{0, 5, 34.7},
				{5, 0, 8.9},
				{34.7, 8.9, 0},
			},
		}
		fileData := []map[string]string{
			{"condition": "condition1", "readout": "readout1", "abundance": "10", "score": "0.01"},
			{"condition": "condition1", "readout": "readout2", "abundance": "5.5", "score": "0.02"},
			{"condition": "condition2", "readout": "readout1", "abundance": "1", "score": "0"},
			{"condition": "condition2", "readout": "readout3", "abundance": "75", "score": "0.01"},
		}
		settings := types.Settings{
			Abundance:     "AvgSpec",
			Condition:     "Condition",
			Readout:       "ReadoutGene",
			PrimaryFilter: 0.01,
			Score:         "BFDR",
			ScoreType:     "lte",
		}

		createCytoscape(fileData, data, settings)

		expectedConditionFile := "source\ttarget\tdistance\n" +
			"condition1\tcondition2\t10\n" +
			"condition1\tcondition3\t74.2\n" +
			"condition2\tcondition3\t90.12\n"
		actualConditionFile, _ := afero.ReadFile(fs.Instance, "cytoscape/Condition-Condition-cytoscape.txt")
		Expect(string(actualConditionFile)).To(Equal(expectedConditionFile), "should write condition distance matrix")

		expectedReadoutFile := "source\ttarget\tdistance\n" +
			"readout1\treadout2\t5\n" +
			"readout1\treadout3\t34.7\n" +
			"readout2\treadout3\t8.9\n"
		actualReadoutFile, _ := afero.ReadFile(fs.Instance, "cytoscape/ReadoutGene-ReadoutGene-cytoscape.txt")
		Expect(string(actualReadoutFile)).To(Equal(expectedReadoutFile), "should write readout distance matrix")

		expectedDataFile := "Condition\tReadoutGene\tAvgSpec\tBFDR\n" +
			"condition1\treadout1\t10.00\t0.01\n" +
			"condition2\treadout1\t1.00\t0\n" +
			"condition2\treadout3\t75.00\t0.01\n"
		actualDataFile, _ := afero.ReadFile(fs.Instance, "cytoscape/Condition-ReadoutGene-cytoscape.txt")
		Expect(string(actualDataFile)).To(Equal(expectedDataFile), "should write input file data")
	})
})

var _ = Describe("Write distance matrix for cytoscape", func() {
	It("should write matrix to file as three-column matix", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("cytoscape", 0755)

		source := []string{"condition1", "condition2", "condition3"}
		matrix := [][]float64{
			{0, 10, 74.2},
			{10, 0, 90.12},
			{74.2, 90.12, 0},
		}

		expected := "source\ttarget\tdistance\n" +
			"condition1\tcondition2\t10\n" +
			"condition1\tcondition3\t74.2\n" +
			"condition2\tcondition3\t90.12\n"

		writeDistanceCytoscape(matrix, source, source, "condition")
		actual, _ := afero.ReadFile(fs.Instance, "cytoscape/condition-condition-cytoscape.txt")
		Expect(string(actual)).To(Equal(expected))
	})
})

var _ = Describe("Write data for cytoscape", func() {
	It("should write input data to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("cytoscape", 0755)

		fileData := []map[string]string{
			{"condition": "condition1", "readout": "readout1", "abundance": "10", "score": "0.01"},
			{"condition": "condition1", "readout": "readout2", "abundance": "5.5", "score": "0.02"},
			{"condition": "condition2", "readout": "readout1", "abundance": "1", "score": "0"},
			{"condition": "condition2", "readout": "readout3", "abundance": "75", "score": "0.01"},
		}
		settings := types.Settings{
			Abundance:     "AvgSpec",
			Condition:     "Condition",
			Readout:       "ReadoutGene",
			PrimaryFilter: 0.01,
			Score:         "BFDR",
			ScoreType:     "lte",
		}

		expected := "Condition\tReadoutGene\tAvgSpec\tBFDR\n" +
			"condition1\treadout1\t10.00\t0.01\n" +
			"condition2\treadout1\t1.00\t0\n" +
			"condition2\treadout3\t75.00\t0.01\n"

		writeFileDataCytoscape(fileData, settings)
		actual, _ := afero.ReadFile(fs.Instance, "cytoscape/Condition-ReadoutGene-cytoscape.txt")
		Expect(string(actual)).To(Equal(expected))
	})
})
